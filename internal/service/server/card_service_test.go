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

func TestCreateCard_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	cardService := NewCardService(mockCardStorage, encryptionKey, l)

	card := model.Card{
		ID:      "1",
		UserID:  "1",
		Number:  "1234567812345678",
		Owner:   "John Doe",
		Expiry:  "12/25",
		CVC:     "123",
		Pin:     "1234",
		Comment: "comment",
	}

	mockCardStorage.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	err = cardService.CreateCard(context.Background(), card)
	assert.NoError(t, err, "expected no error creating card")
}

func TestCreateCard_EncryptionFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "invalid key"
	cardService := NewCardService(mockCardStorage, encryptionKey, l)

	card := model.Card{
		ID:      "1",
		UserID:  "1",
		Number:  "1234567812345678",
		Owner:   "John Doe",
		Expiry:  "12/25",
		CVC:     "123",
		Pin:     "1234",
		Comment: "comment",
	}

	err = cardService.CreateCard(context.Background(), card)
	assert.Error(t, err, "expected error encrypting sensitive data")
}

func TestDeleteCard_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	cardService := NewCardService(mockCardStorage, encryptionKey, l)

	data := model.Data{
		ID:            "1",
		UserID:        "1",
		Type:          model.CardType,
		SensitiveData: make([]byte, 0),
		Comment:       "comment",
	}

	mockCardStorage.EXPECT().GetByID(gomock.Any(), "1").Return(data, nil)
	mockCardStorage.EXPECT().Delete(gomock.Any(), "1").Return(nil)

	err = cardService.DeleteCard(context.Background(), "1", "1")
	assert.NoError(t, err, "unexpected error deleting card")
}

func TestDeleteCard_Forbidden(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	cardService := NewCardService(mockCardStorage, encryptionKey, l)

	data := model.Data{
		ID:            "1",
		UserID:        "1",
		Type:          model.CardType,
		SensitiveData: make([]byte, 0),
		Comment:       "comment",
	}

	mockCardStorage.EXPECT().GetByID(gomock.Any(), "1").Return(data, nil)

	err = cardService.DeleteCard(context.Background(), "2", "1")
	assert.ErrorAs(t, err, &errors.ForbiddenError{}, "expected forbidden error")
}

func TestDeleteCard_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	cardService := NewCardService(mockCardStorage, encryptionKey, l)

	mockCardStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, pgx.ErrNoRows)

	err = cardService.DeleteCard(context.Background(), "1", "1")
	assert.ErrorAs(t, err, &e.NotFoundError{}, "expected not found error")
}

func TestDeleteCard_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	cardService := NewCardService(mockCardStorage, encryptionKey, l)

	mockCardStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, fmt.Errorf("unexpected internal error"))

	err = cardService.DeleteCard(context.Background(), "1", "1")
	assert.Error(t, err, "expected internal error")
}

func TestGetCardByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	cardService := NewCardService(mockCardStorage, encryptionKey, l)

	data := model.Data{
		ID:     "1",
		UserID: "1",
		Type:   model.CardType,
		SensitiveData: []byte{195, 64, 136, 52, 27, 89, 5, 77, 31, 193, 7, 121, 200, 46, 238, 176, 14, 246, 65, 97, 157,
			34, 76, 188, 73, 7, 65, 12, 170, 115, 163, 40, 149, 156, 209, 2, 80, 96, 169, 230, 9, 97, 157, 8, 141, 153,
			142, 144, 227, 201, 220, 154, 216, 108, 108, 112, 18, 29, 217, 99, 154, 20, 107, 101, 82, 252, 245, 202,
			171, 204, 10, 3, 72, 24, 69, 207, 228, 219, 227, 167, 214, 108, 255, 225, 7, 148, 106, 137, 241, 17, 215,
			142, 149, 194, 168, 128, 86, 28, 35, 29, 212, 34, 133, 167, 163, 158, 119, 191, 239, 165, 75, 255, 87, 212,
			88, 33, 157, 15},
		Comment: "comment",
	}

	mockCardStorage.EXPECT().GetByID(gomock.Any(), "1").Return(data, nil)

	_, err = cardService.GetCardByID(context.Background(), "1", "1")
	assert.NoError(t, err, "unexpected error fetching card")
}

func TestGetCardByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	cardService := NewCardService(mockCardStorage, encryptionKey, l)

	mockCardStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, pgx.ErrNoRows)

	_, err = cardService.GetCardByID(context.Background(), "1", "1")
	assert.ErrorAs(t, err, &e.NotFoundError{}, "expected not found error")
}

func TestGetCardByID_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	cardService := NewCardService(mockCardStorage, encryptionKey, l)

	mockCardStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, fmt.Errorf("unexpected internal error"))

	_, err = cardService.GetCardByID(context.Background(), "1", "1")
	assert.Error(t, err, "expected internal error")
}

func TestGetAllCards_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	cardService := NewCardService(mockCardStorage, encryptionKey, l)

	data := []model.Data{
		{
			ID:     "1",
			UserID: "1",
			Type:   model.CardType,
			SensitiveData: []byte{195, 64, 136, 52, 27, 89, 5, 77, 31, 193, 7, 121, 200, 46, 238, 176, 14, 246, 65, 97, 157,
				34, 76, 188, 73, 7, 65, 12, 170, 115, 163, 40, 149, 156, 209, 2, 80, 96, 169, 230, 9, 97, 157, 8, 141, 153,
				142, 144, 227, 201, 220, 154, 216, 108, 108, 112, 18, 29, 217, 99, 154, 20, 107, 101, 82, 252, 245, 202,
				171, 204, 10, 3, 72, 24, 69, 207, 228, 219, 227, 167, 214, 108, 255, 225, 7, 148, 106, 137, 241, 17, 215,
				142, 149, 194, 168, 128, 86, 28, 35, 29, 212, 34, 133, 167, 163, 158, 119, 191, 239, 165, 75, 255, 87, 212,
				88, 33, 157, 15},
			Comment: "comment",
		},
		{
			ID:     "2",
			UserID: "1",
			Type:   model.CardType,
			SensitiveData: []byte{195, 64, 136, 52, 27, 89, 5, 77, 31, 193, 7, 121, 200, 46, 238, 176, 14, 246, 65, 97, 157,
				34, 76, 188, 73, 7, 65, 12, 170, 115, 163, 40, 149, 156, 209, 2, 80, 96, 169, 230, 9, 97, 157, 8, 141, 153,
				142, 144, 227, 201, 220, 154, 216, 108, 108, 112, 18, 29, 217, 99, 154, 20, 107, 101, 82, 252, 245, 202,
				171, 204, 10, 3, 72, 24, 69, 207, 228, 219, 227, 167, 214, 108, 255, 225, 7, 148, 106, 137, 241, 17, 215,
				142, 149, 194, 168, 128, 86, 28, 35, 29, 212, 34, 133, 167, 163, 158, 119, 191, 239, 165, 75, 255, 87, 212,
				88, 33, 157, 15},
			Comment: "comment",
		},
	}

	mockCardStorage.EXPECT().GetAll(gomock.Any(), "1", model.CardType).Return(data, nil)

	cards, err := cardService.GetAllCards(context.Background(), "1")
	assert.NoError(t, err, "unexpected error fetching all cards data")
	assert.Len(t, cards, len(data), "should return two cards data")
}

func TestGetAllCards_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	cardService := NewCardService(mockCardStorage, encryptionKey, l)

	mockCardStorage.EXPECT().GetAll(gomock.Any(), "1", model.CardType).Return([]model.Data{},
		fmt.Errorf("unexpected internal error"))

	cards, err := cardService.GetAllCards(context.Background(), "1")
	assert.Error(t, err, "expected internal error")
	assert.Empty(t, cards, "expected no cards data")
}
