package archive

import (
	"context"
	"io"
)

type File struct {
	In          io.ReadCloser
	OutLocation []string
	OutName     string
}

type Archiver interface {
	// Zip archives each File to File.OutLocation with File.OutName. Each File.In must be closed if no error returned
	Zip(ctx context.Context, files ...*File) (io.ReadCloser, error)
	Unzip(ctx context.Context, archive io.Reader, dir string) error
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (a *Service) Zip(ctx context.Context, files ...*File) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Service) Unzip(ctx context.Context, archive io.Reader, dir string) error {
	//TODO implement me
	panic("implement me")
}
