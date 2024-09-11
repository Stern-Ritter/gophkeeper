package client

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/client"
)

func TestClient_Run(t *testing.T) {
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
