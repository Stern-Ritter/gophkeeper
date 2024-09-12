package server

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

func TestUserStorageCreate(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	userStorage := NewUserStorage(mock, l)

	user := model.User{
		Login:    "user",
		Password: "password",
	}

	mock.ExpectQuery(`
		INSERT INTO gophkeeper.users \(login, password\)
		VALUES \(@login, @password\)
		RETURNING id
	`).
		WithArgs(user.Login, user.Password).
		WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow("1"))

	createdUser, err := userStorage.Create(context.Background(), user)
	assert.NoError(t, err, "Error creating user")
	assert.Equal(t, "1", createdUser.ID, "Returned user ID does not match expected")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}

func TestUserStorageGetOneByLogin(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err, "Error init connection mock")
	defer mock.Close()

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	userStorage := NewUserStorage(mock, l)

	login := "user"
	expectedUser := model.User{
		ID:       "1",
		Login:    login,
		Password: "password",
	}

	mock.ExpectQuery(`
		SELECT id, login, password
		FROM gophkeeper.users
		WHERE login = @login
	`).
		WithArgs(login).
		WillReturnRows(pgxmock.NewRows([]string{"id", "login", "password"}).
			AddRow(expectedUser.ID, expectedUser.Login, expectedUser.Password))

	foundUser, err := userStorage.GetOneByLogin(context.Background(), login)
	assert.NoError(t, err, "Error getting user by login")
	assert.Equal(t, expectedUser, foundUser, "Returned user does not match expected")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "The expected SQL commands were not executed")
}
