package config

import (
	"os"
)

// Config struct holds the configuration settings for the application.
type Config struct {
	LogLevel string // This field stores the log level (e.g., "info", "debug").
	BaseDir  string // This field stores the base directory for file management.
}

// New function creates a new Config instance and returns it along with any error encountered.
func New() (*Config, error) {
	// Get the value of the environment variable "FILE_MANAGER_BASE_DIR".
	baseDir := os.Getenv("FILE_MANAGER_BASE_DIR")
	
	// If the environment variable is not set, use the current directory (".") as the default.
	if baseDir == "" {
		baseDir = "."
	}

	// Create and return a new Config instance with the log level and base directory.
	return &Config{
		LogLevel: os.Getenv("LOG_LEVEL"), // Get the log level from the environment variable "LOG_LEVEL".
		BaseDir:  baseDir, // Use the base directory we determined earlier.
	}, nil // Return nil for the error since there was no error in creating the config.
}
