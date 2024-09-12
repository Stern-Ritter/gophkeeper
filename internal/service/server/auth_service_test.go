package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/Stern-Ritter/gophkeeper/internal/auth"
	e "github.com/Stern-Ritter/gophkeeper/internal/errors"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

func TestAuthServiceSignUp_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	authSecretKey := "secret"
	authService := NewAuthService(mockUserService, authSecretKey, l)

	signUpRequest := model.SignUpRequest{
		Login:    "user",
		Password: "password",
	}

	hashedPassword, err := auth.GetPasswordHash(signUpRequest.Password)
	require.NoError(t, err, "unexpected password hashing error")
	savedUser := model.User{
		ID:       "1",
		Login:    "user",
		Password: hashedPassword,
	}

	mockUserService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, user model.User) (model.User, error) {
			assert.NotEqual(t, signUpRequest.Password, user.Password, "password should be hashed")
			return savedUser, nil
		})

	token, err := authService.SignUp(context.Background(), signUpRequest)
	assert.NoError(t, err, "expected no error while sign up")

	_, err = auth.ValidateToken(token, authSecretKey)
	assert.NoError(t, err, "expected returned JWT token to be valid")
}

func TestAuthServiceSignUp_UserAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	authSecretKey := "secret"
	authService := NewAuthService(mockUserService, authSecretKey, l)

	signUpRequest := model.SignUpRequest{
		Login:    "user",
		Password: "password",
	}

	pgErr := &pgconn.PgError{ConstraintName: UserLoginUniqueConstrainName}
	mockUserService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(model.User{}, pgErr)

	token, err := authService.SignUp(context.Background(), signUpRequest)

	assert.Error(t, err, "expected error when user already exist")
	assert.Empty(t, token, "expected no token when user already exists")
	assert.ErrorAs(t, err, &e.ConflictError{}, "expected conflict error")
}

func TestAuthServiceSignUp_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	authSecretKey := "secret"
	authService := NewAuthService(mockUserService, authSecretKey, l)

	signUpRequest := model.SignUpRequest{
		Login:    "user",
		Password: "password",
	}

	mockUserService.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(model.User{}, fmt.Errorf("internal error"))

	token, err := authService.SignUp(context.Background(), signUpRequest)

	assert.Error(t, err, "expected internal error during user creation")
	assert.Empty(t, token, "expected no token due to internal error")
}

func TestAuthServiceSignIn_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	authSecretKey := "secret"
	authService := NewAuthService(mockUserService, authSecretKey, l)

	signInRequest := model.SignInRequest{
		Login:    "user",
		Password: "password",
	}

	hashedPassword, err := auth.GetPasswordHash(signInRequest.Password)
	require.NoError(t, err, "unexpected password hashing error")
	user := model.User{
		ID:       "1",
		Login:    "user",
		Password: hashedPassword,
	}

	mockUserService.EXPECT().GetUserByLogin(gomock.Any(), signInRequest.Login).Return(user, nil)

	token, err := authService.SignIn(context.Background(), signInRequest)
	assert.NoError(t, err, "expected no error during sign in")

	_, err = auth.ValidateToken(token, authSecretKey)
	assert.NoError(t, err, "expected returned JWT token to be valid")
}

func TestAuthServiceSignIn_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	authSecretKey := "secret"
	authService := NewAuthService(mockUserService, authSecretKey, l)

	signInRequest := model.SignInRequest{
		Login:    "user",
		Password: "password",
	}

	mockUserService.EXPECT().GetUserByLogin(gomock.Any(), signInRequest.Login).Return(model.User{}, pgx.ErrNoRows)

	token, err := authService.SignIn(context.Background(), signInRequest)

	assert.Error(t, err, "expected error when user not found")
	assert.Empty(t, token, "expected no token when user not found")
	assert.ErrorAs(t, err, &e.UnauthorizedError{}, "expected unauthorized error")
}

func TestAuthServiceSignIn_InvalidPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	authSecretKey := "secret"
	authService := NewAuthService(mockUserService, authSecretKey, l)

	signInRequest := model.SignInRequest{
		Login:    "user",
		Password: "wrong password",
	}

	hashedPassword, err := auth.GetPasswordHash("password")
	require.NoError(t, err, "unexpected password hashing error")
	user := model.User{
		ID:       "1",
		Login:    "user",
		Password: hashedPassword,
	}

	mockUserService.EXPECT().GetUserByLogin(gomock.Any(), signInRequest.Login).Return(user, nil)

	token, err := authService.SignIn(context.Background(), signInRequest)

	assert.Error(t, err, "expected error when password is invalid")
	assert.Empty(t, token, "expected no token when password is invalid")
	assert.ErrorAs(t, err, &e.UnauthorizedError{}, "expected unauthorized error")
}

func TestAuthServiceSignIn_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	authSecretKey := "secret"
	authService := NewAuthService(mockUserService, authSecretKey, l)

	signInRequest := model.SignInRequest{
		Login:    "user",
		Password: "password",
	}

	mockUserService.EXPECT().GetUserByLogin(gomock.Any(), signInRequest.Login).Return(model.User{}, fmt.Errorf("internal error"))

	token, err := authService.SignIn(context.Background(), signInRequest)

	assert.Error(t, err, "expected internal error during user retrieval")
	assert.Empty(t, token, "expected no token due to internal error")
}
