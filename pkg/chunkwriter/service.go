package chunkwriter

import (
	"io"
	"os"
	"sync"
)

type ChunkWriter interface {
	WriteChunk(from io.Reader, path string, start int64, n int64) error
}

type file struct {
	mx sync.Mutex
	f  *os.File
}

type ChunkService struct {
	mx    sync.Mutex
	files map[string]file
}

func NewService() *ChunkService {
	return &ChunkService{
		mx:    sync.Mutex{},
		files: make(map[string]file),
	}
}

func (c *ChunkService) WriteChunk(from io.Reader, path string, start int64, n int64) error {
	//TODO implement me
	panic("implement me")
}
