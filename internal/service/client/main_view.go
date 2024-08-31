package client

import (
	"github.com/rivo/tview"
)

// MainView creates and displays the main menu for application.
//
// It initializes a new menu with the following menu items:
// - "Add new data: pair login - password, payment card details, text note, file" with the shortcut 'a'
// - "Viewing data: pairs login - password, payment cards details, text notes, files" with the shortcut 'v'
// - "Logout" with the shortcut 'q'
func (c *Client) MainView() {
	menu := tview.NewList()
	menu.AddItem("Add new data: pair login - password, payment card details, text note, file", "",
		'a', c.AddView).
		AddItem("Viewing data: pairs login- password, payment cards details, text notes, files", "",
			'v', c.DataView).
		AddItem("Application version info", "", 'i', c.VersionView).
		AddItem("Logout", "", 'q', c.AuthView)

	selectView(c.app, menu)
}
