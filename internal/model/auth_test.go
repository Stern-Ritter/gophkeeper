package model

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestMessageToSignUpRequest(t *testing.T) {
	req := &pb.SignUpRequestV1{
		Login:    "signup request login",
		Password: "signup request password",
	}

	signUpRequest := MessageToSignUpRequest(req)

	assert.Equal(t, "signup request login", signUpRequest.Login)
	assert.Equal(t, "signup request password", signUpRequest.Password)
}

func TestMessageToSignInRequest(t *testing.T) {
	req := &pb.SignInRequestV1{
		Login:    "signin request login",
		Password: "signin request password",
	}

	signInRequest := MessageToSignInRequest(req)

	assert.Equal(t, "signin request login", signInRequest.Login)
	assert.Equal(t, "signin request password", signInRequest.Password)
}

func TestSignUpRequestToUser(t *testing.T) {
	signUpRequest := SignUpRequest{
		Login:    "signup login",
		Password: "signup password",
	}

	user := SignUpRequestToUser(signUpRequest)

	assert.Equal(t, "signup login", user.Login)
	assert.Equal(t, "signup password", user.Password)
}
