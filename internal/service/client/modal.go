package client

import (
	"github.com/rivo/tview"
)

// ShowInfoModal displays an informational modal dialog with "OK" button.
// When the "OK" button is pressed, the application returns to the current view.
func (c *Client) ShowInfoModal(text string, currentView tview.Primitive) {
	modal := tview.NewModal().
		SetText(text).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				selectView(c.app, currentView)
			}
		})

	selectView(c.app, modal)
}

// ShowConfirmModal displays a confirmation modal dialog with "Yes" and "No" buttons.
// If "Yes" is pressed, the provided handler function is executed.
// If "No" is pressed, the application returns to the current view.
func (c *Client) ShowConfirmModal(text string, currentView tview.Primitive, handler func()) {
	modal := tview.NewModal().
		SetText(text).
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				handler()
			} else {
				selectView(c.app, currentView)
			}
		})

	selectView(c.app, modal)
}

// ShowRetryModal displays a retry modal dialog with "Retry" and "Cancel" buttons.
// If "Retry" is pressed, the application switches to the current view.
// If "Cancel" is pressed, the application switches back to the previous view.
func (c *Client) ShowRetryModal(text string, currentView tview.Primitive, previousView tview.Primitive) {
	modal := tview.NewModal().
		SetText(text).
		AddButtons([]string{"Retry", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Retry" {
				selectView(c.app, currentView)
			} else {
				selectView(c.app, previousView)
			}
		})

	selectView(c.app, modal)
}
