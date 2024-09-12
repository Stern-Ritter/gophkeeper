package client

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
)

func TestVersionView(t *testing.T) {
	app := tview.NewApplication()

	cfg := &config.ClientConfig{
		BuildVersion: "1.0.0",
		BuildDate:    "2024-09-08",
	}

	client := NewClient(cfg)
	client.SetApp(app)

	element := client.VersionView()

	flex := element.(*tview.Flex)
	header := flex.GetItem(0).(*tview.TextView)
	text := header.GetText(false)
	assert.Contains(t, text, "About application", "Header text should be 'About application'")

	aboutBlock := flex.GetItem(1).(*tview.TextView)
	aboutText := aboutBlock.GetText(false)
	expectedText := "Version: 1.0.0\nBuild Date: 2024-09-08"
	assert.Equal(t, expectedText, aboutText, "About text should match the config values")

	form := flex.GetItem(2).(*tview.Form)
	button := form.GetButton(0)
	assert.NotNil(t, button, "Back button should be present")
	assert.Equal(t, "Back", button.GetLabel(), "Back button label should be 'Back'")
}
