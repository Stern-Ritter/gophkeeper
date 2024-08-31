package client

import "go.uber.org/zap"

// ClientLogger wraps a zap.Logger.
type ClientLogger struct {
	*zap.Logger
}

// Initialize initializes a ClientLogger with the specified logging level.
func Initialize(level string) (*ClientLogger, error) {
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
	return &ClientLogger{logger}, nil
}
