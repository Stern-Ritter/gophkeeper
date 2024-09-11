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
	"github.com/Stern-Ritter/gophkeeper/internal/storage/server"
)

// TextService defines the interface for operations with texts data.
type TextService interface {
	CreateText(ctx context.Context, text model.Text) error
	DeleteText(ctx context.Context, userID string, textID string) error
	GetTextByID(ctx context.Context, userID string, textID string) (model.Text, error)
	GetAllTexts(ctx context.Context, userID string) ([]model.Text, error)
}

// TextServiceImpl is an implementation of the TextService interface.
type TextServiceImpl struct {
	dataStorage         server.DataStorage
	encryptionSecretKey string
	logger              logger.ServerLogger
}

// NewTextService creates a new instance of TextService.
func NewTextService(dataStorage server.DataStorage, encryptionSecretKey string, logger logger.ServerLogger) TextService {
	return &TextServiceImpl{
		dataStorage:         dataStorage,
		encryptionSecretKey: encryptionSecretKey,
		logger:              logger,
	}
}

// CreateText creates a new text data and stores it in the data storage.
func (s *TextServiceImpl) CreateText(ctx context.Context, text model.Text) error {
	data := model.TextToData(text)
	sensitiveData := text.GetSensitiveTextData()
	jsonData, err := json.Marshal(sensitiveData)
	if err != nil {
		s.logger.Error("Failed marshall sensitive text data", zap.String("event", "save text data"),
			zap.Error(err))
		return fmt.Errorf("failed marshall sensitive text data: %w", err)
	}

	encryptedSensitiveData, err := EncryptData(jsonData, []byte(s.encryptionSecretKey))
	if err != nil {
		s.logger.Error("Failed encrypt sensitive text data", zap.String("event", "save text data"),
			zap.Error(err))
		return fmt.Errorf("failed encrypt sensitive text data: %w", err)
	}
	data.SensitiveData = encryptedSensitiveData

	err = s.dataStorage.Create(ctx, data)
	if err != nil {
		s.logger.Error("Failed save text data", zap.String("event", "save text data"),
			zap.Error(err))
		return fmt.Errorf("failed save text data: %w", err)
	}

	return nil
}

// DeleteText deletes the text data with the specified id.
// The user id is checked to ensure that the user has permission to delete the text data.
// Returns an error if the text data is not found, if access is denied, or if the deletion fails.
func (s *TextServiceImpl) DeleteText(ctx context.Context, userID string, textID string) error {
	text, err := s.dataStorage.GetByID(ctx, textID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Text does not exists", zap.String("event", "delete text data"),
				zap.String("text id", textID), zap.String("user id", userID), zap.Error(err))
			return er.NewNotFoundError(fmt.Sprintf("text with id:%s does not exist", textID), err)
		default:
			s.logger.Error("Failed to get text by id", zap.String("event", "delete text data"),
				zap.String("text id", textID), zap.String("user id", userID), zap.Error(err))
			return fmt.Errorf("failed get text by id: %w", err)
		}
	}

	if text.UserID != userID {
		s.logger.Warn("User attempted to access text belonging to another user",
			zap.String("event", "delete text data"), zap.String("text id", textID),
			zap.String("user id", userID), zap.Error(err))
		return er.NewForbiddenError("access denied", err)
	}

	err = s.dataStorage.Delete(ctx, textID)
	if err != nil {
		s.logger.Error("Failed to delete text data", zap.String("event", "delete text data"),
			zap.String("text id", textID), zap.String("user id", userID), zap.Error(err))
		return err
	}

	return nil
}

// GetTextByID retrieves the text data with the specified id.
// The user id is used to ensure the user has permission to access the text data.
// The sensitive data is decrypted before returning the text data.
// Returns the text data with the specified id or an error if the text data is not found or access is denied.
func (s *TextServiceImpl) GetTextByID(ctx context.Context, userID string, textID string) (model.Text, error) {
	data, err := s.dataStorage.GetByID(ctx, textID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("Text does not exists", zap.String("event", "get text data by id"),
				zap.String("text id", textID), zap.String("user id", userID), zap.Error(err))
			return model.Text{}, er.NewNotFoundError(fmt.Sprintf("text with id:%s does not exist", textID), err)
		default:
			s.logger.Error("Failed to get text by id", zap.String("event", "get text data by id"),
				zap.String("text id", textID), zap.String("user id", userID), zap.Error(err))
			return model.Text{}, fmt.Errorf("failed get text by id: %w", err)
		}
	}

	text := model.DataToText(data)
	decryptedSensitiveData, err := DecryptData(data.SensitiveData, []byte(s.encryptionSecretKey))
	if err != nil {
		s.logger.Error("Failed decrypt sensitive text data", zap.String("event", "get text data by id"),
			zap.String("text id", textID), zap.String("user id", userID), zap.Error(err))
		return model.Text{}, fmt.Errorf("failed decrypt sensitive text data: %w", err)
	}

	sensitiveData := model.SensitiveTextData{}
	err = json.Unmarshal(decryptedSensitiveData, &sensitiveData)
	if err != nil {
		s.logger.Error("Failed unmarshall sensitive text data", zap.String("event", "get text data by id"),
			zap.String("text id", textID), zap.String("user id", userID), zap.Error(err))
		return model.Text{}, fmt.Errorf("failed unmarshall sensitive text data: %w", err)
	}
	text.SetSensitiveTextData(sensitiveData)

	return text, nil
}

// GetAllTexts retrieves all texts data associated with the specified user.
// Sensitive data is decrypted before returning texts data.
// Returns a slice of texts data or an error if the retrieval fails.
func (s *TextServiceImpl) GetAllTexts(ctx context.Context, userID string) ([]model.Text, error) {
	textsData, err := s.dataStorage.GetAll(ctx, userID, model.TextType)
	if err != nil {
		s.logger.Error("Failed get texts data", zap.String("event", "get texts data"),
			zap.String("user id", userID), zap.Error(err))
		return []model.Text{}, fmt.Errorf("failed get texts data: %w", err)
	}

	texts := make([]model.Text, 0)
	for _, data := range textsData {
		text := model.DataToText(data)
		decryptedSensitiveData, err := DecryptData(data.SensitiveData, []byte(s.encryptionSecretKey))
		if err != nil {
			s.logger.Error("Failed decrypt sensitive texts data", zap.String("event", "get texts data"),
				zap.String("user id", userID), zap.Error(err))
			return []model.Text{}, fmt.Errorf("failed decrypt sensitive text data: %w", err)
		}

		sensitiveData := model.SensitiveTextData{}
		err = json.Unmarshal(decryptedSensitiveData, &sensitiveData)
		if err != nil {
			s.logger.Error("Failed unmarshall sensitive texts data", zap.String("event", "get texts data"),
				zap.String("user id", userID), zap.Error(err))
			return []model.Text{}, fmt.Errorf("failed unmarshall sensitive text data: %w", err)
		}
		text.SetSensitiveTextData(sensitiveData)

		texts = append(texts, text)
	}

	return texts, nil
}
