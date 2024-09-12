package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	e "github.com/Stern-Ritter/gophkeeper/internal/errors"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/server"
	"github.com/Stern-Ritter/gophkeeper/internal/model"
)

func TestCreateText_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	text := model.Text{
		ID:      "1",
		UserID:  "1",
		Text:    "sensitive text",
		Comment: "comment",
	}

	mockTextStorage.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	err = textService.CreateText(context.Background(), text)
	assert.NoError(t, err, "expected no error creating text")
}

func TestCreateText_EncryptionFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "invalid key"
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	text := model.Text{
		ID:      "1",
		UserID:  "1",
		Text:    "sensitive text",
		Comment: "comment",
	}

	err = textService.CreateText(context.Background(), text)
	assert.Error(t, err, "expected error encrypting sensitive data")
}

func TestCreateText_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	text := model.Text{
		ID:      "1",
		UserID:  "1",
		Text:    "sensitive text",
		Comment: "comment",
	}

	mockTextStorage.EXPECT().Create(gomock.Any(), gomock.Any()).Return(fmt.Errorf("failed to save text"))

	err = textService.CreateText(context.Background(), text)
	assert.Error(t, err, "expected error saving text")
}

func TestDeleteText_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	data := model.Data{
		ID:            "1",
		UserID:        "1",
		Type:          model.TextType,
		SensitiveData: make([]byte, 0),
		Comment:       "comment",
	}

	mockTextStorage.EXPECT().GetByID(gomock.Any(), "1").Return(data, nil)
	mockTextStorage.EXPECT().Delete(gomock.Any(), "1").Return(nil)

	err = textService.DeleteText(context.Background(), "1", "1")
	assert.NoError(t, err, "unexpected error deleting text")
}

func TestDeleteText_Forbidden(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	data := model.Data{
		ID:            "1",
		UserID:        "1",
		Type:          model.TextType,
		SensitiveData: make([]byte, 0),
		Comment:       "comment",
	}

	mockTextStorage.EXPECT().GetByID(gomock.Any(), "1").Return(data, nil)

	err = textService.DeleteText(context.Background(), "2", "1")
	assert.ErrorAs(t, err, &e.ForbiddenError{}, "expected forbidden error")
}

func TestDeleteText_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	mockTextStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, pgx.ErrNoRows)

	err = textService.DeleteText(context.Background(), "1", "1")
	assert.ErrorAs(t, err, &e.NotFoundError{}, "expected not found error")
}

func TestDeleteText_GetByIDInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	mockTextStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, fmt.Errorf("unexpected internal error"))

	err = textService.DeleteText(context.Background(), "1", "1")
	assert.Error(t, err, "expected internal error")
}

func TestDeleteText_DeleteError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	data := model.Data{
		ID:            "1",
		UserID:        "1",
		Type:          model.TextType,
		SensitiveData: make([]byte, 0),
		Comment:       "comment",
	}

	mockTextStorage.EXPECT().GetByID(gomock.Any(), "1").Return(data, nil)
	mockTextStorage.EXPECT().Delete(gomock.Any(), "1").Return(fmt.Errorf("failed to delete text"))

	err = textService.DeleteText(context.Background(), "1", "1")
	assert.Error(t, err, "expected error deleting text")
}

func TestGetTextByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	data := model.Data{
		ID:     "1",
		UserID: "1",
		Type:   model.TextType,
		SensitiveData: []byte{250, 10, 23, 249, 70, 187, 80, 165, 13, 29, 67, 52, 68, 71, 35, 253, 126, 168, 105, 25,
			170, 101, 180, 104, 149, 209, 182, 157, 85, 81, 4, 182, 115, 253, 111, 76, 58, 9, 32, 2, 157, 233, 220, 85,
			165, 99, 174, 229, 68, 112, 61, 98, 226},
		Comment: "comment",
	}

	mockTextStorage.EXPECT().GetByID(gomock.Any(), "1").Return(data, nil)

	_, err = textService.GetTextByID(context.Background(), "1", "1")
	assert.NoError(t, err, "unexpected error fetching text")
}

func TestGetTextByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	mockTextStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, pgx.ErrNoRows)

	_, err = textService.GetTextByID(context.Background(), "1", "1")
	assert.ErrorAs(t, err, &e.NotFoundError{}, "expected not found error")
}

func TestGetTextByID_DecryptionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := ""
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	data := model.Data{
		ID:     "1",
		UserID: "1",
		Type:   model.TextType,
		SensitiveData: []byte{250, 10, 23, 249, 70, 187, 80, 165, 13, 29, 67, 52, 68, 71, 35, 253, 126, 168, 105, 25,
			170, 101, 180, 104, 149, 209, 182, 157, 85, 81, 4, 182, 115, 253, 111, 76, 58, 9, 32, 2, 157, 233, 220, 85,
			165, 99, 174, 229, 68, 112, 61, 98, 226},
		Comment: "comment",
	}

	mockTextStorage.EXPECT().GetByID(gomock.Any(), "1").Return(data, nil)

	_, err = textService.GetTextByID(context.Background(), "1", "1")
	assert.Error(t, err, "expected error decrypting sensitive data")
}

func TestGetTextByID_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	mockTextStorage.EXPECT().GetByID(gomock.Any(), "1").Return(model.Data{}, fmt.Errorf("unexpected internal error"))

	_, err = textService.GetTextByID(context.Background(), "1", "1")
	assert.Error(t, err, "expected internal error")
}

func TestGetAllTexts_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	data := []model.Data{
		{
			ID:     "1",
			UserID: "1",
			SensitiveData: []byte{250, 10, 23, 249, 70, 187, 80, 165, 13, 29, 67, 52, 68, 71, 35, 253, 126, 168, 105, 25,
				170, 101, 180, 104, 149, 209, 182, 157, 85, 81, 4, 182, 115, 253, 111, 76, 58, 9, 32, 2, 157, 233, 220, 85,
				165, 99, 174, 229, 68, 112, 61, 98, 226},
		},
		{
			ID:     "2",
			UserID: "1",
			SensitiveData: []byte{250, 10, 23, 249, 70, 187, 80, 165, 13, 29, 67, 52, 68, 71, 35, 253, 126, 168, 105, 25,
				170, 101, 180, 104, 149, 209, 182, 157, 85, 81, 4, 182, 115, 253, 111, 76, 58, 9, 32, 2, 157, 233, 220, 85,
				165, 99, 174, 229, 68, 112, 61, 98, 226},
		},
	}

	mockTextStorage.EXPECT().GetAll(gomock.Any(), "1", model.TextType).Return(data, nil)

	texts, err := textService.GetAllTexts(context.Background(), "1")
	assert.NoError(t, err, "unexpected error fetching all texts data")
	assert.Len(t, texts, len(data), "should return two texts data")
}

func TestGetAllTexts_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := "HmQkWX1zTb6l3P8V8f3Eiw=="
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	mockTextStorage.EXPECT().GetAll(gomock.Any(), "1", model.TextType).Return([]model.Data{},
		fmt.Errorf("unexpected internal error"))

	texts, err := textService.GetAllTexts(context.Background(), "1")
	assert.Error(t, err, "expected internal error")
	assert.Empty(t, texts, "expected no texts data")
}

func TestGetAllTexts_DecryptionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextStorage := NewMockDataStorage(ctrl)

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	encryptionKey := ""
	textService := NewTextService(mockTextStorage, encryptionKey, l)

	data := []model.Data{
		{
			ID:     "1",
			UserID: "1",
			SensitiveData: []byte{250, 10, 23, 249, 70, 187, 80, 165, 13, 29, 67, 52, 68, 71, 35, 253, 126, 168, 105, 25,
				170, 101, 180, 104, 149, 209, 182, 157, 85, 81, 4, 182, 115, 253, 111, 76, 58, 9, 32, 2, 157, 233, 220, 85,
				165, 99, 174, 229, 68, 112, 61, 98, 226},
		},
	}

	mockTextStorage.EXPECT().GetAll(gomock.Any(), "1", model.TextType).Return(data, nil)

	texts, err := textService.GetAllTexts(context.Background(), "1")
	assert.Error(t, err, "expected error decrypting sensitive data")
	assert.Empty(t, texts, "expected no texts data")
}
