package httpserver

import (
	"context"
	"log"
	"net/http"
	"time"
)

const (
	_defaultReadHeaderTimeout = 5 * time.Second
	_defaultAddr              = ":80"
	_defaultShutdownTimeout   = 3 * time.Second
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
	printf          func(string, ...interface{})
}

func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:           handler,
		ReadHeaderTimeout: _defaultReadHeaderTimeout,
		Addr:              _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		printf:          log.Printf,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start() {
	go func() {
		s.printf("starting server: '%s'", s.server.Addr)
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
