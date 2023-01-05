package main

import (
	"log"

	"github.com/kaz-as/zip/config"
	"github.com/kaz-as/zip/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("Application create error: %s", err)
	}

	err = application.Run()
	if err != nil {
		log.Fatalf("Application error: %s", err)
	}
}
