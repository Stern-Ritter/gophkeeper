package model

import (
	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

type SignUpRequest struct {
	Login    string
	Password string
}

type SignInRequest struct {
	Login    string
	Password string
}

func MessageToSignUpRequest(req *pb.SignUpRequestV1) SignUpRequest {
	return SignUpRequest{
		Login:    req.Login,
		Password: req.Password,
	}
}

func MessageToSignInRequest(req *pb.SignInRequestV1) SignInRequest {
	return SignInRequest{
		Login:    req.Login,
		Password: req.Password,
	}
}

func SignUpRequestToUser(req SignUpRequest) User {
	return User{
		Login:    req.Login,
		Password: req.Password,
	}
}
