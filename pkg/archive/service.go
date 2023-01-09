package archive

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/kaz-as/zip/pkg/logger"
)

type File struct {
	In          io.ReadCloser
	OutLocation []string
	OutName     string
}

type Archiver interface {
	// Zip archives each File to File.OutLocation with File.OutName. Each File.In must be closed if no error returned
	Zip(ctx context.Context, files ...*File) (io.ReadCloser, error)
	Unzip(ctx context.Context, archive io.ReaderAt, size int64, dest string) error
}

type ErrBadFile struct {
	Location string
}

func (e ErrBadFile) Error() string {
	return fmt.Sprintf("cannot add file %s to archive", e.Location)
}

type Service struct {
	l logger.Interface // logging of async errors
}

func NewService(l logger.Interface) *Service {
	return &Service{l: l}
}

func (a *Service) Zip(_ context.Context, files ...*File) (io.ReadCloser, error) {
	r, w := io.Pipe()
	zw := zip.NewWriter(w)

	go func() {
		defer func() {
			err := zw.Close()
			if err != nil {
				a.l.Error("cannot close zip: %s", err)
			}
		}()

		var firstUnclosedFile int

		defer func() {
			for i := firstUnclosedFile; i < len(files); i++ {
				file := files[i]
				if file == nil {
					continue
				}
				err := file.In.Close()
				if err != nil {
					a.l.Error("cannot close zip: %s", err)
				}
			}
		}()

		for _, file := range files {
			path := make([]string, len(file.OutLocation)+1)
			copy(path, file.OutLocation)
			path[len(path)-1] = file.OutName

			pathStr := filepath.Join(path...)

			zf, err := zw.Create(pathStr)
			if err != nil {
				a.l.Error("cannot add file to zip on path: %s", pathStr)
				return
			}

			_, err = io.Copy(zf, file.In)
			if err != nil {
				a.l.Error("io.Copy failed during zip on path: %s", pathStr)
				return
			}

			firstUnclosedFile++
		}
	}()

	return r, nil
}

func (a *Service) Unzip(_ context.Context, archive io.ReaderAt, size int64, dest string) error {
	r, err := zip.NewReader(archive, size)
	if err != nil {
		return fmt.Errorf("create zip reader failed: %w", err)
	}

	err = os.MkdirAll(dest, 0755)
	if err != nil {
		return fmt.Errorf("create folder %s failed: %w", dest, err)
	}

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(path, f.Mode())
			if err != nil {
				return fmt.Errorf("create folder %s failed: %w", path, err)
			}
		} else {
			pth := filepath.Dir(path)
			err = os.MkdirAll(pth, f.Mode())
			if err != nil {
				return fmt.Errorf("create folder %s failed: %w", pth, err)
			}

			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
