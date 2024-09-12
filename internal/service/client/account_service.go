package client

import (
	"context"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

// AccountService defines an interface for user accounts data.
type AccountService interface {
	CreateAccount(login string, password string, comment string) error
	GetAllAccounts() ([]*pb.AccountV1, error)
	DeleteAccount(accountID string) error
}

// AccountServiceImpl is implementation of the AccountService interface.
type AccountServiceImpl struct {
	accountClient pb.AccountServiceV1Client
}

// NewAccountService creates a new instance of AccountService.
func NewAccountService(accountClient pb.AccountServiceV1Client) AccountService {
	return &AccountServiceImpl{
		accountClient: accountClient,
	}
}

// CreateAccount creates new account data with the given login, password, and comment by sending a request
// to the account service with gRPC.
// Returns an error if the account data creation fails.
func (s *AccountServiceImpl) CreateAccount(login string, password string, comment string) error {
	ctx := context.Background()
	req := &pb.AddAccountRequestV1{
		Login:    login,
		Password: password,
		Comment:  comment,
	}

	_, err := s.accountClient.AddAccount(ctx, req)
	return err
}

// GetAllAccounts get all accounts data by sending a request to the account service with gRPC.
func (s *AccountServiceImpl) GetAllAccounts() ([]*pb.AccountV1, error) {
	ctx := context.Background()
	req := &pb.GetAccountsRequestV1{}
	resp, err := s.accountClient.GetAccounts(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Accounts, nil
}

// DeleteAccount deletes account data identified by the given id by sending a request
// to the account service with gRPC.
// Returns an error if the account data deletion fails.
func (s *AccountServiceImpl) DeleteAccount(accountID string) error {
	ctx := context.Background()
	req := &pb.DeleteAccountRequestV1{Id: accountID}

	_, err := s.accountClient.DeleteAccount(ctx, req)
	return err
}
