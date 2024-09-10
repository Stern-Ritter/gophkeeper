package client

import (
	"context"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"google.golang.org/grpc/status"

	"github.com/Stern-Ritter/gophkeeper/internal/utils"
)

const (
	downloadBtnText = "[green::bl]DOWNLOAD" // downloadBtnText represents the text for the download button.
	deleteBtnText   = "[yellow::bl]DELETE"  // deleteBtnText represents the text for the delete button
)

// DataView creates and displays a menu for viewing user data.
//
// It initializes a new menu with the following menu items:
// - "View accounts data: login password pairs" with the shortcut 'a'.
// - "View cards data: payment cards details" with the shortcut 'c'.
// - "View text data: text notes" with the shortcut 't'.
// - "View file data: uploaded files" with the shortcut 'f'.
// - "Back" with the shortcut 'b'.
func (c *ClientImpl) DataView() tview.Primitive {
	menu := tview.NewList()
	menu.AddItem("View accounts data: login password pairs", "", 'a', func() { c.AccountsView(menu) }).
		AddItem("View cards data: payment cards details", "", 'c', func() { c.CardsView(menu) }).
		AddItem("View text data: text notes", "", 't', func() { c.TextsView(menu) }).
		AddItem("View file data: uploaded files", "", 'f', func() { c.FilesView(menu) }).
		AddItem("Back", "", 'b', func() { c.MainView() })

	selectView(c.app, menu)

	return menu
}

// AccountsView displays a table of accounts data, including login, password, comment, and a delete button.
func (c *ClientImpl) AccountsView(previousView tview.Primitive) tview.Primitive {
	accounts, err := c.accountService.GetAllAccounts()
	if err != nil {
		c.ShowInfoModal(fmt.Sprintf("Get accounts data error: %s", err.Error()), previousView)
	}

	table := tview.NewTable().SetSelectable(true, true).SetBorders(true)
	setTableHeader(table, []string{"Login", "Password", "Comment", "Delete"})
	addTableListeners(table, c.app, previousView)

	for i, account := range accounts {
		row := i + 1
		column := getColumnCounter()
		table.SetCell(row, column(), tview.NewTableCell(account.Login))
		table.SetCell(row, column(), tview.NewTableCell(account.Password))
		table.SetCell(row, column(), tview.NewTableCell(account.Comment))
		table.SetCell(row, column(), newClickableCell(deleteBtnText, func() { c.deleteAccountHandler(account.Id, table, previousView) }))
	}

	selectView(c.app, table)

	return table
}

// deleteAccountHandler handles the deletion of account data.
// It shows a confirmation modal and deletes the account if confirmed.
func (c *ClientImpl) deleteAccountHandler(accountID string, currentView tview.Primitive, previousView tview.Primitive) {
	c.ShowConfirmModal("Are you sure you want to delete this account data?", currentView,
		func() {
			err := c.accountService.DeleteAccount(accountID)
			if err != nil {
				st, ok := status.FromError(err)
				if !ok {
					c.ShowInfoModal(fmt.Sprintf("Failed to delete account data: %s", err.Error()), currentView)
				} else {
					errMsg := st.Message()
					c.ShowInfoModal(fmt.Sprintf("Failed to delete account data: %s", errMsg), currentView)
				}
			}
			c.ShowInfoModal("Account data deleted successfully", currentView)
			c.AccountsView(previousView)
		})
}

// CardsView displays a table of cards data, including number, owner, expiry, cvc, pin, comment and a delete button.
func (c *ClientImpl) CardsView(previousView tview.Primitive) tview.Primitive {
	cards, err := c.cardService.GetAllCards()
	if err != nil {
		c.ShowInfoModal(fmt.Sprintf("Get cards data error: %s", err.Error()), previousView)
	}

	table := tview.NewTable().SetSelectable(true, true).SetBorders(true)
	setTableHeader(table, []string{"Number", "Owner", "Expiry", "CVC", "PIN", "Comment", "Delete"})
	addTableListeners(table, c.app, previousView)

	for i, card := range cards {
		row := i + 1
		column := getColumnCounter()
		table.SetCell(row, column(), tview.NewTableCell(card.Number))
		table.SetCell(row, column(), tview.NewTableCell(card.Owner))
		table.SetCell(row, column(), tview.NewTableCell(card.Expiry))
		table.SetCell(row, column(), tview.NewTableCell(card.Cvc))
		table.SetCell(row, column(), tview.NewTableCell(card.Pin))
		table.SetCell(row, column(), tview.NewTableCell(card.Comment))
		table.SetCell(row, column(), newClickableCell(deleteBtnText, func() { c.deleteCardHandler(card.Id, table, previousView) }))
	}

	selectView(c.app, table)

	return table
}

