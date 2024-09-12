package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/rivo/tview"
	"google.golang.org/grpc/status"
)

// AddView creates and displays a menu for adding user data.
//
// It initializes a new menu with the following menu items:
// - "Add account data: login password pair" with the shortcut 'a'.
// - "Add card data: payment card details" with the shortcut 'c'.
// - "Add text data: text note" with the shortcut 't'.
// - "Add file data: upload file" with the shortcut 'f'.
// - "Back" with the shortcut 'b'.
func (c *ClientImpl) AddView() tview.Primitive {
	menu := tview.NewList()
	menu.AddItem("Add account data: login password pair", "", 'a', func() { c.AddAccountView() }).
		AddItem("Add card data: payment card details", "", 'c', func() { c.AddCardView() }).
		AddItem("Add text data: text note", "", 't', func() { c.AddTextView() }).
		AddItem("Add file data: upload file", "", 'f', func() { c.AddFileView() }).
		AddItem("Back", "", 'b', func() { c.MainView() })

	selectView(c.app, menu)

	return menu
}

// AddAccountView creates and displays a form for adding user account data.
func (c *ClientImpl) AddAccountView() tview.Primitive {
	form := tview.NewForm()
	form.AddInputField("Login", "", 20, nil, nil).
		AddInputField("Password", "", 20, nil, nil).
		AddInputField("Comment", "", 20, nil, nil).
		AddButton("Add", func() { addAccountHandler(c, c.accountService, form) }).
		AddButton("Cancel", func() { c.MainView() })

	selectView(c.app, form)

	return form
}

// addAccountHandler processes the form submit for adding user account data.
func addAccountHandler(c Client, accountService AccountService, form *tview.Form) {
	login := form.GetFormItemByLabel("Login").(*tview.InputField).GetText()
	password := form.GetFormItemByLabel("Password").(*tview.InputField).GetText()
	comment := form.GetFormItemByLabel("Comment").(*tview.InputField).GetText()

	err := accountService.CreateAccount(login, password, comment)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowInfoModal(fmt.Sprintf("Error adding account data: %s", err.Error()), form)
		} else {
			errMsg := st.Message()
			c.ShowInfoModal(fmt.Sprintf("Error adding account data: %s", errMsg), form)
		}
	} else {
		c.ShowInfoModal("Success added account data", form)
		c.AddAccountView()
	}
}

// AddCardView creates and displays a form for adding user card data.
func (c *ClientImpl) AddCardView() tview.Primitive {
	form := tview.NewForm()

	form.AddInputField("Number", "", 20, nil, nil).
		AddInputField("Owner", "", 20, nil, nil).
		AddInputField("Expiry", "", 20, nil, nil).
		AddInputField("CVC", "", 20, nil, nil).
		AddInputField("PIN", "", 20, nil, nil).
		AddInputField("Comment", "", 20, nil, nil).
		AddButton("Add", func() { addCardHandler(c, c.cardService, form) }).
		AddButton("Cancel", func() { c.MainView() })

	selectView(c.app, form)

	return form
}

// addCardHandler processes the form submit for adding user card data.
func addCardHandler(c Client, cardService CardService, form *tview.Form) {
	number := form.GetFormItemByLabel("Number").(*tview.InputField).GetText()
	owner := form.GetFormItemByLabel("Owner").(*tview.InputField).GetText()
	expiry := form.GetFormItemByLabel("Expiry").(*tview.InputField).GetText()
	cvc := form.GetFormItemByLabel("CVC").(*tview.InputField).GetText()
	pin := form.GetFormItemByLabel("PIN").(*tview.InputField).GetText()
	comment := form.GetFormItemByLabel("Comment").(*tview.InputField).GetText()

	err := cardService.CreateCard(number, owner, expiry, cvc, pin, comment)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowInfoModal(fmt.Sprintf("Error adding card data: %s", err.Error()), form)
		} else {
			errMsg := st.Message()
			c.ShowInfoModal(fmt.Sprintf("Error adding card data: %s", errMsg), form)
		}
	} else {
		c.ShowInfoModal("Success added card data", form)
		c.AddCardView()
	}
}

// AddTextView creates and displays a form for adding user text data.
func (c *ClientImpl) AddTextView() tview.Primitive {
	form := tview.NewForm()
	form.AddInputField("Text", "", 40, nil, nil).
		AddInputField("Comment", "", 20, nil, nil).
		AddButton("Add", func() { addTextHandler(c, c.textService, form) }).
		AddButton("Cancel", func() { c.MainView() })

	selectView(c.app, form)

	return form
}

// addTextHandler processes the form submission for adding user text data.
func addTextHandler(c Client, textService TextService, form *tview.Form) {
	text := form.GetFormItemByLabel("Text").(*tview.InputField).GetText()
	comment := form.GetFormItemByLabel("Comment").(*tview.InputField).GetText()

	err := textService.CreateText(text, comment)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowInfoModal(fmt.Sprintf("Error adding text data: %s", err.Error()), form)
		} else {
			errMsg := st.Message()
			c.ShowInfoModal(fmt.Sprintf("Error adding text data: %s", errMsg), form)
		}
	} else {
		c.ShowInfoModal("Success added text data", form)
		c.AddTextView()
	}
}

// AddFileView creates and displays a form for upload user file to the server.
func (c *ClientImpl) AddFileView() tview.Primitive {
	form := tview.NewForm()

	form.AddInputField("File path", "", 40, nil, nil).
		AddInputField("Comment", "", 40, nil, nil).
		AddButton("Upload", func() { uploadFileHandler(c, c.fileService, form) }).
		AddButton("Cancel", func() { c.MainView() })

	selectView(c.app, form)

	return form
}

// uploadFileHandler processes the form submit for upload user file to the server.
func uploadFileHandler(c Client, fileService FileService, form *tview.Form) {
	filePath := form.GetFormItemByLabel("File path").(*tview.InputField).GetText()
	comment := form.GetFormItemByLabel("Comment").(*tview.InputField).GetText()

	ctx, cancel := context.WithCancel(context.Background())

	progressText := tview.NewTextView().
		SetText("Uploading... 0%").
		SetChangedFunc(func() { c.UpdateDraw() })

	progressBar := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("Uploading file:"), 1, 1, false).
		AddItem(progressText, 1, 1, false)

	modal := tview.NewModal().
		SetText("Uploading file...").
		AddButtons([]string{"Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Cancel" {
				cancel()
				c.SelectView(form)
			}
		})

	pages := tview.NewPages().
		AddPage("progressBar", progressBar, true, true).
		AddPage("modal", modal, true, true)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := fileService.UploadFile(ctx, filePath, comment, func(progress float64) {
			c.QueueUpdateDraw(func() { progressText.SetText(fmt.Sprintf("Uploading... %.2f%%", progress)) })
		})
		if err != nil {
			c.QueueUpdateDraw(func() { c.ShowInfoModal("Failed to upload file", form) })
		} else {
			c.QueueUpdateDraw(func() {
				clearForm(form)
				c.ShowInfoModal("File uploaded successfully", form)
			})
		}
	}()

	c.SelectView(pages)
	wg.Wait()
}
