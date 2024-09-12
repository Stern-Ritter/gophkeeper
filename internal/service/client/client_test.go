package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
)

func TestSetAuthService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := NewMockAuthService(ctrl)

	cfg := &config.ClientConfig{}
	c := NewClient(cfg)

	c.SetAuthService(mockAuthService)

	assert.Equal(t, mockAuthService, c.GetAuthService(), "auth service should be set")
}

func TestSetAccountService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountService := NewMockAccountService(ctrl)

	cfg := &config.ClientConfig{}
	c := NewClient(cfg)

	c.SetAccountService(mockAccountService)

	assert.Equal(t, mockAccountService, c.GetAccountService(), "account service should be set")
}

func TestSetCardService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardService := NewMockCardService(ctrl)

	cfg := &config.ClientConfig{}
	c := NewClient(cfg)

	c.SetCardService(mockCardService)

	assert.Equal(t, mockCardService, c.GetCardService(), "card service should be set")
}

func TestSetTextService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTextService := NewMockTextService(ctrl)

	cfg := &config.ClientConfig{}
	c := NewClient(cfg)

	c.SetTextService(mockTextService)

	assert.Equal(t, mockTextService, c.GetTextService(), "text service should be set")
}

func TestSetFileService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFileService := NewMockFileService(ctrl)

	cfg := &config.ClientConfig{}
	c := NewClient(cfg)

	c.SetFileService(mockFileService)

	assert.Equal(t, mockFileService, c.GetFileService(), "fileservice should be set")
}

func TestClient_SetApp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockApp := NewMockApplication(ctrl)

	cfg := &config.ClientConfig{}
	c := NewClient(cfg)

	c.SetApp(mockApp)

	assert.Equal(t, mockApp, c.GetApp(), "app should be set")
}
