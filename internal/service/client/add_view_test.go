package client

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
)

func TestAddView(t *testing.T) {
	app := tview.NewApplication()

	cfg := &config.ClientConfig{}

	client := NewClient(cfg)
	client.SetApp(app)

	element := client.AddView()

	menu, ok := element.(*tview.List)
	assert.True(t, ok, "Should create a menu")
	assert.Equal(t, 5, menu.GetItemCount(), "Menu should have 4 items")

	firstItemText, _ := menu.GetItemText(0)
	assert.Contains(t, firstItemText, "Add account data: login password pair", "First menu item should be 'Add account data: login password pair'")

	secondItemText, _ := menu.GetItemText(1)
	assert.Contains(t, secondItemText, "Add card data: payment card details", "Second menu item should be 'Add card data: payment card details'")

	thirdItemText, _ := menu.GetItemText(2)
	assert.Contains(t, thirdItemText, "Add text data: text note", "Third menu item should be 'Add text data: text note'")

	fourthItemText, _ := menu.GetItemText(3)
	assert.Contains(t, fourthItemText, "Add file data: upload file", "Fourth menu item should be 'Add file data: upload file'")

	fifthItemText, _ := menu.GetItemText(4)
	assert.Contains(t, fifthItemText, "Back", "Fourth menu item should be 'Back'")
}

func TestAddAccountView(t *testing.T) {
	app := tview.NewApplication()

	cfg := &config.ClientConfig{}
	client := NewClient(cfg)
	client.SetApp(app)

	element := client.AddAccountView()

	form := element.(*tview.Form)
	assert.Equal(t, 3, form.GetFormItemCount(), "Form should have 3 input fields")

	loginField := form.GetFormItemByLabel("Login").(*tview.InputField)
	assert.NotNil(t, loginField, "Login input field should be in form")

	passwordField := form.GetFormItemByLabel("Password").(*tview.InputField)
	assert.NotNil(t, passwordField, "Password input field should be in form")

	commentField := form.GetFormItemByLabel("Comment").(*tview.InputField)
	assert.NotNil(t, commentField, "Comment input field should be in form")

	addButton := form.GetButton(0)
	assert.Equal(t, "Add", addButton.GetLabel(), "First button should be 'Add'")

	cancelButton := form.GetButton(1)
	assert.Equal(t, "Cancel", cancelButton.GetLabel(), "Second button should be 'Cancel'")
}

func TestAddCardView(t *testing.T) {
	app := tview.NewApplication()

	cfg := &config.ClientConfig{}
	client := NewClient(cfg)
	client.SetApp(app)

	element := client.AddCardView()

	form := element.(*tview.Form)
	assert.Equal(t, 6, form.GetFormItemCount(), "Form should have 6 input fields")

	numberField := form.GetFormItemByLabel("Number").(*tview.InputField)
	assert.NotNil(t, numberField, "Number input field should be in form")

	ownerField := form.GetFormItemByLabel("Owner").(*tview.InputField)
	assert.NotNil(t, ownerField, "Owner input field should be in form")

	expiryField := form.GetFormItemByLabel("Expiry").(*tview.InputField)
	assert.NotNil(t, expiryField, "Expiry input field should be in form")

	cvcField := form.GetFormItemByLabel("CVC").(*tview.InputField)
	assert.NotNil(t, cvcField, "CVC input field should be in form")

	pinField := form.GetFormItemByLabel("PIN").(*tview.InputField)
	assert.NotNil(t, pinField, "PIN input field should be in form")

	commentField := form.GetFormItemByLabel("Comment").(*tview.InputField)
	assert.NotNil(t, commentField, "Comment input field should be in form")

	addButton := form.GetButton(0)
	assert.Equal(t, "Add", addButton.GetLabel(), "First button should be 'Add'")

	cancelButton := form.GetButton(1)
	assert.Equal(t, "Cancel", cancelButton.GetLabel(), "Second button should be 'Cancel'")
}

func TestAddTextView(t *testing.T) {
	app := tview.NewApplication()

	cfg := &config.ClientConfig{}
	client := NewClient(cfg)
	client.SetApp(app)

	element := client.AddTextView()

	form := element.(*tview.Form)
	assert.Equal(t, 2, form.GetFormItemCount(), "Form should have 2 input fields")

	textField := form.GetFormItemByLabel("Text").(*tview.InputField)
	assert.NotNil(t, textField, "Text input field should be in form")

	commentField := form.GetFormItemByLabel("Comment").(*tview.InputField)
	assert.NotNil(t, commentField, "Comment input field should be in form")

	addButton := form.GetButton(0)
	assert.Equal(t, "Add", addButton.GetLabel(), "First button should be 'Add'")

	cancelButton := form.GetButton(1)
	assert.Equal(t, "Cancel", cancelButton.GetLabel(), "Second button should be 'Cancel'")
}

func TestAddFileView(t *testing.T) {
	app := tview.NewApplication()

	cfg := &config.ClientConfig{}
	client := NewClient(cfg)
	client.SetApp(app)

	element := client.AddFileView()

	form := element.(*tview.Form)
	assert.Equal(t, 2, form.GetFormItemCount(), "Form should have 2 input fields")

	filePathField := form.GetFormItemByLabel("File path").(*tview.InputField)
	assert.NotNil(t, filePathField, "File path input field should be in form")

	commentField := form.GetFormItemByLabel("Comment").(*tview.InputField)
	assert.NotNil(t, commentField, "Comment input field should be in form")

	uploadButton := form.GetButton(0)
	assert.Equal(t, "Upload", uploadButton.GetLabel(), "First button should be 'Upload'")

	cancelButton := form.GetButton(1)
	assert.Equal(t, "Cancel", cancelButton.GetLabel(), "Second button should be 'Cancel'")
}
