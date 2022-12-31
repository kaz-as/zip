package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kaz-as/zip/config"
	"github.com/kaz-as/zip/pkg/httpserver"
	"github.com/kaz-as/zip/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	var l logger.Interface = logger.New(cfg.Level)

	r := mux.NewRouter().StrictSlash(true)

	srv := httpserver.New(r, httpserver.Port(cfg.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-srv.Notify():
		l.Error(fmt.Errorf("app - Run - srv.Notify: %w", err))
	}

	err := srv.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - srv.Shutdown: %w", err))
	}
}
