package main

import (
	"github.com/kaz-as/zip/config"
	"github.com/kaz-as/zip/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
