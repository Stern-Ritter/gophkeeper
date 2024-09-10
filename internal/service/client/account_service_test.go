package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	pb "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
)

func TestCreateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountClient := NewMockAccountServiceV1Client(mockCtrl)
	accountClient.EXPECT().AddAccount(gomock.Any(), &pb.AddAccountRequestV1{
		Login:    "login",
		Password: "password",
		Comment:  "comment",
	}).Return(&pb.AddAccountResponseV1{}, nil)

	service := NewAccountService(accountClient)
	err := service.CreateAccount("login", "password", "comment")

	assert.NoError(t, err, "Expected no error when creating account")
}

func TestGetAllAccounts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountClient := NewMockAccountServiceV1Client(mockCtrl)
	accountClient.EXPECT().GetAccounts(gomock.Any(), &pb.GetAccountsRequestV1{}).
		Return(&pb.GetAccountsResponseV1{
			Accounts: []*pb.AccountV1{
				{Login: "login 1", Password: "password 1", Comment: "comment 1"},
				{Login: "login 2", Password: "password 2", Comment: "comment 2"},
			},
		}, nil)

	service := NewAccountService(accountClient)
	accounts, err := service.GetAllAccounts()

	assert.NoError(t, err, "Expected no error when retrieving accounts")
	assert.Len(t, accounts, 2, "Expected two accounts")

	assert.Equal(t, "login 1", accounts[0].Login)
	assert.Equal(t, "password 1", accounts[0].Password)
	assert.Equal(t, "comment 1", accounts[0].Comment)

	assert.Equal(t, "login 2", accounts[1].Login)
	assert.Equal(t, "password 2", accounts[1].Password)
	assert.Equal(t, "comment 2", accounts[1].Comment)
}

func TestDeleteAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	accountClient := NewMockAccountServiceV1Client(mockCtrl)
	accountClient.EXPECT().DeleteAccount(gomock.Any(), &pb.DeleteAccountRequestV1{Id: "1"}).
		Return(&pb.DeleteAccountResponseV1{}, nil)

	service := NewAccountService(accountClient)
	err := service.DeleteAccount("1")

	assert.NoError(t, err, "Expected no error when deleting account")
}
