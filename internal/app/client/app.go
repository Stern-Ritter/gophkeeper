package client

import (
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
	crypto "github.com/Stern-Ritter/gophkeeper/internal/crypto/client"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/client"
	service "github.com/Stern-Ritter/gophkeeper/internal/service/client"
	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func Run(client service.Client, app service.Application, cfg *config.ClientConfig, logger *logger.ClientLogger) error {
	opts := make([]grpc.DialOption, 0)
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	opts = append(opts, grpc.WithUnaryInterceptor(client.AuthInterceptor))
	opts = append(opts, grpc.WithStreamInterceptor(client.AuthStreamInterceptor))
	creds, err := crypto.GetTransportCredentials()
	if err != nil {
		logger.Fatal(err.Error(), zap.String("event", "load credentials"))
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.Dial(cfg.ServerURL, opts...)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to connect to server: %s", cfg.ServerURL),
			zap.String("event", "connect server"), zap.Error(err))
	}
	defer conn.Close()

	authService := service.NewAuthService(pb.NewAuthServiceV1Client(conn))
	accountService := service.NewAccountService(pb.NewAccountServiceV1Client(conn))
	cardService := service.NewCardService(pb.NewCardServiceV1Client(conn))
	textService := service.NewTextService(pb.NewTextServiceV1Client(conn))
	fileService := service.NewFileService(pb.NewFileServiceV1Client(conn))

	client.SetAuthService(authService)
	client.SetAccountService(accountService)
	client.SetCardService(cardService)
	client.SetTextService(textService)
	client.SetFileService(fileService)
	client.SetApp(app)

	client.AuthView()
	if err := app.Run(); err != nil {
		logger.Error(fmt.Sprintf("Failed to run application: %s", cfg.ServerURL),
			zap.String("event", "run application"), zap.Error(err))
		return err
	}

	return nil
}
