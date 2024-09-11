package client

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestSignUp_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authClient := NewMockAuthServiceV1Client(mockCtrl)
	authClient.EXPECT().SignUp(gomock.Any(), &pb.SignUpRequestV1{
		Login:    "login",
		Password: "password",
	}).Return(&pb.SignUpResponseV1{Token: "token"}, nil)

	service := NewAuthService(authClient)
	token, err := service.SignUp("login", "password")

	assert.NoError(t, err, "Expected no error when signing up")
	assert.Equal(t, "token", token, "Expected correct token")
}

func TestSignUp_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authClient := NewMockAuthServiceV1Client(mockCtrl)
	authClient.EXPECT().SignUp(gomock.Any(), &pb.SignUpRequestV1{
		Login:    "login",
		Password: "password",
	}).Return(nil, errors.New("sign up error"))

	service := NewAuthService(authClient)
	token, err := service.SignUp("login", "password")

	assert.Error(t, err, "Expected error when signing up")
	assert.Equal(t, "", token, "Expected empty token on error")
}

func TestSignIn_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authClient := NewMockAuthServiceV1Client(mockCtrl)
	authClient.EXPECT().SignIn(gomock.Any(), &pb.SignInRequestV1{
		Login:    "login",
		Password: "password",
	}).Return(&pb.SignInResponseV1{Token: "token"}, nil)

	service := NewAuthService(authClient)
	token, err := service.SignIn("login", "password")

	assert.NoError(t, err, "Expected no error when signing in")
	assert.Equal(t, "token", token, "Expected correct token")
}

func TestSignIn_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authClient := NewMockAuthServiceV1Client(mockCtrl)
	authClient.EXPECT().SignIn(gomock.Any(), &pb.SignInRequestV1{
		Login:    "login",
		Password: "password",
	}).Return(nil, errors.New("sign in error"))

	service := NewAuthService(authClient)
	token, err := service.SignIn("login", "password")

	assert.Error(t, err, "Expected error when signing in")
	assert.Equal(t, "", token, "Expected empty token on error")
}
