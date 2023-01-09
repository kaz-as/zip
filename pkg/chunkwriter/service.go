package chunkwriter

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/kaz-as/zip/pkg/logger"
)

type ErrIncorrectLength struct {
	path    string
	needed  int64
	written int64
}

func (e *ErrIncorrectLength) Error() string {
	if e.written > e.needed {
		return fmt.Sprintf("file %s: should be written %d bytes, but buffer size is larger", e.path, e.needed)
	}
	return fmt.Sprintf("file %s: should be written %d bytes, but size of buffer = %d", e.path, e.needed, e.written)
}

type ChunkWriter interface {
	WriteChunk(from io.Reader, path string, start int64, n int64) error
}

type file struct {
	mx         sync.Mutex
	f          *os.File
	lastAccess time.Time
}

func newFile(path string) (*file, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &file{
		mx:         sync.Mutex{},
		f:          f,
		lastAccess: time.Now(),
	}, nil
}

const (
	durationForStartCleaner  = time.Minute * 10
	initCallsForStartCleaner = 100
	durationAfterFileAccess  = time.Minute * 3
)

type ChunkService struct {
	l logger.Interface

	rwmx sync.RWMutex

	mx    sync.Mutex
	files map[string]*file

	callsForStartCleaner int
	prevStartedCleaner   time.Time

	exit chan struct{}
}

func NewService(l logger.Interface) *ChunkService {
	s := &ChunkService{
		l:                    l,
		files:                make(map[string]*file),
		callsForStartCleaner: initCallsForStartCleaner,
		exit:                 make(chan struct{}),
	}

	go func() {
		for {
			select {
			case <-time.After(s.prevStartedCleaner.Add(durationForStartCleaner).Sub(time.Now())):
				s.clean()
			case <-s.exit:
				return
			}
		}
	}()

	return s
}

func (c *ChunkService) Close() error {
	close(c.exit)
	return nil
}

func (c *ChunkService) WriteChunk(from io.Reader, path string, start int64, n int64) error {
	c.rwmx.RLock()
	defer c.rwmx.RUnlock()

	defer func() { c.callsForStartCleaner-- }()

	f, err := c.getFile(path)
	if err != nil {
		return fmt.Errorf("get file by path %s failed: %w", path, err)
	}

	f.mx.Lock()
	defer f.mx.Unlock()

	_, err = f.f.Seek(start, 0)
	if err != nil {
		return fmt.Errorf("seek file by path %s failed: %w", path, err)
	}

	reader := io.LimitReader(from, n)
	written, err := io.Copy(f.f, reader)
	if err != nil {
		return fmt.Errorf("write to the file by path %s failed: %w", path, err)
	}
	if written != n {
		return &ErrIncorrectLength{
			path:    path,
			needed:  n,
			written: written,
		}
	}
	unaryCheckBuffer := make([]byte, 1)
	_, err = from.Read(unaryCheckBuffer)
	if err == nil || !errors.Is(err, io.EOF) {
		return &ErrIncorrectLength{
			path:    path,
			needed:  n,
			written: written + 1,
		}
	}

	if c.callsForStartCleaner <= 0 {
		go c.clean()
	}

	return nil
}

func (c *ChunkService) getFile(path string) (f *file, err error) {
	c.mx.Lock()
	defer c.mx.Unlock()

	f, ok := c.files[path]
	if !ok {
		f, err = newFile(path)
	}

	if f != nil {
		f.lastAccess = time.Now()
	}

	return
}

func (c *ChunkService) clean() {
	c.rwmx.Lock()
	defer c.rwmx.Unlock()

	for k, v := range c.files {
		if v.lastAccess.Add(durationAfterFileAccess).Before(time.Now()) {
			err := v.f.Close()
			if err != nil {
				c.l.Error("file %s writer failed: %s", k, err)
			}

			delete(c.files, k)
		}
	}

	c.callsForStartCleaner = initCallsForStartCleaner
	c.prevStartedCleaner = time.Now()
}
