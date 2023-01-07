package archive

import (
	"context"
	"io"
)

type Archiver interface {
	Zip(ctx context.Context, files ...File) (io.Reader, error)
	Unzip(ctx context.Context, archive io.Reader, dir string) error
}
