package logger

import (
    "go.uber.org/zap"
)

type Logger struct {
    *zap.SugaredLogger
}

func New(level string) (*Logger, error) {
    config := zap.NewProductionConfig()
    
    if level != "" {
        var zapLevel zap.AtomicLevel
        if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
            return nil, err
        }
        config.Level = zapLevel
    }

    logger, err := config.Build()
    if err != nil {
        return nil, err
    }

    return &Logger{logger.Sugar()}, nil
}