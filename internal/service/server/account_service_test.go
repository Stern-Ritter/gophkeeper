package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/Stern-Ritter/gophkeeper/internal/errors"
	e "github.com/Stern-Ritter/gophkeeper/internal/errors"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

func TestCreateAccount_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	accountService := NewAccountService(mockAccountStorage, encryptionKey, l)

	account := model.Account{
		ID:       "1",
		UserID:   "1",
		Login:    "user",
		Password: "password",
		Comment:  "comment",
	}

	mockAccountStorage.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	err = accountService.CreateAccount(context.Background(), account)
	assert.NoError(t, err, "expected no error creating account")
}

func TestCreateAccount_EncryptionFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "invalid key"
	accountService := NewAccountService(mockAccountStorage, encryptionKey, l)

	account := model.Account{
		ID:       "1",
		UserID:   "1",
		Login:    "user",
		Password: "password",
		Comment:  "comment",
	}

	err = accountService.CreateAccount(context.Background(), account)
	assert.Error(t, err, "expected error encrypting sensitive data")
}

func TestDeleteAccount_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	accountService := NewAccountService(mockAccountStorage, encryptionKey, l)

	data := model.Data{
		ID:            "1",
		UserID:        "1",
		Type:          model.AccountType,
		SensitiveData: make([]byte, 0),
		Comment:       "comment",
	}

	mockAccountStorage.EXPECT().GetByID(gomock.Any(), "1").Return(data, nil)
	mockAccountStorage.EXPECT().Delete(gomock.Any(), "1").Return(nil)

	err = accountService.DeleteAccount(context.Background(), "1", "1")
	assert.NoError(t, err, "unexpected error deleting account")
}

func TestDeleteAccount_Forbidden(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	accountService := NewAccountService(mockAccountStorage, encryptionKey, l)

	data := model.Data{
		ID:            "1",
		UserID:        "1",
		Type:          model.AccountType,
		SensitiveData: make([]byte, 0),
		Comment:       "comment",
	}

	mockAccountStorage.EXPECT().GetByID(gomock.Any(), "1").Return(data, nil)

	err = accountService.DeleteAccount(context.Background(), "2", "1")
	assert.ErrorAs(t, err, &errors.ForbiddenError{}, "expected forbidden error")
}

func TestDeleteAccount_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	accountService := NewAccountService(mockAccountStorage, encryptionKey, l)

	mockAccountStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, pgx.ErrNoRows)

	err = accountService.DeleteAccount(context.Background(), "1", "1")
	assert.ErrorAs(t, err, &e.NotFoundError{}, "expected not found error")
}

func TestDeleteAccount_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	accountService := NewAccountService(mockAccountStorage, encryptionKey, l)

	mockAccountStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, fmt.Errorf("unexpected internal error"))

	err = accountService.DeleteAccount(context.Background(), "1", "1")
	assert.Error(t, err, "expected internal error")
}

func TestGetAccountByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	accountService := NewAccountService(mockAccountStorage, encryptionKey, l)

	data := model.Data{
		ID:     "1",
		UserID: "1",
		Type:   model.AccountType,
		SensitiveData: []byte{69, 245, 110, 5, 132, 219, 203, 67, 127, 213, 189, 213, 243, 63, 42, 61, 245, 231, 12,
			40, 3, 160, 221, 7, 96, 45, 176, 45, 4, 233, 77, 251, 32, 126, 22, 168, 177, 89, 16, 241, 225, 43, 178, 89,
			134, 66, 228, 85, 2, 99, 184, 36, 71, 192, 52, 26, 249, 122, 145, 194, 248, 188, 41, 149, 130, 91},
		Comment: "comment",
	}

	mockAccountStorage.EXPECT().GetByID(gomock.Any(), "1").Return(data, nil)

	_, err = accountService.GetAccountByID(context.Background(), "1", "1")
	assert.NoError(t, err, "unexpected error fetching account")
}

func TestGetAccountByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	accountService := NewAccountService(mockAccountStorage, encryptionKey, l)

	mockAccountStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, pgx.ErrNoRows)

	_, err = accountService.GetAccountByID(context.Background(), "1", "1")
	assert.ErrorAs(t, err, &e.NotFoundError{}, "expected not found error")
}

func TestGetAccountByID_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	accountService := NewAccountService(mockAccountStorage, encryptionKey, l)

	mockAccountStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, fmt.Errorf("unexpected internal error"))

	_, err = accountService.GetAccountByID(context.Background(), "1", "1")
	assert.Error(t, err, "expected internal error")
}

func TestGetAllAccounts_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	accountService := NewAccountService(mockAccountStorage, encryptionKey, l)

	data := []model.Data{
		{
			ID:     "1",
			UserID: "1",
			SensitiveData: []byte{69, 245, 110, 5, 132, 219, 203, 67, 127, 213, 189, 213, 243, 63, 42, 61, 245, 231, 12,
				40, 3, 160, 221, 7, 96, 45, 176, 45, 4, 233, 77, 251, 32, 126, 22, 168, 177, 89, 16, 241, 225, 43, 178, 89,
				134, 66, 228, 85, 2, 99, 184, 36, 71, 192, 52, 26, 249, 122, 145, 194, 248, 188, 41, 149, 130, 91},
		},
		{
			ID:     "2",
			UserID: "1",
			SensitiveData: []byte{69, 245, 110, 5, 132, 219, 203, 67, 127, 213, 189, 213, 243, 63, 42, 61, 245, 231, 12,
				40, 3, 160, 221, 7, 96, 45, 176, 45, 4, 233, 77, 251, 32, 126, 22, 168, 177, 89, 16, 241, 225, 43, 178, 89,
				134, 66, 228, 85, 2, 99, 184, 36, 71, 192, 52, 26, 249, 122, 145, 194, 248, 188, 41, 149, 130, 91},
		},
	}

	mockAccountStorage.EXPECT().GetAll(gomock.Any(), "1", model.AccountType).Return(data, nil)

	accounts, err := accountService.GetAllAccounts(context.Background(), "1")
	assert.NoError(t, err, "unexpected error fetching all accounts data")
	assert.Len(t, accounts, len(data), "should return two accounts data")
}

func TestGetAllAccounts_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	accountService := NewAccountService(mockAccountStorage, encryptionKey, l)

	mockAccountStorage.EXPECT().GetAll(gomock.Any(), "1", model.AccountType).Return([]model.Data{},
		fmt.Errorf("unexpected internal error"))

	account, err := accountService.GetAllAccounts(context.Background(), "1")
	assert.Error(t, err, "expected internal error")
	assert.Empty(t, account, "expected no accounts data")
}
