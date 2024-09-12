package server

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/server"
	er "github.com/Stern-Ritter/gophkeeper/internal/errors"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name                 string
		req                  *pb.SignUpRequestV1
		expectedResp         *pb.SignUpResponseV1
		authServiceSignUpErr error
		expectedErr          error
	}{
		{
			name: "should return conflict error if user already exists",
			req: &pb.SignUpRequestV1{
				Login:    "user",
				Password: "password",
			},
			authServiceSignUpErr: er.NewConflictError("user already exists", nil),
			expectedErr:          status.Error(codes.AlreadyExists, "user already exists"),
		},
		{
			name: "should sign up successfully",
			req: &pb.SignUpRequestV1{
				Login:    "user",
				Password: "password",
			},
			authServiceSignUpErr: nil,
			expectedResp: &pb.SignUpResponseV1{
				Token: "token",
			},
			expectedErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockAuthService.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return("token", tt.authServiceSignUpErr).Times(1)

			resp, err := s.SignUp(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name                 string
		req                  *pb.SignInRequestV1
		expectedResp         *pb.SignInResponseV1
		authServiceSignInErr error
		expectedErr          error
	}{
		{
			name: "should return unauthorized error if credentials are incorrect",
			req: &pb.SignInRequestV1{
				Login:    "user",
				Password: "password",
			},
			authServiceSignInErr: er.NewUnauthorizedError("unauthorized", nil),
			expectedErr:          status.Error(codes.Unauthenticated, "unauthorized"),
		},
		{
			name: "should sign in successfully",
			req: &pb.SignInRequestV1{
				Login:    "user",
				Password: "password",
			},
			authServiceSignInErr: nil,
			expectedResp: &pb.SignInResponseV1{
				Token: "token",
			},
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockAuthService.EXPECT().SignIn(gomock.Any(), gomock.Any()).Return("token", tc.authServiceSignInErr).Times(1)

			resp, err := s.SignIn(context.Background(), tc.req)
			if tc.expectedErr != nil {
				assert.ErrorIs(t, tc.expectedErr, err, "should return error: %s, got: %s", tc.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tc.expectedResp, resp, "should return response: %v, got: %v", tc.expectedResp, resp)
			}
		})
	}
}

func TestAddAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.AddAccountRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callCreateAccount  bool
		createAccountErr   error
		expectedResp       *pb.AddAccountResponseV1
		expectedErr        error
	}{
		{
			name: "should return unauthenticated error if authentication fails",
			req: &pb.AddAccountRequestV1{
				Login:    "login",
				Password: "password",
				Comment:  "comment",
			},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name: "should return invalid argument error if request is invalid",
			req: &pb.AddAccountRequestV1{
				Login:    "",
				Password: "password",
				Comment:  "comment",
			},
			callGetCurrentUser: true,
			expectedErr:        status.Error(codes.InvalidArgument, "validation error:\n - login: Login should not be empty [login]"),
		},
		{
			name: "should return internal error if account service return error",
			req: &pb.AddAccountRequestV1{
				Login:    "login",
				Password: "password",
				Comment:  "comment",
			},
			callGetCurrentUser: true,
			callCreateAccount:  true,
			createAccountErr:   errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name: "should add account successfully",
			req: &pb.AddAccountRequestV1{
				Login:    "login",
				Password: "password",
				Comment:  "comment",
			},
			callGetCurrentUser: true,
			callCreateAccount:  true,
			expectedResp:       &pb.AddAccountResponseV1{},
			expectedErr:        nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callCreateAccount {
				mockAccountService.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(tt.createAccountErr)
			}

			resp, err := s.AddAccount(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestDeleteAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.DeleteAccountRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callDeleteAccount  bool
		deleteAccountErr   error
		expectedResp       *pb.DeleteAccountResponseV1
		expectedErr        error
	}{
		{
			name: "should return unauthenticated error if authentication fails",
			req: &pb.DeleteAccountRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name: "should return invalid argument error if request is invalid",
			req: &pb.DeleteAccountRequestV1{
				Id: "",
			},
			callGetCurrentUser: true,
			expectedErr:        status.Error(codes.InvalidArgument, "validation error:\n - id: Deleted account data ID should not be empty [id]"),
		},
		{
			name: "should return not found error if account does not exist",
			req: &pb.DeleteAccountRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteAccount:  true,
			deleteAccountErr:   er.NewNotFoundError("account not found", nil),
			expectedErr:        status.Error(codes.NotFound, "account not found"),
		},
		{
			name: "should return permission denied error if user does not have permission",
			req: &pb.DeleteAccountRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteAccount:  true,
			deleteAccountErr:   er.NewForbiddenError("access denied", nil),
			expectedErr:        status.Error(codes.PermissionDenied, "access denied"),
		},
		{
			name: "should return internal error if account service return error",
			req: &pb.DeleteAccountRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteAccount:  true,
			deleteAccountErr:   errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name: "should delete account successfully",
			req: &pb.DeleteAccountRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteAccount:  true,
			expectedResp:       &pb.DeleteAccountResponseV1{},
			expectedErr:        nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callDeleteAccount {
				mockAccountService.EXPECT().DeleteAccount(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.deleteAccountErr)
			}

			resp, err := s.DeleteAccount(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestGetAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.GetAccountsRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetAccounts    bool
		getAllAccountsErr  error
		accounts           []model.Account
		expectedResp       *pb.GetAccountsResponseV1
		expectedErr        error
	}{
		{
			name:               "should return unauthenticated error if authentication fails",
			req:                &pb.GetAccountsRequestV1{},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name:               "should return internal error if account service return error",
			req:                &pb.GetAccountsRequestV1{},
			callGetCurrentUser: true,
			callGetAccounts:    true,
			getAllAccountsErr:  errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name:               "should return accounts successfully",
			req:                &pb.GetAccountsRequestV1{},
			callGetCurrentUser: true,
			callGetAccounts:    true,
			accounts: []model.Account{
				{ID: "1", UserID: "1", Login: "login 1", Password: "password 1", Comment: "comment 1"},
				{ID: "2", UserID: "1", Login: "login 2", Password: "password 2", Comment: "comment 2"},
			},
			expectedResp: &pb.GetAccountsResponseV1{
				Accounts: []*pb.AccountV1{
					{Id: "1", UserId: "1", Login: "login 1", Password: "password 1", Comment: "comment 1"},
					{Id: "2", UserId: "1", Login: "login 2", Password: "password 2", Comment: "comment 2"},
				},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callGetAccounts {
				mockAccountService.EXPECT().GetAllAccounts(gomock.Any(), "1").Return(tt.accounts, tt.getAllAccountsErr)
			}

			resp, err := s.GetAccounts(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestAddCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.AddCardRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callCreateCard     bool
		createCardErr      error
		expectedResp       *pb.AddCardResponseV1
		expectedErr        error
	}{
		{
			name: "should return unauthenticated error if authentication fails",
			req: &pb.AddCardRequestV1{
				Number:  "1234-5678-1234-5678",
				Owner:   "John Doe",
				Expiry:  "12/25",
				Cvc:     "123",
				Pin:     "1234",
				Comment: "comment",
			},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name: "should return invalid argument error if request is invalid",
			req: &pb.AddCardRequestV1{
				Number:  "1234-5678-1234-5678",
				Owner:   "",
				Expiry:  "12/25",
				Cvc:     "123",
				Pin:     "1234",
				Comment: "comment",
			},
			callGetCurrentUser: true,
			expectedErr:        status.Error(codes.InvalidArgument, "validation error:\n - owner: Owner must not be empty [owner]"),
		},
		{
			name: "should return internal error if card service return error",
			req: &pb.AddCardRequestV1{
				Number:  "1234-5678-1234-5678",
				Owner:   "John Doe",
				Expiry:  "12/25",
				Cvc:     "123",
				Pin:     "1234",
				Comment: "comment",
			},
			callGetCurrentUser: true,
			callCreateCard:     true,
			createCardErr:      errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name: "should add card successfully",
			req: &pb.AddCardRequestV1{
				Number:  "1234-5678-1234-5678",
				Owner:   "John Doe",
				Expiry:  "12/25",
				Cvc:     "123",
				Pin:     "1234",
				Comment: "comment",
			},
			callGetCurrentUser: true,
			callCreateCard:     true,
			expectedResp:       &pb.AddCardResponseV1{},
			expectedErr:        nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "user-id"}, tt.getCurrentUserErr)
			}
			if tt.callCreateCard {
				mockCardService.EXPECT().CreateCard(gomock.Any(), gomock.Any()).Return(tt.createCardErr)
			}

			resp, err := s.AddCard(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestDeleteCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.DeleteCardRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callDeleteCard     bool
		deleteCardErr      error
		expectedResp       *pb.DeleteCardResponseV1
		expectedErr        error
	}{
		{
			name: "should return unauthenticated error if authentication fails",
			req: &pb.DeleteCardRequestV1{
				Id: "card-id",
			},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name: "should return invalid argument error if request is invalid",
			req: &pb.DeleteCardRequestV1{
				Id: "",
			},
			callGetCurrentUser: true,
			expectedErr:        status.Error(codes.InvalidArgument, "validation error:\n - id: Deleted card data ID must not be empty [id]"),
		},
		{
			name: "should return not found error if card does not exist",
			req: &pb.DeleteCardRequestV1{
				Id: "card-id",
			},
			callGetCurrentUser: true,
			callDeleteCard:     true,
			deleteCardErr:      er.NewNotFoundError("card not found", nil),
			expectedErr:        status.Error(codes.NotFound, "card not found"),
		},
		{
			name: "should return permission denied error if user does not have permission",
			req: &pb.DeleteCardRequestV1{
				Id: "card-id",
			},
			callGetCurrentUser: true,
			callDeleteCard:     true,
			deleteCardErr:      er.NewForbiddenError("access denied", nil),
			expectedErr:        status.Error(codes.PermissionDenied, "access denied"),
		},
		{
			name: "should return internal error if card service return error",
			req: &pb.DeleteCardRequestV1{
				Id: "card-id",
			},
			callGetCurrentUser: true,
			callDeleteCard:     true,
			deleteCardErr:      errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name: "should delete card successfully",
			req: &pb.DeleteCardRequestV1{
				Id: "card-id",
			},
			callGetCurrentUser: true,
			callDeleteCard:     true,
			expectedResp:       &pb.DeleteCardResponseV1{},
			expectedErr:        nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "user-id"}, tt.getCurrentUserErr)
			}
			if tt.callDeleteCard {
				mockCardService.EXPECT().DeleteCard(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.deleteCardErr)
			}

			resp, err := s.DeleteCard(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestGetCards(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.GetCardsRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetCards       bool
		getCardsErr        error
		cards              []model.Card
		expectedResp       *pb.GetCardsResponseV1
		expectedErr        error
	}{
		{
			name:               "should return unauthenticated error if authentication fails",
			req:                &pb.GetCardsRequestV1{},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name:               "should return internal error if card service return error",
			req:                &pb.GetCardsRequestV1{},
			callGetCurrentUser: true,
			callGetCards:       true,
			getCardsErr:        errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name:               "should get cards successfully",
			req:                &pb.GetCardsRequestV1{},
			callGetCurrentUser: true,
			callGetCards:       true,
			cards:              []model.Card{{ID: "1", Owner: "John Doe"}},
			expectedResp: &pb.GetCardsResponseV1{
				Cards: []*pb.CardV1{
					{Id: "1", Owner: "John Doe"},
				},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "user-id"}, tt.getCurrentUserErr)
			}
			if tt.callGetCards {
				mockCardService.EXPECT().GetAllCards(gomock.Any(), gomock.Any()).Return(tt.cards, tt.getCardsErr)
			}

			resp, err := s.GetCards(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestAddText(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.AddTextRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callCreateText     bool
		createTextErr      error
		expectedResp       *pb.AddTextResponseV1
		expectedErr        error
	}{
		{
			name: "should return unauthenticated error if authentication fails",
			req: &pb.AddTextRequestV1{
				Text:    "text",
				Comment: "comment",
			},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name: "should return invalid argument error if request is invalid",
			req: &pb.AddTextRequestV1{
				Text:    "",
				Comment: "comment",
			},
			callGetCurrentUser: true,
			expectedErr:        status.Error(codes.InvalidArgument, "validation error:\n - text: Text must not be empty [text]"),
		},
		{
			name: "should return internal error if text service return error",
			req: &pb.AddTextRequestV1{
				Text:    "text",
				Comment: "comment",
			},
			callGetCurrentUser: true,
			callCreateText:     true,
			createTextErr:      errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name: "should add text successfully",
			req: &pb.AddTextRequestV1{
				Text:    "text",
				Comment: "comment",
			},
			callGetCurrentUser: true,
			callCreateText:     true,
			expectedResp:       &pb.AddTextResponseV1{},
			expectedErr:        nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "user-id"}, tt.getCurrentUserErr)
			}
			if tt.callCreateText {
				mockTextService.EXPECT().CreateText(gomock.Any(), gomock.Any()).Return(tt.createTextErr)
			}

			resp, err := s.AddText(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestDeleteText(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.DeleteTextRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callDeleteText     bool
		deleteTextErr      error
		expectedResp       *pb.DeleteTextResponseV1
		expectedErr        error
	}{
		{
			name: "should return unauthenticated error if authentication fails",
			req: &pb.DeleteTextRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name: "should return invalid argument error if request is invalid",
			req: &pb.DeleteTextRequestV1{
				Id: "",
			},
			callGetCurrentUser: true,
			expectedErr:        status.Error(codes.InvalidArgument, "validation error:\n - id: Deleted text data ID must not be empty [id]"),
		},
		{
			name: "should return not found error if text does not exist",
			req: &pb.DeleteTextRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteText:     true,
			deleteTextErr:      er.NewNotFoundError("text not found", nil),
			expectedErr:        status.Error(codes.NotFound, "text not found"),
		},
		{
			name: "should return permission denied error if user does not have permission",
			req: &pb.DeleteTextRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteText:     true,
			deleteTextErr:      er.NewForbiddenError("access denied", nil),
			expectedErr:        status.Error(codes.PermissionDenied, "access denied"),
		},
		{
			name: "should return internal error if text service return error",
			req: &pb.DeleteTextRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteText:     true,
			deleteTextErr:      errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name: "should delete text successfully",
			req: &pb.DeleteTextRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteText:     true,
			expectedResp:       &pb.DeleteTextResponseV1{},
			expectedErr:        nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "user-id"}, tt.getCurrentUserErr)
			}
			if tt.callDeleteText {
				mockTextService.EXPECT().DeleteText(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.deleteTextErr)
			}

			resp, err := s.DeleteText(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestGetTexts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.GetTextsRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetTexts       bool
		getTextsErr        error
		texts              []model.Text
		expectedResp       *pb.GetTextsResponseV1
		expectedErr        error
	}{
		{
			name:               "should return unauthenticated error if authentication fails",
			req:                &pb.GetTextsRequestV1{},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name:               "should return internal error if text service return error",
			req:                &pb.GetTextsRequestV1{},
			callGetCurrentUser: true,
			callGetTexts:       true,
			getTextsErr:        errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name:               "should return texts successfully",
			req:                &pb.GetTextsRequestV1{},
			callGetCurrentUser: true,
			callGetTexts:       true,
			texts:              []model.Text{{ID: "1", Text: "text", Comment: "comment"}},
			expectedResp:       &pb.GetTextsResponseV1{Texts: []*pb.TextV1{{Id: "1", Text: "text", Comment: "comment"}}},
			expectedErr:        nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "user-id"}, tt.getCurrentUserErr)
			}
			if tt.callGetTexts {
				mockTextService.EXPECT().GetAllTexts(gomock.Any(), gomock.Any()).Return(tt.texts, tt.getTextsErr)
			}

			resp, err := s.GetTexts(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

type mockUploadStream struct {
	pb.FileServiceV1_UploadFileServer
}

func (m mockUploadStream) SendAndClose(resp *pb.UploadFileResponseV1) error {
	return nil
}

func (m mockUploadStream) Recv() (*pb.UploadFileRequestV1, error) {
	return &pb.UploadFileRequestV1{}, io.EOF
}

func (m mockUploadStream) Context() context.Context {
	claims := jwt.MapClaims{
		"login": "login",
	}

	ctx := context.WithValue(context.Background(), AuthorizationTokenContextKey, claims)
	return ctx
}

func TestUploadFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.UploadFileRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callUploadFile     bool
		uploadFileErr      error
		expectedErr        error
	}{
		{
			name: "should return unauthenticated error if authentication fails",
			req: &pb.UploadFileRequestV1{
				Name: "file",
				Data: []byte("data"),
			},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name: "should return internal error if file service return error",
			req: &pb.UploadFileRequestV1{
				Name: "file",
				Data: []byte("data"),
			},
			callGetCurrentUser: true,
			callUploadFile:     true,
			uploadFileErr:      errors.New("internal server error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name: "should upload file successfully",
			req: &pb.UploadFileRequestV1{
				Name: "file",
				Data: []byte("data"),
			},
			callGetCurrentUser: true,
			callUploadFile:     true,
			expectedErr:        nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callUploadFile {
				mockFileService.EXPECT().UploadFile(gomock.Any(), "1", gomock.Any()).Return(tt.uploadFileErr)
			}

			stream := mockUploadStream{}
			err := s.UploadFile(stream)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
			}
		})
	}
}

type mockDownloadStream struct {
	pb.FileServiceV1_DownloadFileServer
}

func (m mockDownloadStream) SendAndClose(resp *pb.DownloadFileResponseV1) error {
	return nil
}

func (m mockDownloadStream) Context() context.Context {
	claims := jwt.MapClaims{
		"login": "login",
	}

	ctx := context.WithValue(context.Background(), AuthorizationTokenContextKey, claims)
	return ctx
}

func TestDownloadFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, nil, nil, nil, nil, mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.DownloadFileRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callDownloadFile   bool
		downloadFileErr    error
		expectedResp       *pb.DownloadFileResponseV1
		expectedErr        error
	}{
		{
			name: "should return unauthenticated error if authentication fails",
			req: &pb.DownloadFileRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name: "should return internal error if file service return error",
			req: &pb.DownloadFileRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDownloadFile:   true,
			downloadFileErr:    errors.New("internal server error"),
			expectedErr:        status.Error(codes.Internal, "Internal server error"),
		},
		{
			name: "should download file successfully",
			req: &pb.DownloadFileRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDownloadFile:   true,
			expectedResp: &pb.DownloadFileResponseV1{
				Name: "file",
				Size: 1024,
				Data: []byte("data"),
			},
			expectedErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callDownloadFile {
				mockFileService.EXPECT().DownloadFile(gomock.Any(), "1", tt.req.Id, gomock.Any()).Return(tt.downloadFileErr)
			}

			stream := mockDownloadStream{}
			err := s.DownloadFile(tt.req, stream)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
			}
		})
	}
}

func TestDeleteFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.DeleteFileRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callDeleteFile     bool
		deleteFileErr      error
		expectedResp       *pb.DeleteFileResponseV1
		expectedErr        error
	}{
		{
			name: "should return unauthenticated error if authentication fails",
			req: &pb.DeleteFileRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name: "should return invalid argument error if request is invalid",
			req: &pb.DeleteFileRequestV1{
				Id: "",
			},
			callGetCurrentUser: true,
			expectedErr:        status.Error(codes.InvalidArgument, "validation error:\n - id: Deleted file ID must not be empty [id]"),
		},
		{
			name: "should return not found error if file does not exist",
			req: &pb.DeleteFileRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteFile:     true,
			deleteFileErr:      er.NewNotFoundError("file not found", nil),
			expectedErr:        status.Error(codes.NotFound, "file not found"),
		},
		{
			name: "should return permission denied error if user does not have permission",
			req: &pb.DeleteFileRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteFile:     true,
			deleteFileErr:      er.NewForbiddenError("access denied", nil),
			expectedErr:        status.Error(codes.PermissionDenied, "access denied"),
		},
		{
			name: "should return internal error if file service return error",
			req: &pb.DeleteFileRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteFile:     true,
			deleteFileErr:      errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name: "should delete file successfully",
			req: &pb.DeleteFileRequestV1{
				Id: "1",
			},
			callGetCurrentUser: true,
			callDeleteFile:     true,
			expectedResp:       &pb.DeleteFileResponseV1{},
			expectedErr:        nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callDeleteFile {
				mockFileService.EXPECT().DeleteFile(gomock.Any(), "1", tt.req.Id).Return(tt.deleteFileErr)
			}

			resp, err := s.DeleteFile(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}

func TestGetFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := NewMockUserService(ctrl)
	mockAuthService := NewMockAuthService(ctrl)
	mockAccountService := NewMockAccountService(ctrl)
	mockCardService := NewMockCardService(ctrl)
	mockTextService := NewMockTextService(ctrl)
	mockFileService := NewMockFileService(ctrl)
	cfg := &config.ServerConfig{}

	validator, err := protovalidate.New()
	require.NoError(t, err, "unexpected error creating validator")

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	s := NewServer(mockUserService, mockAuthService, mockAccountService, mockCardService, mockTextService,
		mockFileService, validator, cfg, l)

	testCases := []struct {
		name               string
		req                *pb.GetFilesRequestV1
		callGetCurrentUser bool
		getCurrentUserErr  error
		callGetFiles       bool
		files              []model.File
		getFilesErr        error
		expectedResp       *pb.GetFilesResponseV1
		expectedErr        error
	}{
		{
			name:               "should return unauthenticated error if authentication fails",
			req:                &pb.GetFilesRequestV1{},
			callGetCurrentUser: true,
			getCurrentUserErr:  errors.New("authentication error"),
			expectedErr:        status.Error(codes.Unauthenticated, "authentication error"),
		},
		{
			name:               "should return internal error if file service return error",
			req:                &pb.GetFilesRequestV1{},
			callGetCurrentUser: true,
			callGetFiles:       true,
			files:              make([]model.File, 0),
			getFilesErr:        errors.New("internal error"),
			expectedErr:        status.Error(codes.Internal, "internal server error"),
		},
		{
			name:               "should get files successfully",
			req:                &pb.GetFilesRequestV1{},
			callGetCurrentUser: true,
			callGetFiles:       true,
			files: []model.File{{Name: "file 1", Size: 1028},
				{Name: "file 2", Size: 2056}},
			expectedResp: &pb.GetFilesResponseV1{
				Files: []*pb.FileV1{
					{Name: "file 1", Size: 1028},
					{Name: "file 2", Size: 2056},
				},
			},
			expectedErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.callGetCurrentUser {
				mockUserService.EXPECT().GetCurrentUser(gomock.Any()).Return(model.User{ID: "1"}, tt.getCurrentUserErr)
			}
			if tt.callGetFiles {
				mockFileService.EXPECT().GetAllFiles(gomock.Any(), "1").Return(tt.files, tt.getFilesErr)
			}

			resp, err := s.GetFiles(context.Background(), tt.req)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr, "should return error: %s, got: %s", tt.expectedErr, err)
			} else {
				assert.NoError(t, err, "shouldn't return error, but got: %s", err)
				assert.Equal(t, tt.expectedResp, resp, "should return response: %v, got: %v", tt.expectedResp, resp)
			}
		})
	}
}
