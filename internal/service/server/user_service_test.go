package server

import (
	"context"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

func TestUserServiceCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStorage := NewMockUserStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	userService := NewUserService(mockUserStorage, l)

	user := model.User{
		Login:    "user",
		Password: "password",
	}

	savedUser := model.User{
		ID:       "1",
		Login:    "user",
		Password: "password hash",
	}

	mockUserStorage.EXPECT().Create(gomock.Any(), user).Return(savedUser, nil)

	createdUser, err := userService.CreateUser(context.Background(), user)

	assert.NoError(t, err, "unexpected error creating user")
	assert.Equal(t, savedUser, createdUser, "returned created user does not match expected user")
}

func TestUserServiceGetUserByLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStorage := NewMockUserStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	userService := NewUserService(mockUserStorage, l)

	login := "user"
	user := model.User{
		ID:       "1",
		Login:    login,
		Password: "password",
	}

	mockUserStorage.EXPECT().GetOneByLogin(gomock.Any(), login).Return(user, nil)

	savedUser, err := userService.GetUserByLogin(context.Background(), login)

	assert.NoError(t, err, "unexpected error getting user by login")
	assert.Equal(t, user, savedUser, "Fetched user does not match expected user")
}

func TestUserServiceGetCurrentUserSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStorage := NewMockUserStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	userService := NewUserService(mockUserStorage, l)

	login := "user"
	user := model.User{
		ID:       "1",
		Login:    login,
		Password: "password",
	}

	claims := jwt.MapClaims{
		"login": login,
	}
	ctx := context.WithValue(context.Background(), AuthorizationTokenContextKey, claims)

	mockUserStorage.EXPECT().GetOneByLogin(gomock.Any(), login).Return(user, nil)

	currentUser, err := userService.GetCurrentUser(ctx)

	assert.NoError(t, err, "Error while getting current user")
	assert.Equal(t, user, currentUser, "Current user does not match expected user")
}

func TestUserServiceGetCurrentUserUnauthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStorage := NewMockUserStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	userService := NewUserService(mockUserStorage, l)

	ctx := context.Background()

	_, err = userService.GetCurrentUser(ctx)

	assert.Error(t, err, "Expected unauthorized error when JWT claims are missing")
}
