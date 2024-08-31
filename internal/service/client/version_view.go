package client

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// VersionView displays the application version and build date, and back button to return to the main menu.
func (c *Client) VersionView() {
	header := tview.NewTextView().
		SetText("About application").
		SetTextColor(tcell.ColorYellow).
		SetTextAlign(tview.AlignLeft)

	aboutText := fmt.Sprintf("Version: %s\nBuild Date: %s", c.config.BuildVersion, c.config.BuildDate)
	aboutBlock := tview.NewTextView().
		SetText(aboutText).
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)

	buttons := tview.NewForm().
		AddButton("Back", func() { c.MainView() })

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(header, 3, 1, false).
		AddItem(aboutBlock, 0, 1, false).
		AddItem(buttons, 3, 1, true)

	selectView(c.app, flex)
}
