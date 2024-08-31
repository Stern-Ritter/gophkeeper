package server

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"

	er "github.com/Stern-Ritter/gophkeeper/internal/errors"

	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

// SignUp registers a new user.
// If successful, it returns a SignUpResponse containing an authentication token.
func (s *Server) SignUp(ctx context.Context, in *pb.SignUpRequestV1) (*pb.SignUpResponseV1, error) {
	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req := model.MessageToSignUpRequest(in)
	token, err := s.AuthService.SignUp(ctx, req)
	if err != nil {
		var conflictErr er.ConflictError
		if errors.As(err, &conflictErr) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	resp := pb.SignUpResponseV1{
		Token: token,
	}

	return &resp, nil
}

// SignIn authenticates a user based on the provided credentials.
// If successful, it returns a SignInResponse containing an authentication token.
func (s *Server) SignIn(ctx context.Context, in *pb.SignInRequestV1) (*pb.SignInResponseV1, error) {
	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req := model.MessageToSignInRequest(in)
	token, err := s.AuthService.SignIn(ctx, req)

	if err != nil {
		var unauthorizedErr er.UnauthorizedError
		if errors.As(err, &unauthorizedErr) {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, "internal server error")
	}

	resp := pb.SignInResponseV1{
		Token: token,
	}

	return &resp, nil
}

// AddAccount add new account data for the authenticated user.
func (s *Server) AddAccount(ctx context.Context, in *pb.AddAccountRequestV1) (*pb.AddAccountResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req := model.AddAccountRequestToAccount(in)
	req.UserID = user.ID

	err = s.AccountService.CreateAccount(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.AddAccountResponseV1{}, nil
}

// DeleteAccount removes account data for the authenticated user.
func (s *Server) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequestV1) (*pb.DeleteAccountResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	accountID := in.Id
	err = s.AccountService.DeleteAccount(ctx, user.ID, accountID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.DeleteAccountResponseV1{}, nil
}

// GetAccounts get all accounts associated with the authenticated user.
func (s *Server) GetAccounts(ctx context.Context, in *pb.GetAccountsRequestV1) (*pb.GetAccountsResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	a, err := s.AccountService.GetAllAccounts(ctx, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	accounts := model.AccountsToRepeatedAccountMessage(a)
	resp := pb.GetAccountsResponseV1{
		Accounts: accounts,
	}

	return &resp, nil
}

// AddCard adds new card data for the authenticated user.
func (s *Server) AddCard(ctx context.Context, in *pb.AddCardRequestV1) (*pb.AddCardResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req := model.AddCardRequestToCard(in)
	req.UserID = user.ID

	err = s.CardService.CreateCard(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.AddCardResponseV1{}, nil
}

// DeleteCard removes card data for the authenticated user.
func (s *Server) DeleteCard(ctx context.Context, in *pb.DeleteCardRequestV1) (*pb.DeleteCardResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	cardID := in.Id
	err = s.CardService.DeleteCard(ctx, user.ID, cardID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.DeleteCardResponseV1{}, nil
}

// GetCards get all cards data associated with the authenticated user.
func (s *Server) GetCards(ctx context.Context, in *pb.GetCardsRequestV1) (*pb.GetCardsResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	c, err := s.CardService.GetAllCards(ctx, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	cards := model.CardsToRepeatedCardMessage(c)
	resp := pb.GetCardsResponseV1{
		Cards: cards,
	}

	return &resp, nil
}

// AddText adds new text data for the authenticated user.
func (s *Server) AddText(ctx context.Context, in *pb.AddTextRequestV1) (*pb.AddTextResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req := model.AddTextRequestToText(in)
	req.UserID = user.ID

	err = s.TextService.CreateText(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.AddTextResponseV1{}, nil
}

// DeleteText removes text data for the authenticated user.
func (s *Server) DeleteText(ctx context.Context, in *pb.DeleteTextRequestV1) (*pb.DeleteTextResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	textID := in.Id
	err = s.TextService.DeleteText(ctx, user.ID, textID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.DeleteTextResponseV1{}, nil
}

// GetTexts get all text data associated with the authenticated user.
func (s *Server) GetTexts(ctx context.Context, in *pb.GetTextsRequestV1) (*pb.GetTextsResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	t, err := s.TextService.GetAllTexts(ctx, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	texts := model.TextsToRepeatedTextMessage(t)
	resp := pb.GetTextsResponseV1{
		Texts: texts,
	}

	return &resp, nil
}

// UploadFile handles the file upload to the user storage process for the authenticated user with gRPC stream.
func (s *Server) UploadFile(stream pb.FileServiceV1_UploadFileServer) error {
	ctx := stream.Context()
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	err = s.FileService.UploadFile(ctx, user.ID, stream)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return stream.SendAndClose(&pb.UploadFileResponseV1{})
}

// DownloadFile handles the file download from the user storage process for the authenticated user with gRPC stream.
func (s *Server) DownloadFile(in *pb.DownloadFileRequestV1, stream pb.FileServiceV1_DownloadFileServer) error {
	ctx := stream.Context()
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	err = s.FileService.DownloadFile(ctx, user.ID, in.GetId(), stream)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return status.Error(codes.PermissionDenied, err.Error())
		default:
			return status.Error(codes.Internal, "Internal server error")
		}
	}

	return nil
}

// DeleteFile deletes a file for the authenticated user.
func (s *Server) DeleteFile(ctx context.Context, in *pb.DeleteFileRequestV1) (*pb.DeleteFileResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if err := s.Validator.Validate(in); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	fileID := in.Id
	err = s.FileService.DeleteFile(ctx, user.ID, fileID)
	if err != nil {
		var notFoundErr er.NotFoundError
		var forbiddenErr er.ForbiddenError
		switch {
		case errors.As(err, &notFoundErr):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.As(err, &forbiddenErr):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, "internal server error")
		}
	}

	return &pb.DeleteFileResponseV1{}, nil
}

// GetFiles get all files associated with the authenticated user.
func (s *Server) GetFiles(ctx context.Context, in *pb.GetFilesRequestV1) (*pb.GetFilesResponseV1, error) {
	user, err := s.UserService.GetCurrentUser(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	f, err := s.FileService.GetAllFiles(ctx, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	files := model.FilesToRepeatedFileMessage(f)
	resp := pb.GetFilesResponseV1{
		Files: files,
	}

	return &resp, nil
}
