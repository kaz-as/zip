package httpserver

import (
	"log"
	"net"
	"unsafe"

	"github.com/kaz-as/zip/pkg/logger"
)

type Option func(*Server)

func Port(port string) Option {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}

type srvErrLog struct {
	logger logger.Interface
}

func (s srvErrLog) Write(p []byte) (int, error) {
	s.logger.Error(*(*string)(unsafe.Pointer(&p)))
	return len(p), nil
}

func Logger(l logger.Interface) Option {
	return func(s *Server) {
		lg := log.New(srvErrLog{logger: l}, "server error: ", log.LstdFlags|log.Llongfile)
		s.server.ErrorLog = lg
		s.printf = l.Info
	}
}