// deleteCardHandler handles the deletion of card data.
// It shows a confirmation modal and deletes the card if confirmed.
func (c *ClientImpl) deleteCardHandler(cardID string, currentView tview.Primitive, previousView tview.Primitive) {
	c.ShowConfirmModal("Are you sure you want to delete this card data?", currentView,
		func() {
			err := c.cardService.DeleteCard(cardID)
			if err != nil {
				st, ok := status.FromError(err)
				if !ok {
					c.ShowInfoModal(fmt.Sprintf("Failed to delete card data: %s", err.Error()), currentView)
				} else {
					errMsg := st.Message()
					c.ShowInfoModal(fmt.Sprintf("Failed to delete card data: %s", errMsg), currentView)
				}
			}
			c.ShowInfoModal("Card data deleted successfully", currentView)
			c.CardsView(previousView)
		})
}

// TextsView displays a table of texts data, including text, comment and a delete button.
func (c *ClientImpl) TextsView(previousView tview.Primitive) tview.Primitive {
	texts, err := c.textService.GetAllTexts()
	if err != nil {
		c.ShowInfoModal(fmt.Sprintf("Get texts data error: %s", err.Error()), previousView)
	}

	table := tview.NewTable().SetSelectable(true, true).SetBorders(true)
	setTableHeader(table, []string{"Text", "Comment", "Delete"})
	addTableListeners(table, c.app, previousView)

	for i, text := range texts {
		row := i + 1
		column := getColumnCounter()
		table.SetCell(row, column(), tview.NewTableCell(text.Text))
		table.SetCell(row, column(), tview.NewTableCell(text.Comment))
		table.SetCell(row, column(), newClickableCell(deleteBtnText, func() { c.deleteTextHandler(text.Id, table, previousView) }))
	}

	selectView(c.app, table)

	return table
}

// deleteTextHandler handles the deletion of text data.
// It shows a confirmation modal and deletes the text if confirmed.
func (c *ClientImpl) deleteTextHandler(textID string, currentView tview.Primitive, previousView tview.Primitive) {
	c.ShowConfirmModal("Are you sure you want to delete this text data?", currentView,
		func() {
			err := c.textService.DeleteText(textID)
			if err != nil {
				st, ok := status.FromError(err)
				if !ok {
					c.ShowInfoModal(fmt.Sprintf("Failed to delete text data: %s", err.Error()), currentView)
				} else {
					errMsg := st.Message()
					c.ShowInfoModal(fmt.Sprintf("Failed to delete text data: %s", errMsg), currentView)
				}
			}
			c.ShowInfoModal("Text data deleted successfully", currentView)
			c.TextsView(previousView)
		})
}

// FilesView displays a table of file data, including file name, size, comment, and options to download or delete file.
func (c *ClientImpl) FilesView(previousView tview.Primitive) tview.Primitive {
	files, err := c.fileService.GetAllFiles()
	if err != nil {
		c.ShowInfoModal(fmt.Sprintf("Get files error: %s", err.Error()), previousView)
	}

	table := tview.NewTable().SetSelectable(true, true).SetBorders(true)
	setTableHeader(table, []string{"Name", "Size", "Comment", "Download", "Delete"})
	addTableListeners(table, c.app, previousView)

	for i, file := range files {
		row := i + 1
		column := getColumnCounter()
		table.SetCell(row, column(), tview.NewTableCell(file.Name))
		table.SetCell(row, column(), tview.NewTableCell(utils.FormatBytes(file.Size)))
		table.SetCell(row, column(), tview.NewTableCell(file.Comment))
		table.SetCell(row, column(), newClickableCell(downloadBtnText, func() { c.downloadFileHandler(file.Id, table) }))
		table.SetCell(row, column(), newClickableCell(deleteBtnText, func() { c.deleteFileHandler(file.Id, table, previousView) }))
	}

	selectView(c.app, table)

	return table
}

