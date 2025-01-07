package main

import (
	"log" // Importing the log package to log messages
	"os"  // Importing the os package to interact with the operating system

	"github.com/baffoatta/filemanager/internal/app"    // Importing the app package for application logic
	"github.com/baffoatta/filemanager/internal/config" // Importing the config package to handle configuration
	"github.com/baffoatta/filemanager/pkg/logger"     // Importing the logger package for logging functionality
)

func main() {
	// Create a new configuration object
	cfg, err := config.New()
	if err != nil {
		// If there is an error creating the config, log the error and exit the program
		log.Fatalf("Failed to initialize config: %v", err)
	}

	// Create a new logger with the specified log level from the configuration
	l, err := logger.New(cfg.LogLevel)
	if err != nil {
		// If there is an error creating the logger, log the error and exit the program
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer l.Sync() // Ensure that all log messages are flushed before the program exits

	// Create a new application instance with the configuration and logger
	app := app.New(cfg, l)
	// Run the application and check for errors
	if err := app.Run(); err != nil {
		// If the application fails to run, log the error and exit the program
		l.Error("Failed to run application", err)
		os.Exit(1) // Exit the program with a status code of 1 indicating an error
	}
}
