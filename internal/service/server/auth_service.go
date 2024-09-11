package server

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"

	"github.com/Stern-Ritter/gophkeeper/internal/auth"
	er "github.com/Stern-Ritter/gophkeeper/internal/errors"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

const (
	AuthTokenExpiration          = time.Hour * 24 // AuthTokenExpiration defines the duration for which an authentication token is valid.
	UserLoginUniqueConstrainName = "users_login_unique"
)

// AuthService provides an interface for user authentication operations.
type AuthService interface {
	SignUp(ctx context.Context, req model.SignUpRequest) (string, error)
	SignIn(ctx context.Context, req model.SignInRequest) (string, error)
}

// AuthServiceImpl is an implementation of the AuthService interface.
type AuthServiceImpl struct {
	userService   UserService
	authSecretKey string
	logger        logger.ServerLogger
}

// NewAuthService creates a new instance of AuthService.
func NewAuthService(userService UserService, authSecretKey string, logger logger.ServerLogger) AuthService {
	return &AuthServiceImpl{
		userService:   userService,
		authSecretKey: authSecretKey,
		logger:        logger,
	}
}

// SignUp registers a new user based on the provided sign-up request. It hashes the user's password and generates a JWT token.
// Returns the JWT token or an error if the registration fails.
func (s *AuthServiceImpl) SignUp(ctx context.Context, req model.SignUpRequest) (string, error) {
	user := model.SignUpRequestToUser(req)

	passwordHash, err := auth.GetPasswordHash(user.Password)
	if err != nil {
		s.logger.Error("Failed to generate password hash", zap.String("event", "user registration"),
			zap.Error(err))
		return "", fmt.Errorf("failed to generate password hash: %w", err)
	}

	user.Password = passwordHash

	savedUser, err := s.userService.CreateUser(ctx, user)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.ConstraintName == UserLoginUniqueConstrainName {
			s.logger.Info("User already exists", zap.String("event", "user registration"),
				zap.String("login", user.Login), zap.Error(err))
			return "", er.NewConflictError(fmt.Sprintf("user with login: %s already exists", user.Login), err)
		}

		return "", err
	}

	token, err := auth.NewToken(savedUser, s.authSecretKey, AuthTokenExpiration)
	if err != nil {
		s.logger.Error("Failed to generate jwt token", zap.String("event", "user registration"),
			zap.String("login", user.Login), zap.Error(err))
		return "", fmt.Errorf("failed to generate jwt token: %w", err)
	}

	return token, nil
}

// SignIn authenticates a user based on the provided sign-in request. It verifies the user's login and password,
// and if successful, generates a JWT token.
// Returns the JWT token or an error if the authentication fails.
func (s *AuthServiceImpl) SignIn(ctx context.Context, req model.SignInRequest) (string, error) {
	user, err := s.userService.GetUserByLogin(ctx, req.Login)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("User not exists", zap.String("event", "user authentication"),
				zap.String("login", req.Login), zap.Error(err))
			return "", er.NewUnauthorizedError("invalid login or password", err)
		default:
			s.logger.Error("Failed to get user by login", zap.String("event", "user authentication"),
				zap.String("login", req.Login), zap.Error(err))
			return "", fmt.Errorf("failed to get user by login: %w", err)
		}
	}

	if !auth.CheckPasswordHash(req.Password, user.Password) {
		s.logger.Info("Invalid password", zap.String("event", "user authentication"),
			zap.String("login", req.Login), zap.Error(err))
		return "", er.NewUnauthorizedError("invalid login or password", err)
	}

	token, err := auth.NewToken(user, s.authSecretKey, AuthTokenExpiration)
	if err != nil {
		s.logger.Error("Failed to generate jwt token", zap.String("event", "user authentication"),
			zap.String("login", req.Login), zap.Error(err))
		return "", fmt.Errorf("failed to generate jwt token: %w", err)
	}

	return token, nil
}