// deleteFileHandler handles the deletion of a file. It shows a confirmation modal and deletes the file if confirmed.
func (c *ClientImpl) deleteFileHandler(fileID string, currentView tview.Primitive, previousView tview.Primitive) {
	c.ShowConfirmModal("Are you sure you want to delete this file?", currentView,
		func() {
			err := c.fileService.DeleteFile(fileID)
			if err != nil {
				st, ok := status.FromError(err)
				if !ok {
					c.ShowInfoModal(fmt.Sprintf("Failed to delete file: %s", err.Error()), currentView)
				} else {
					errMsg := st.Message()
					c.ShowInfoModal(fmt.Sprintf("Failed to delete file: %s", errMsg), currentView)
				}
			}
			c.ShowInfoModal("File deleted successfully", currentView)
			c.FilesView(previousView)
		})
}

// downloadFileHandler displays a form for entering the directory path to download a file.
// It starts the download when the "Download" button is clicked.
func (c *ClientImpl) downloadFileHandler(fileID string, currentView tview.Primitive) {
	form := tview.NewForm()
	form.AddInputField("Directory path", "", 40, nil, nil).
		AddButton("Download", func() {
			dirPath := form.GetFormItemByLabel("Directory path").(*tview.InputField).GetText()
			c.downloadFile(fileID, dirPath, currentView)
		}).
		AddButton("Cancel", func() { selectView(c.app, currentView) })

	selectView(c.app, form)
}

// downloadFile initiates the download of a file and displays a progress indicator.
// It uses a context to manage cancellation and provides a progress view with a modal dialog.
// The progress is shown as a percentage, and the user can cancel the download.
// Once the download is complete, a message indicating success or failure is displayed.
func (c *ClientImpl) downloadFile(fileID string, dirPath string, currentView tview.Primitive) {
	ctx, cancel := context.WithCancel(context.Background())

	progressText := tview.NewTextView().
		SetText("Downloading... 0%").
		SetChangedFunc(func() { c.app.Draw() })

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("Downloading file:"), 1, 1, false).
		AddItem(progressText, 1, 1, false)

	modal := tview.NewModal().
		SetText("Downloading file...").
		AddButtons([]string{"Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Cancel" {
				cancel()
				selectView(c.app, currentView)
			}
		})

	pages := tview.NewPages().
		AddPage("flex", flex, true, true).
		AddPage("modal", modal, true, true)

	go func() {
		err := c.fileService.DownloadFile(ctx, fileID, dirPath, func(progress float64) {
			c.app.QueueUpdateDraw(func() { progressText.SetText(fmt.Sprintf("Downloading... %.2f%%", progress)) })
		})
		if err != nil {
			c.app.QueueUpdateDraw(func() {
				st, ok := status.FromError(err)
				if !ok {
					c.ShowInfoModal(fmt.Sprintf("Failed to download file: %s", err.Error()), currentView)
				} else {
					errMsg := st.Message()
					c.ShowInfoModal(fmt.Sprintf("Failed to download file: %s", errMsg), currentView)
				}
			})
		} else {
			c.app.QueueUpdateDraw(func() { c.ShowInfoModal("File downloaded successfully", currentView) })
		}
	}()

	selectView(c.app, pages)
}

// addTableListeners sets up listeners for a table.
// It handles Enter key presses for triggering actions on clickable cells and
// handles Esc key presses for returning to the previous view.
func addTableListeners(table *tview.Table, app *tview.Application, previousView tview.Primitive) {
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			row, column := table.GetSelection()
			cell := table.GetCell(row, column)
			switch cell.Text {
			case downloadBtnText, deleteBtnText:
				if cell.Clicked != nil {
					cell.Clicked()
				}
			}
		case tcell.KeyEsc:
			app.SetRoot(previousView, true)
			return nil
		}
		return event
	})
}
