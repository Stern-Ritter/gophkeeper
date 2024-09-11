package client

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
	logger "github.com/Stern-Ritter/gophkeeper/internal/logger/client"
)

func TestClient_Run(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := NewMockClient(mockCtrl)

	gomock.InOrder(
		mockClient.EXPECT().SetAuthService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetAccountService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetCardService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetTextService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetFileService(gomock.Any()).Times(1),
		mockClient.EXPECT().SetApp(gomock.Any()).Times(1),
		mockClient.EXPECT().AuthView().Times(1),
	)

	cfg := &config.ClientConfig{ServerURL: "localhost:8080"}

	l, err := logger.Initialize("error")
	require.NoError(t, err, "Error init logger")

	resultChan := make(chan error, 1)

	go func() {
		err := Run(mockClient, cfg, l)
		resultChan <- err
	}()

	select {
	case err := <-resultChan:
		require.NoError(t, err, "Expected client run be succeed")
	case <-time.After(1 * time.Second):
	}
}
