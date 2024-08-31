package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bufbuild/protovalidate-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"

	_ "github.com/jackc/pgx/v5/stdlib"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/server"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	service "github.com/Stern-Ritter/gophkeeper/internal/service/server"
	storage "github.com/Stern-Ritter/gophkeeper/internal/storage/server"
	"github.com/Stern-Ritter/gophkeeper/migrations"
)

func Run(cfg *config.ServerConfig, logger *logger.ServerLogger) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	idleConnsClosed := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := pgxpool.New(ctx, cfg.DatabaseDSN)
	if err != nil {
		logger.Fatal("Failed to create database connection", zap.String("event", "create database connection"),
			zap.String("database url", cfg.DatabaseDSN), zap.Error(err))
	}
	err = db.Ping(ctx)
	if err != nil {
		logger.Fatal("Failed to connect database", zap.String("event", "connect database"),
			zap.String("database url", cfg.DatabaseDSN), zap.Error(err))
	}
	err = migrateDatabase(cfg.DatabaseDSN)
	if err != nil {
		logger.Fatal("Failed to migrate database", zap.String("event", "migrate database"),
			zap.String("database url", cfg.DatabaseDSN), zap.Error(err))
	}

	userStorage := storage.NewUserStorage(db, logger)
	dataStorage := storage.NewDataStorage(db, logger)
	fileStorage := storage.NewFileStorage(db, logger)

	userService := service.NewUserService(userStorage, logger)
	authService := service.NewAuthService(userService, cfg.AuthenticationKey, logger)
	accountService := service.NewAccountService(dataStorage, cfg.EncryptionKey, logger)
	cardService := service.NewCardService(dataStorage, cfg.EncryptionKey, logger)
	textService := service.NewTextService(dataStorage, cfg.EncryptionKey, logger)
	fileService := service.NewFileService(fileStorage, cfg.FileStoragePath, logger)

	validator, err := protovalidate.New()
	if err != nil {
		logger.Fatal("Failed to initialize validator", zap.String("event", "initialize validator"),
			zap.Error(err))
	}

	server := service.NewServer(userService, authService, accountService, cardService, textService, fileService,
		validator, cfg, logger)

	err = runGrpcServer(server, signals, idleConnsClosed)
	return err
}

func migrateDatabase(databaseDsn string) error {
	goose.SetBaseFS(migrations.Migrations)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose failed to set postgres dialect: %w", err)
	}

	db, err := goose.OpenDBWithDriver("pgx", databaseDsn)
	if err != nil {
		return fmt.Errorf("goose failed to open database connection: %w", err)
	}

	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("goose failed to migrate database: %w", err)
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("goose failed to close database connection: %w", err)
	}

	return nil
}

func runGrpcServer(server *service.Server, signals chan os.Signal, idleConnsClosed chan struct{}) error {
	listen, err := net.Listen("tcp", server.Config.URL)
	if err != nil {
		return err
	}

	opts := make([]grpc.ServerOption, 0)
	opts = append(opts, grpc.ChainUnaryInterceptor(logger.LoggerInterceptor(server.Logger)))
	opts = append(opts, grpc.ChainUnaryInterceptor(server.AuthInterceptor))
	opts = append(opts, grpc.ChainStreamInterceptor(logger.StreamLoggerInterceptor(server.Logger)))
	opts = append(opts, grpc.ChainStreamInterceptor(server.AuthStreamInterceptor))
	creds, err := credentials.NewServerTLSFromFile(server.Config.TLSCertPath, server.Config.TLSKeyPath)
	if err != nil {
		server.Logger.Fatal(err.Error(), zap.String("event", "load credentials"))
	}
	opts = append(opts, grpc.Creds(creds))

	srv := grpc.NewServer(opts...)
	pb.RegisterAuthServiceV1Server(srv, server)
	pb.RegisterAccountServiceV1Server(srv, server)
	pb.RegisterCardServiceV1Server(srv, server)
	pb.RegisterTextServiceV1Server(srv, server)
	pb.RegisterFileServiceV1Server(srv, server)

	go func() {
		<-signals

		server.Logger.Info("Shutting down server", zap.String("event", "shutdown server"))
		srv.GracefulStop()

		close(idleConnsClosed)
	}()

	server.Logger.Info("Starting server", zap.String("event", "start server"),
		zap.String("url", server.Config.URL))
	if err := srv.Serve(listen); err != nil {
		return err
	}

	<-idleConnsClosed
	server.Logger.Info("Server shutdown complete", zap.String("event", "shutdown server"))

	return nil
}
