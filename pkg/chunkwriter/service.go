package chunkwriter

import (
	"io"
	"os"
	"sync"

	"github.com/kaz-as/zip/pkg/logger"
)

type ChunkWriter interface {
	WriteChunk(from io.Reader, path string, start int64, n int64) error
}

type file struct {
	mx sync.Mutex
	f  *os.File
}

type ChunkService struct {
	l logger.Interface

	mx    sync.Mutex
	files map[string]file
}

func NewService(l logger.Interface) *ChunkService {
	return &ChunkService{
		l:     l,
		mx:    sync.Mutex{},
		files: make(map[string]file),
	}
}

func (c *ChunkService) WriteChunk(from io.Reader, path string, start int64, n int64) error {
	//TODO implement me
	panic("implement me")
}
