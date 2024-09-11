package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	er "github.com/Stern-Ritter/gophkeeper/internal/errors"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
	storage "github.com/Stern-Ritter/gophkeeper/internal/storage/server"
)

// AccountService defines the interface for operations with accounts data.
type AccountService interface {
	CreateAccount(ctx context.Context, account model.Account) error
	DeleteAccount(ctx context.Context, userID string, accountID string) error
	GetAccountByID(ctx context.Context, userID string, accountID string) (model.Account, error)
	GetAllAccounts(ctx context.Context, userID string) ([]model.Account, error)
}

// AccountServiceImpl is an implementation of the AccountService interface.
type AccountServiceImpl struct {
	dataStorage         storage.DataStorage
	encryptionSecretKey string
	logger              logger.ServerLogger
}

// NewAccountService creates a new instance of AccountService.
func NewAccountService(dataStorage storage.DataStorage, encryptionSecretKey string, logger logger.ServerLogger) AccountService {
	return &AccountServiceImpl{
		dataStorage:         dataStorage,
		encryptionSecretKey: encryptionSecretKey,
		logger:              logger,
	}
}

// CreateAccount creates a new account data and stores it in the data storage.
func (s *AccountServiceImpl) CreateAccount(ctx context.Context, account model.Account) error {
	data := model.AccountToData(account)
	sensitiveData := account.GetSensitiveAccountData()
	jsonData, err := json.Marshal(sensitiveData)
	if err != nil {
		s.logger.Error("Failed marshall sensitive account data", zap.String("event", "save account data"),
			zap.Error(err))
		return fmt.Errorf("failed marshall sensitive account data: %w", err)
	}

	encryptedSensitiveData, err := EncryptData(jsonData, []byte(s.encryptionSecretKey))
	if err != nil {
		s.logger.Error("Failed encrypt sensitive account data", zap.String("event", "save account data"),
			zap.Error(err))
		return fmt.Errorf("failed encrypt sensitive account data: %w", err)
	}
	data.SensitiveData = encryptedSensitiveData

	err = s.dataStorage.Create(ctx, data)
	if err != nil {
		s.logger.Error("Failed save account data", zap.String("event", "save account data"),
			zap.Error(err))
		return fmt.Errorf("failed save account data: %w", err)
	}

	return nil
}

// DeleteAccount deletes the account data with the specified id.
// The user id is checked to ensure that the user has permission to delete the account data.
// Returns an error if the account data is not found, if access is denied, or if the deletion fails.
func (s *AccountServiceImpl) DeleteAccount(ctx context.Context, userID string, accountID string) error {
	account, err := s.dataStorage.GetByID(ctx, accountID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Account does not exists", zap.String("event", "delete account data"),
				zap.String("account id", accountID), zap.String("user id", userID), zap.Error(err))
			return er.NewNotFoundError(fmt.Sprintf("account with id:%s does not exist", accountID), err)
		default:
			s.logger.Error("Failed to get account by id", zap.String("event", "delete account data"),
				zap.String("account id", accountID), zap.String("user id", userID), zap.Error(err))
			return fmt.Errorf("failed get account by id: %w", err)
		}
	}

	if account.UserID != userID {
		s.logger.Warn("User attempted to access account belonging to another user",
			zap.String("event", "delete account data"), zap.String("account id", accountID),
			zap.String("user id", userID), zap.Error(err))
		return er.NewForbiddenError("access denied", err)
	}

	err = s.dataStorage.Delete(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to delete account data", zap.String("event", "delete account data"),
			zap.String("account id", accountID), zap.String("user id", userID), zap.Error(err))
		return err
	}

	return nil
}

// GetAccountByID retrieves the account data with the specified id.
// The user id is used to ensure the user has permission to access the account data.
// The sensitive data is decrypted before returning the account data.
// Returns the account data with the specified id or an error if the account data is not found or access is denied.
func (s *AccountServiceImpl) GetAccountByID(ctx context.Context, userID string, accountID string) (model.Account, error) {
	data, err := s.dataStorage.GetByID(ctx, accountID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Account does not exists", zap.String("event", "get account data by id"),
				zap.String("account id", accountID), zap.String("user id", userID), zap.Error(err))
			return model.Account{}, er.NewNotFoundError(fmt.Sprintf("account with id:%s does not exist", accountID), err)
		default:
			s.logger.Error("Failed to get account by id", zap.String("event", "get account data by id"),
				zap.String("account id", accountID), zap.String("user id", userID), zap.Error(err))
			return model.Account{}, fmt.Errorf("failed get account by id: %w", err)
		}
	}

	account := model.DataToAccount(data)
	decryptedSensitiveData, err := DecryptData(data.SensitiveData, []byte(s.encryptionSecretKey))
	if err != nil {
		s.logger.Error("Failed decrypt sensitive account data", zap.String("event", "get account data by id"),
			zap.String("account id", accountID), zap.String("user id", userID), zap.Error(err))
		return model.Account{}, fmt.Errorf("failed decrypt sensitive account data: %w", err)
	}

	sensitiveData := model.SensitiveAccountData{}
	err = json.Unmarshal(decryptedSensitiveData, &sensitiveData)
	if err != nil {
		s.logger.Error("Failed unmarshall sensitive account data", zap.String("event", "get account data by id"),
			zap.String("account id", accountID), zap.String("user id", userID), zap.Error(err))
		return model.Account{}, fmt.Errorf("failed unmarshall sensitive account data: %w", err)
	}
	account.SetSensitiveAccountData(sensitiveData)

	return account, nil
}

// GetAllAccounts retrieves all accounts data associated with the specified user.
// Sensitive data is decrypted before returning accounts data.
// Returns a slice of account data or an error if the retrieval fails.
func (s *AccountServiceImpl) GetAllAccounts(ctx context.Context, userID string) ([]model.Account, error) {
	accountsData, err := s.dataStorage.GetAll(ctx, userID, model.AccountType)
	if err != nil {
		s.logger.Error("Failed get accounts data", zap.String("event", "get accounts data"),
			zap.String("user id", userID), zap.Error(err))
		return []model.Account{}, fmt.Errorf("failed get accounts data: %w", err)
	}

	accounts := make([]model.Account, 0)
	for _, data := range accountsData {
		account := model.DataToAccount(data)
		decryptedSensitiveData, err := DecryptData(data.SensitiveData, []byte(s.encryptionSecretKey))
		if err != nil {
			s.logger.Error("Failed decrypt sensitive accounts data", zap.String("event", "get accounts data"),
				zap.String("user id", userID), zap.Error(err))
			return []model.Account{}, fmt.Errorf("failed decrypt sensitive account data: %w", err)
		}

		sensitiveData := model.SensitiveAccountData{}
		err = json.Unmarshal(decryptedSensitiveData, &sensitiveData)
		if err != nil {
			s.logger.Error("Failed unmarshall sensitive accounts data", zap.String("event", "get accounts data"),
				zap.String("user id", userID), zap.Error(err))
			return []model.Account{}, fmt.Errorf("failed unmarshall sensitive account data: %w", err)
		}
		account.SetSensitiveAccountData(sensitiveData)

		accounts = append(accounts, account)
	}

	return accounts, nil
}
