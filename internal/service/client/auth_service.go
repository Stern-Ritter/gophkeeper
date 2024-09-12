package client

import (
	"context"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

// AuthService defines an interface for user authentication.
type AuthService interface {
	SignUp(login string, password string) (string, error)
	SignIn(login string, password string) (string, error)
}

// AuthServiceImpl is implementation of the AuthService interface.
type AuthServiceImpl struct {
	authClient pb.AuthServiceV1Client
}

// NewAuthService creates a new instance of AuthService.
func NewAuthService(authClient pb.AuthServiceV1Client) AuthService {
	return &AuthServiceImpl{
		authClient: authClient,
	}
}

// SignUp registers a new user with the given login and password by sending a request
// to the authentication service with gRPC. If successful, it returns an authentication token.
func (s *AuthServiceImpl) SignUp(login string, password string) (string, error) {
	ctx := context.Background()
	req := &pb.SignUpRequestV1{Login: login, Password: password}
	resp, err := s.authClient.SignUp(ctx, req)
	if err != nil {
		return "", err
	}

	token := resp.Token
	return token, nil
}

// SignIn authenticates a user with the given login and password by sending a request
// to the authentication service with gRPC. If successful, it returns an authentication token.
func (s *AuthServiceImpl) SignIn(login string, password string) (string, error) {
	ctx := context.Background()
	req := pb.SignInRequestV1{Login: login, Password: password}
	resp, err := s.authClient.SignIn(ctx, &req)
	if err != nil {
		return "", err
	}

	token := resp.Token
	return token, nil
}
