package client

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/client"
)

func TestRun_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := NewMockClient(mockCtrl)
	mockApplication := NewMockApplication(mockCtrl)

	gomock.InOrder(
		mockClient.EXPECT().SetAuthService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetAccountService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetCardService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetTextService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetFileService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetApp(gomock.Any()).Times(1),
		mockClient.EXPECT().AuthView().Times(1),
		mockApplication.EXPECT().Run().Return(nil).Times(1),
	)

	cfg := &config.ClientConfig{ServerURL: ":8080"}

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	err = Run(mockClient, mockApplication, cfg, l)
	require.NoError(t, err, "unexpected error run application")
}

func TestRun_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := NewMockClient(mockCtrl)
	mockApplication := NewMockApplication(mockCtrl)

	err := errors.New("run application error")
	gomock.InOrder(
		mockClient.EXPECT().SetAuthService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetAccountService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetCardService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetTextService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetFileService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetApp(gomock.Any()).Times(1),
		mockClient.EXPECT().AuthView().Times(1),
		mockApplication.EXPECT().Run().Return(err).Times(1),
	)

	cfg := &config.ClientConfig{ServerURL: ":8080"}

	l, err := logger.Initialize("error")
	require.NoError(t, err, "unexpected error init logger")

	err = Run(mockClient, mockApplication, cfg, l)
	require.Error(t, err, "expected error run application")
}
