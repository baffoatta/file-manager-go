package main

import (
	"log"
	"os"

	"github.com/baffoatta/filemanager/internal/app"
	"github.com/baffoatta/filemanager/internal/config"
	"github.com/baffoatta/filemanager/pkg/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	l, err := logger.New(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer l.Sync()	

	app := app.New(cfg, l)
	if err := app.Run(); err != nil {
		l.Error("Failed to run application", err)
		os.Exit(1)
	}
}
