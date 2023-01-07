package main

import (
	"fmt"
	"log"

	"github.com/kaz-as/zip/config"
	"github.com/kaz-as/zip/internal/app"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("config error: %s", err)
	}

	application, err := app.New(cfg)
	defer application.Close()
	if err != nil {
		return fmt.Errorf("application create error: %s", err)
	}

	err = application.Run()
	if err != nil {
		return fmt.Errorf("application error: %s", err)
	}

	return nil
}
