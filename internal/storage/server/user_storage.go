package server

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

// UserStorage defines the interface for user storage operations.
type UserStorage interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	GetOneByLogin(ctx context.Context, login string) (model.User, error)
}

// UserStorageImpl is an implementation of the UserStorage interface.
// It uses a PgxIface for database operations and a Logger for logging errors.
type UserStorageImpl struct {
	db     PgxIface
	Logger *logger.ServerLogger
}

// NewUserStorage creates a new instance of UserStorage.
func NewUserStorage(db PgxIface, logger *logger.ServerLogger) UserStorage {
	return &UserStorageImpl{
		db:     db,
		Logger: logger,
	}
}

// Create inserts a new user into the database.
func (s *UserStorageImpl) Create(ctx context.Context, user model.User) (model.User, error) {
	var userID string
	row := s.db.QueryRow(ctx, `
		INSERT INTO gophkeeper.users (login, password)
		VALUES (@login, @password)
		RETURNING id
	`, pgx.NamedArgs{
		"login":    user.Login,
		"password": user.Password,
	})

	err := row.Scan(&userID)
	if err != nil {
		s.Logger.Debug("Failed to insert user", zap.String("event", "create user"),
			zap.String("user", fmt.Sprintf("%v", user)), zap.Error(err))
		return model.User{}, err
	}

	user.ID = userID
	return user, nil
}

// GetOneByLogin get a user by login from the database.
func (s *UserStorageImpl) GetOneByLogin(ctx context.Context, login string) (model.User, error) {
	row := s.db.QueryRow(ctx, `
		SELECT id, login, password
		FROM gophkeeper.users
		WHERE login = @login
	`, pgx.NamedArgs{
		"login": login,
	})

	user := model.User{}
	err := row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		s.Logger.Debug("Failed select user by login", zap.String("event", "get user by login"),
			zap.String("login", login), zap.Error(err))
		return model.User{}, err
	}

	return user, nil
}
