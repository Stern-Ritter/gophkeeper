package client

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
)

func TestShowInfoModal(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	app := tview.NewApplication()
	cfg := &config.ClientConfig{}
	client := &ClientImpl{
		app:    app,
		config: cfg,
	}

	client.SetApp(app)
	currentView := tview.NewTextView().SetText("Current view")
	client.ShowInfoModal("Info message", currentView)

	focusedElement, ok := app.GetFocus().(*tview.Button)
	assert.True(t, ok, "OK button should be in focus")
	assert.NotNil(t, focusedElement, "OK button should be in focus")
	assert.Equal(t, focusedElement.GetLabel(), "OK", "OK button should be in focus")
}

func TestShowConfirmModal(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	app := tview.NewApplication()
	cfg := &config.ClientConfig{}
	client := &ClientImpl{
		app:    app,
		config: cfg,
	}

	client.SetApp(app)
	currentView := tview.NewTextView().SetText("Current view")
	client.ShowConfirmModal("Info message", currentView, func() {})

	focusedElement, ok := app.GetFocus().(*tview.Button)
	assert.True(t, ok, "Yes button should be in focus")
	assert.NotNil(t, focusedElement, "Yes button should be in focus")
	assert.Equal(t, focusedElement.GetLabel(), "Yes", "Yes button should be in focus")
}

func TestShowRetryModal(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	app := tview.NewApplication()
	cfg := &config.ClientConfig{}
	client := &ClientImpl{
		app:    app,
		config: cfg,
	}

	client.SetApp(app)
	previousView := tview.NewTextView().SetText("Previous view")
	currentView := tview.NewTextView().SetText("Current view")
	client.ShowRetryModal("Info message", currentView, previousView)

	focusedElement, ok := app.GetFocus().(*tview.Button)
	assert.True(t, ok, "Retry button should be in focus")
	assert.NotNil(t, focusedElement, "Retry button should be in focus")
	assert.Equal(t, focusedElement.GetLabel(), "Retry", "Retry button should be in focus")
}
