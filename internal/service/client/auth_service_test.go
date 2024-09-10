package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestSignUp(t *testing.T) {
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

func TestSignIn(t *testing.T) {
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
