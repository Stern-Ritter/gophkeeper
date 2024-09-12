package server

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Logger defines an interface for a logger that can be configured with options.
type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	WithOptions(opts ...zap.Option) *zap.Logger
}

// NewInterceptorLogger creates a new logging.Logger that logs messages using the provided Logger.
func NewInterceptorLogger(l Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)

		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			value := fields[i+1]

			switch v := value.(type) {
			case string:
				f = append(f, zap.String(key.(string), v))
			case int:
				f = append(f, zap.Int(key.(string), v))
			case bool:
				f = append(f, zap.Bool(key.(string), v))
			default:
				f = append(f, zap.Any(key.(string), v))
			}
		}

		logger := l.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			logger.Debug(msg)
		case logging.LevelInfo:
			logger.Info(msg)
		case logging.LevelWarn:
			logger.Warn(msg)
		case logging.LevelError:
			logger.Error(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

// LoggerInterceptor returns a new UnaryServerInterceptor that logs the details of each gRPC call.
// It uses the provided logger to log events related to the gRPC calls. The interceptor logs
// only when the call finishes.
func LoggerInterceptor(logger Logger) grpc.UnaryServerInterceptor {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.FinishCall),
	}
	return logging.UnaryServerInterceptor(NewInterceptorLogger(logger), opts...)
}

// StreamLoggerInterceptor returns a new gRPC StreamServerInterceptor that logs the details of each streaming gRPC call.
// It uses the provided logger to record events related to streaming gRPC calls. The interceptor logs
// events when the stream is initiated and when it is completed.
func StreamLoggerInterceptor(logger Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		logger.Info("Stream started", zap.String("method", info.FullMethod))

		err := handler(srv, ss)

		if err != nil {
			logger.Error("Stream finished with error", zap.String("method", info.FullMethod), zap.Error(err))
		} else {
			logger.Info("Stream finished", zap.String("method", info.FullMethod))
		}

		return err
	}
}
