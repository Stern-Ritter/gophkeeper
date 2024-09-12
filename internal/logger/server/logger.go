package server

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ServerLogger defines an interface for a logger with methods for structured logging.
type ServerLogger interface {
	Sugar() *zap.SugaredLogger
	Named(s string) *zap.Logger
	WithOptions(opts ...zap.Option) *zap.Logger
	With(fields ...zap.Field) *zap.Logger
	WithLazy(fields ...zap.Field) *zap.Logger
	Level() zapcore.Level
	Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
	Log(lvl zapcore.Level, msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Sync() error
	Core() zapcore.Core
	Name() string
}

// ServerLoggerImpl is an implementation of the ServerLogger interface using
// the zap.Logger as the underlying logging mechanism.
type ServerLoggerImpl struct {
	*zap.Logger
}

// Initialize initializes a ServerLogger with the specified logging level.
func Initialize(level string) (ServerLogger, error) {
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
	return &ServerLoggerImpl{logger}, nil
}
