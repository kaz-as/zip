package app

import (
	"database/sql"
	"fmt"
	archiverepo "github.com/kaz-as/zip/internal/archive/repository/postgres"
	"github.com/kaz-as/zip/internal/archive/usecase"
	chunkrepo "github.com/kaz-as/zip/internal/chunk/repository/postgres"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/kaz-as/zip/config"
	"github.com/kaz-as/zip/internal/handlers"
	"github.com/kaz-as/zip/internal/middlewares"
	"github.com/kaz-as/zip/pkg/archive"
	"github.com/kaz-as/zip/pkg/httpserver"
	"github.com/kaz-as/zip/pkg/logger"
)

type App struct {
	log  logger.Interface
	srv  *httpserver.Server
	conn *sql.DB
}

func New(cfg *config.Config) (app App, _ error) {
	var err error

	var l logger.Interface = logger.New(cfg.Level)

	app.log = l

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	app.conn, err = sql.Open("pgx", dsn)
	if err != nil {
		return app, fmt.Errorf("db open: %s", err)
	}
	err = app.conn.Ping()
	if err != nil {
		return app, fmt.Errorf("db ping: %s", err)
	}

	archiveService := archive.NewService()

	archiveRepo := archiverepo.New(app.conn)
	chunkRepo := chunkrepo.New(app.conn)
	archiveUseCase := usecase.New(archiveRepo, chunkRepo, cfg.DB.Timeout)

	h, err := handlers.New(
		l,
		archiveUseCase,
		archiveService,
		[]middlewares.Middleware{})
	if err != nil {
		return app, fmt.Errorf("creating main handler: %s", err)
	}

	mwGlobal := middlewares.Chain([]middlewares.Middleware{
		middlewares.Logger(l),
		middlewares.Recoverer(l),
	})

	app.srv = httpserver.New(
		mwGlobal(h),
		httpserver.Port(cfg.Port),
		httpserver.Logger(l),
	)

	return app, nil
}

// Close must be called at app stop or exit
func (app App) Close() {
	if app.conn != nil {
		err := app.conn.Close()
		if err != nil {
			app.log.Error("db connection cannot be closed: %s", err)
		}
	}

	if app.srv != nil {
		err := app.srv.Shutdown()
		if err != nil {
			app.log.Error("srv.Shutdown: %s", err)
		}
	}
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
		ret = fmt.Errorf("srv.Notify: %w", err)
		app.log.Error(ret)
	}

	return
}
