package client

import (
	"errors"
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

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

func TestAddAccountHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockClient(ctrl)
	mockAccountService := NewMockAccountService(ctrl)

	form := tview.NewForm()
	form.AddInputField("Login", "", 20, nil, nil)
	form.AddInputField("Password", "", 20, nil, nil)
	form.AddInputField("Comment", "", 20, nil, nil)

	form.GetFormItemByLabel("Login").(*tview.InputField).SetText("user")
	form.GetFormItemByLabel("Password").(*tview.InputField).SetText("password")
	form.GetFormItemByLabel("Comment").(*tview.InputField).SetText("comment")

	mockAccountService.EXPECT().CreateAccount("user", "password", "comment").Return(nil)
	mockClient.EXPECT().ShowInfoModal("Success added account data", gomock.Any()).Times(1)
	mockClient.EXPECT().AddAccountView().Times(1)

	addAccountHandler(mockClient, mockAccountService, form)
}

func TestAddAccountHandler_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockClient(ctrl)
	mockAccountService := NewMockAccountService(ctrl)

	form := tview.NewForm()
	form.AddInputField("Login", "", 20, nil, nil)
	form.AddInputField("Password", "", 20, nil, nil)
	form.AddInputField("Comment", "", 20, nil, nil)

	form.GetFormItemByLabel("Login").(*tview.InputField).SetText("user")
	form.GetFormItemByLabel("Password").(*tview.InputField).SetText("password")
	form.GetFormItemByLabel("Comment").(*tview.InputField).SetText("comment")

	err := errors.New("failed to create account")

	mockAccountService.EXPECT().CreateAccount("user", "password", "comment").Return(err)
	mockClient.EXPECT().ShowInfoModal("Error adding account data: failed to create account", gomock.Any()).Times(1)

	addAccountHandler(mockClient, mockAccountService, form)
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

func TestAddCardHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockClient(ctrl)
	mockCardService := NewMockCardService(ctrl)

	form := tview.NewForm()
	form.AddInputField("Number", "", 20, nil, nil)
	form.AddInputField("Owner", "", 20, nil, nil)
	form.AddInputField("Expiry", "", 20, nil, nil)
	form.AddInputField("CVC", "", 20, nil, nil)
	form.AddInputField("PIN", "", 20, nil, nil)
	form.AddInputField("Comment", "", 20, nil, nil)

	form.GetFormItemByLabel("Number").(*tview.InputField).SetText("1234 1234 1234 1234")
	form.GetFormItemByLabel("Owner").(*tview.InputField).SetText("John Doe")
	form.GetFormItemByLabel("Expiry").(*tview.InputField).SetText("12/25")
	form.GetFormItemByLabel("CVC").(*tview.InputField).SetText("123")
	form.GetFormItemByLabel("PIN").(*tview.InputField).SetText("0000")
	form.GetFormItemByLabel("Comment").(*tview.InputField).SetText("Debit bank card")

	mockCardService.EXPECT().CreateCard("1234 1234 1234 1234", "John Doe", "12/25", "123",
		"0000", "Debit bank card").Return(nil)
	mockClient.EXPECT().ShowInfoModal("Success added card data", gomock.Any()).Times(1)
	mockClient.EXPECT().AddCardView().Times(1)

	addCardHandler(mockClient, mockCardService, form)
}

func TestAddCardHandler_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockClient(ctrl)
	mockCardService := NewMockCardService(ctrl)

	form := tview.NewForm()
	form.AddInputField("Number", "", 20, nil, nil)
	form.AddInputField("Owner", "", 20, nil, nil)
	form.AddInputField("Expiry", "", 20, nil, nil)
	form.AddInputField("CVC", "", 20, nil, nil)
	form.AddInputField("PIN", "", 20, nil, nil)
	form.AddInputField("Comment", "", 20, nil, nil)

	form.GetFormItemByLabel("Number").(*tview.InputField).SetText("1234 1234 1234 1234")
	form.GetFormItemByLabel("Owner").(*tview.InputField).SetText("John Doe")
	form.GetFormItemByLabel("Expiry").(*tview.InputField).SetText("12/25")
	form.GetFormItemByLabel("CVC").(*tview.InputField).SetText("123")
	form.GetFormItemByLabel("PIN").(*tview.InputField).SetText("0000")
	form.GetFormItemByLabel("Comment").(*tview.InputField).SetText("Debit bank card")

	err := errors.New("failed to create card")

	mockCardService.EXPECT().CreateCard("1234 1234 1234 1234", "John Doe", "12/25", "123",
		"0000", "Debit bank card").Return(err)
	mockClient.EXPECT().ShowInfoModal("Error adding card data: failed to create card", gomock.Any()).Times(1)

	addCardHandler(mockClient, mockCardService, form)
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

func TestAddTextHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockClient(ctrl)
	mockTextService := NewMockTextService(ctrl)

	form := tview.NewForm()
	form.AddInputField("Text", "", 20, nil, nil)
	form.AddInputField("Comment", "", 20, nil, nil)

	form.GetFormItemByLabel("Text").(*tview.InputField).SetText("text")
	form.GetFormItemByLabel("Comment").(*tview.InputField).SetText("comment")

	mockTextService.EXPECT().CreateText("text", "comment").Return(nil)
	mockClient.EXPECT().ShowInfoModal("Success added text data", gomock.Any()).Times(1)
	mockClient.EXPECT().AddTextView().Times(1)

	addTextHandler(mockClient, mockTextService, form)
}

func TestAddTextHandler_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockClient(ctrl)
	mockTextService := NewMockTextService(ctrl)

	form := tview.NewForm()
	form.AddInputField("Text", "", 20, nil, nil)
	form.AddInputField("Comment", "", 20, nil, nil)

	form.GetFormItemByLabel("Text").(*tview.InputField).SetText("text")
	form.GetFormItemByLabel("Comment").(*tview.InputField).SetText("comment")

	err := errors.New("failed to create text")

	mockTextService.EXPECT().CreateText("text", "comment").Return(err)
	mockClient.EXPECT().ShowInfoModal("Error adding text data: failed to create text", gomock.Any()).Times(1)

	addTextHandler(mockClient, mockTextService, form)
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

func TestUploadFileHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockClient(ctrl)
	mockFileService := NewMockFileService(ctrl)

	form := tview.NewForm()
	form.AddInputField("File path", "", 20, nil, nil)
	form.AddInputField("Comment", "", 20, nil, nil)

	form.GetFormItemByLabel("File path").(*tview.InputField).SetText("path")
	form.GetFormItemByLabel("Comment").(*tview.InputField).SetText("comment")

	mockFileService.EXPECT().UploadFile(gomock.Any(), "path", "comment", gomock.Any()).Return(nil)

	mockClient.EXPECT().SelectView(gomock.Any()).Times(1)
	mockClient.EXPECT().QueueUpdateDraw(gomock.Any()).AnyTimes()

	uploadFileHandler(mockClient, mockFileService, form)
}
