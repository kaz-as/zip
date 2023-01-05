package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kaz-as/zip/config"
	"github.com/kaz-as/zip/internal/handlers"
	"github.com/kaz-as/zip/internal/middlewares"
	"github.com/kaz-as/zip/pkg/httpserver"
	"github.com/kaz-as/zip/pkg/logger"
)

type App struct {
	log logger.Interface
	srv *httpserver.Server
}

func New(cfg *config.Config) (App, error) {
	var l logger.Interface = logger.New(cfg.Level)

	h, err := handlers.New(l, []middlewares.Middleware{})
	if err != nil {
		return App{}, fmt.Errorf("creating main handler: %s", err)
	}

	mwGlobal := middlewares.Chain([]middlewares.Middleware{
		middlewares.Logger(l),
		middlewares.Recoverer(l),
	})

	srv := httpserver.New(
		mwGlobal(h),
		httpserver.Port(cfg.Port),
		httpserver.Logger(l),
	)

	return App{
		log: l,
		srv: srv,
	}, nil
}

// Run returns only application error that cause shutdown, else nil
func (app App) Run() (ret error) {
	app.srv.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		app.log.Info("app - Run - signal: " + s.String())
	case err := <-app.srv.Notify():
		ret = fmt.Errorf("app - Run - srv.Notify: %w", err)
		app.log.Error(ret)
	}

	err := app.srv.Shutdown()
	if err != nil {
		app.log.Error(fmt.Errorf("app - Run - srv.Shutdown: %w", err))
	}

	return
}
