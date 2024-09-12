package client

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
)

func TestMainView(t *testing.T) {
	app := tview.NewApplication()

	cfg := &config.ClientConfig{}

	client := NewClient(cfg)
	client.SetApp(app)

	client.MainView()

	menu := app.GetFocus().(*tview.List)
	assert.NotNil(t, menu, "Menu should be focused")

	assert.Equal(t, 4, menu.GetItemCount(), "Menu should have 4 items")

	firstItemText, _ := menu.GetItemText(0)
	assert.Contains(t, firstItemText, "Add new data", "First menu item should be 'Add new data: pair login - password, payment card details, text note, file'")

	secondItemText, _ := menu.GetItemText(1)
	assert.Contains(t, secondItemText, "Viewing data", "Second menu item should be 'Viewing data: pairs login- password, payment cards details, text notes, files'")

	thirdItemText, _ := menu.GetItemText(2)
	assert.Contains(t, thirdItemText, "Application version info", "Third menu item should be 'Application version info'")

	fourthItemText, _ := menu.GetItemText(3)
	assert.Contains(t, fourthItemText, "Logout", "Fourth menu item should be 'Logout'")
}
