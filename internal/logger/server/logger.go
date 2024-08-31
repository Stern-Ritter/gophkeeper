package server

import (
	"go.uber.org/zap"
)

// ServerLogger wraps a zap.Logger.
type ServerLogger struct {
	*zap.Logger
}

// Initialize initializes a ServerLogger with the specified logging level.
func Initialize(level string) (*ServerLogger, error) {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return &ServerLogger{logger}, nil
}
