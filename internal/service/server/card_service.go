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

// CardService defines the interface for operations with cards data.
type CardService interface {
	CreateCard(ctx context.Context, card model.Card) error
	DeleteCard(ctx context.Context, userID string, cardID string) error
	GetCardByID(ctx context.Context, userID string, cardID string) (model.Card, error)
	GetAllCards(ctx context.Context, userID string) ([]model.Card, error)
}

// CardServiceImpl is an implementation of the CardService interface.
type CardServiceImpl struct {
	dataStorage         storage.DataStorage
	encryptionSecretKey string
	logger              logger.ServerLogger
}

// NewCardService creates a new instance of CardService.
func NewCardService(dataStorage storage.DataStorage, encryptionSecretKey string, logger logger.ServerLogger) CardService {
	return &CardServiceImpl{
		dataStorage:         dataStorage,
		encryptionSecretKey: encryptionSecretKey,
		logger:              logger,
	}
}

// CreateCard creates a new card data and stores it in the data storage.
func (s *CardServiceImpl) CreateCard(ctx context.Context, card model.Card) error {
	data := model.CardToData(card)
	sensitiveData := card.GetSensitiveCardData()
	jsonData, err := json.Marshal(sensitiveData)
	if err != nil {
		s.logger.Error("Failed marshall sensitive card data", zap.String("event", "save card data"),
			zap.Error(err))
		return fmt.Errorf("failed marshall sensitive card data: %w", err)
	}

	encryptedSensitiveData, err := EncryptData(jsonData, []byte(s.encryptionSecretKey))
	if err != nil {
		s.logger.Error("Failed encrypt sensitive card data", zap.String("event", "save card data"),
			zap.Error(err))
		return fmt.Errorf("failed encrypt sensitive card data: %w", err)
	}
	data.SensitiveData = encryptedSensitiveData

	err = s.dataStorage.Create(ctx, data)
	if err != nil {
		s.logger.Error("Failed save card data", zap.String("event", "save card data"),
			zap.Error(err))
		return fmt.Errorf("failed save card data: %w", err)
	}

	return nil
}

// DeleteCard deletes the card data with the specified id.
// The user id is checked to ensure that the user has permission to delete the card data.
// Returns an error if the card data is not found, if access is denied, or if the deletion fails.
func (s *CardServiceImpl) DeleteCard(ctx context.Context, userID string, cardID string) error {
	card, err := s.dataStorage.GetByID(ctx, cardID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Card does not exists", zap.String("event", "delete card data"),
				zap.String("card id", cardID), zap.String("user id", userID), zap.Error(err))
			return er.NewNotFoundError(fmt.Sprintf("card with id:%s does not exist", cardID), err)
		default:
			s.logger.Error("Failed to get card by id", zap.String("event", "delete card data"),
				zap.String("card id", cardID), zap.String("user id", userID), zap.Error(err))
			return fmt.Errorf("failed get card by id: %w", err)
		}
	}

	if card.UserID != userID {
		s.logger.Warn("User attempted to access card belonging to another user",
			zap.String("event", "delete card data"), zap.String("card id", cardID),
			zap.String("user id", userID), zap.Error(err))
		return er.NewForbiddenError("access denied", err)
	}

	err = s.dataStorage.Delete(ctx, cardID)
	if err != nil {
		s.logger.Error("Failed to delete card data", zap.String("event", "delete card data"),
			zap.String("card id", cardID), zap.String("user id", userID), zap.Error(err))
		return err
	}

	return nil
}

// GetCardByID retrieves the card data with the specified id.
// The user id is used to ensure the user has permission to access the card data.
// The sensitive data is decrypted before returning the card data.
// Returns the card data with the specified id or an error if the card data is not found or access is denied.
func (s *CardServiceImpl) GetCardByID(ctx context.Context, userID string, cardID string) (model.Card, error) {
	data, err := s.dataStorage.GetByID(ctx, cardID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Card does not exists", zap.String("event", "get card data by id"),
				zap.String("card id", cardID), zap.String("user id", userID), zap.Error(err))
			return model.Card{}, er.NewNotFoundError(fmt.Sprintf("card with id:%s does not exist", cardID), err)
		default:
			s.logger.Error("Failed to get card by id", zap.String("event", "get card data by id"),
				zap.String("card id", cardID), zap.String("user id", userID), zap.Error(err))
			return model.Card{}, fmt.Errorf("failed get card by id: %w", err)
		}
	}

	card := model.DataToCard(data)
	decryptedSensitiveData, err := DecryptData(data.SensitiveData, []byte(s.encryptionSecretKey))
	if err != nil {
		s.logger.Error("Failed decrypt sensitive card data", zap.String("event", "get card data by id"),
			zap.String("card id", cardID), zap.String("user id", userID), zap.Error(err))
		return model.Card{}, fmt.Errorf("failed decrypt sensitive card data: %w", err)
	}

	sensitiveData := model.SensitiveCardData{}
	err = json.Unmarshal(decryptedSensitiveData, &sensitiveData)
	if err != nil {
		s.logger.Error("Failed unmarshall sensitive card data", zap.String("event", "get card data by id"),
			zap.String("card id", cardID), zap.String("user id", userID), zap.Error(err))
		return model.Card{}, fmt.Errorf("failed unmarshall sensitive card data: %w", err)
	}
	card.SetSensitiveCardData(sensitiveData)

	return card, nil
}

// GetAllCards retrieves all cards data associated with the specified user.
// Sensitive data is decrypted before returning cards data.
// Returns a slice of cards data or an error if the retrieval fails.
func (s *CardServiceImpl) GetAllCards(ctx context.Context, userID string) ([]model.Card, error) {
	cardsData, err := s.dataStorage.GetAll(ctx, userID, model.CardType)
	if err != nil {
		s.logger.Error("Failed get cards data", zap.String("event", "get cards data"),
			zap.String("user id", userID), zap.Error(err))
		return []model.Card{}, fmt.Errorf("failed get cards data: %w", err)
	}

	cards := make([]model.Card, 0)
	for _, data := range cardsData {
		card := model.DataToCard(data)
		decryptedSensitiveData, err := DecryptData(data.SensitiveData, []byte(s.encryptionSecretKey))
		if err != nil {
			s.logger.Error("Failed decrypt sensitive cards data", zap.String("event", "get cards data"),
				zap.String("user id", userID), zap.Error(err))
			return []model.Card{}, fmt.Errorf("failed decrypt sensitive cards data: %w", err)
		}

		sensitiveData := model.SensitiveCardData{}
		err = json.Unmarshal(decryptedSensitiveData, &sensitiveData)
		if err != nil {
			s.logger.Error("Failed unmarshall sensitive cards data", zap.String("event", "get cards data"),
				zap.String("user id", userID), zap.Error(err))
			return []model.Card{}, fmt.Errorf("failed unmarshall sensitive cards data: %w", err)
		}
		card.SetSensitiveCardData(sensitiveData)

		cards = append(cards, card)
	}

	return cards, nil
}
