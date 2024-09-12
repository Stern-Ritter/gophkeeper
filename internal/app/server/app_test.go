package server

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/server"
	service "github.com/Stern-Ritter/gophkeeper/internal/service/server"
)

func TestRunGrpcServer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockLogger := NewMockServerLogger(mockCtrl)

	mockLogger.EXPECT().Info("Starting server", zap.String("event", "start server"),
		zap.String("url", "localhost:8080")).Times(1)
	mockLogger.EXPECT().Info("Shutting down server", zap.String("event", "shutdown server")).Times(1)
	mockLogger.EXPECT().Info("Server shutdown complete", zap.String("event", "shutdown server")).Times(1)

	cfg := &config.ServerConfig{
		URL:             "localhost:8080",
		TLSCertPath:     "../../../testdata/certs/server-cert.pem",
		TLSKeyPath:      "../../../testdata/certs/server-key.pem",
		DatabaseDSN:     "postgresql://user:password@localhost/dbname",
		EncryptionKey:   "encryption-key",
		FileStoragePath: "/path-to-storage",
	}

	server := &service.Server{
		Config: cfg,
		Logger: mockLogger,
	}

	signals := make(chan os.Signal, 1)
	idleConnsClosed := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		signals <- os.Interrupt
	}()

	err := runGrpcServer(server, signals, idleConnsClosed)
	require.NoError(t, err, "Expected no error when running gRPC server")
}
