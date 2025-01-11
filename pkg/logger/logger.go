// File: pkg/logger/logger.go
package logger

import (
    "go.uber.org/zap"
)

// Logger struct that embeds a SugaredLogger from zap for easier logging
type Logger struct {
    *zap.SugaredLogger
}

// New function creates a new Logger instance with the specified log level
func New(level string) (*Logger, error) {
    config := zap.NewProductionConfig()
    
    // If a log level is provided, set it in the configuration
    if level != "" {
        var zapLevel zap.AtomicLevel
        // Unmarshal the provided log level into zapLevel
        if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
            return nil, err
        }
        config.Level = zapLevel
    }

    // Build the logger with the specified configuration
    logger, err := config.Build()
    if err != nil {
        return nil, err
    }

    sugar := logger.Sugar()
    
    return &Logger{sugar}, nil
}

// Info method logs an informational message with optional key-value pairs
func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
    l.SugaredLogger.Infow(msg, keysAndValues...)
}

// Error method logs an error message with optional key-value pairs
func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
    l.SugaredLogger.Errorw(msg, keysAndValues...)
}