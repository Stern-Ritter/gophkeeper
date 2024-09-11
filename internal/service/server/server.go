package server

import (
	"github.com/bufbuild/protovalidate-go"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/server"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

// Server represents the gRPC server that handles services in the application.
type Server struct {
	UserService    UserService
	AuthService    AuthService
	AccountService AccountService
	CardService    CardService
	TextService    TextService
	FileService    FileService
	Validator      *protovalidate.Validator
	Config         *config.ServerConfig
	Logger         logger.ServerLogger
	pb.UnimplementedAuthServiceV1Server
	pb.UnimplementedAccountServiceV1Server
	pb.UnimplementedCardServiceV1Server
	pb.UnimplementedTextServiceV1Server
	pb.UnimplementedFileServiceV1Server
}

// NewServer creates a new instance of the Server
func NewServer(userService UserService, authService AuthService, accountService AccountService, cardService CardService,
	textService TextService, fileService FileService, validator *protovalidate.Validator, serverConfig *config.ServerConfig,
	logger logger.ServerLogger) *Server {
	return &Server{
		UserService:    userService,
		AuthService:    authService,
		AccountService: accountService,
		CardService:    cardService,
		TextService:    textService,
		FileService:    fileService,
		Validator:      validator,
		Config:         serverConfig,
		Logger:         logger,
	}
}
