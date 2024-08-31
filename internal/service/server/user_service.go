package server

import (
	"context"

	"github.com/golang-jwt/jwt/v4"

	er "github.com/Stern-Ritter/gophkeeper/internal/errors"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
	storage "github.com/Stern-Ritter/gophkeeper/internal/storage/server"
)

// UserService defines the interface for operations with users.
type UserService interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByLogin(ctx context.Context, login string) (model.User, error)
	GetCurrentUser(ctx context.Context) (model.User, error)
}

// UserServiceImpl is an implementation of the UserService interface.
type UserServiceImpl struct {
	userStorage storage.UserStorage
	logger      *logger.ServerLogger
}

// NewUserService creates a new instance of UserService.
func NewUserService(userStorage storage.UserStorage, logger *logger.ServerLogger) UserService {
	return &UserServiceImpl{
		userStorage: userStorage,
		logger:      logger,
	}
}

// CreateUser creates a new user using the provided user model.
func (s *UserServiceImpl) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return s.userStorage.Create(ctx, user)
}

// GetUserByLogin get a user by login.
func (s *UserServiceImpl) GetUserByLogin(ctx context.Context, login string) (model.User, error) {
	return s.userStorage.GetOneByLogin(ctx, login)
}

// GetCurrentUser get the currently authenticated user based on the context.
// It extracts the login from the JWT claims in the context and get the user from storage.
// Returns the current user or an unauthorized error if the user cannot be determined.
func (s *UserServiceImpl) GetCurrentUser(ctx context.Context) (model.User, error) {
	claims, ok := ctx.Value(AuthorizationTokenContextKey).(jwt.MapClaims)
	if !ok {
		return model.User{}, er.NewUnauthorizedError("user is not authorized to access this resource", nil)
	}

	if _, ok := claims["login"]; !ok {
		return model.User{}, er.NewUnauthorizedError("user is not authorized to access this resource", nil)
	}

	login, ok := claims["login"].(string)
	if !ok {
		return model.User{}, er.NewUnauthorizedError("user is not authorized to access this resource", nil)
	}

	currentUser, err := s.userStorage.GetOneByLogin(ctx, login)
	if err != nil {
		return model.User{}, er.NewUnauthorizedError("user is not authorized to access this resource", nil)
	}
	return currentUser, err
}
