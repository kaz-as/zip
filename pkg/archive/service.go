package archive

import (
	"context"
	"io"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (a *Service) Zip(ctx context.Context, files ...File) (io.Reader, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Service) Unzip(ctx context.Context, archive io.Reader, dir string) error {
	//TODO implement me
	panic("implement me")
}
