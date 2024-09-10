package client

import (
	"fmt"
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
)

func TestAuthView(t *testing.T) {
	app := tview.NewApplication()

	cfg := &config.ClientConfig{}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	client := NewClient(cfg)
	authService := NewMockAuthService(mockCtrl)
	client.SetAuthService(authService)
	client.SetApp(app)

	element := client.AuthView()

	form := element.(*tview.Form)
	assert.Equal(t, 2, form.GetFormItemCount(), "Form should have 2 input fields")

	loginField := form.GetFormItemByLabel("Login").(*tview.InputField)
	assert.NotNil(t, loginField, "Login input field should be in form")

	passwordField := form.GetFormItemByLabel("Password").(*tview.InputField)
	assert.NotNil(t, passwordField, "Password input field should be in form")

	signInButton := form.GetButton(0)
	assert.Equal(t, "Sign in", signInButton.GetLabel(), "First button should be 'Sign in'")

	signUpButton := form.GetButton(1)
	assert.Equal(t, "Sign up", signUpButton.GetLabel(), "Second button should be 'Sign up'")

	quitButton := form.GetButton(2)
	assert.Equal(t, "Quit", quitButton.GetLabel(), "Third button should be 'Quit'")
}

func TestSignUpHandler_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authService := NewMockAuthService(mockCtrl)
	authService.EXPECT().SignUp("user", "password").Return("auth token", nil)

	app := tview.NewApplication()
	cfg := &config.ClientConfig{}
	client := ClientImpl{
		authService: authService,
		config:      cfg,
	}
	client.SetApp(app)
	client.SetAuthService(authService)

	form := tview.NewForm()
	form.AddInputField("Login", "user", 20, nil, nil)
	form.AddInputField("Password", "password", 20, nil, nil)

	client.signUpHandler(form)

	assert.Equal(t, "auth token", client.authToken, "Auth token should be set")
}

func TestSignUpHandler_Failure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authService := NewMockAuthService(mockCtrl)
	authService.EXPECT().SignUp("user", "password").Return("", fmt.Errorf("registration failed"))

	app := tview.NewApplication()
	cfg := &config.ClientConfig{}
	client := ClientImpl{
		authService: authService,
		config:      cfg,
	}
	client.SetApp(app)
	client.SetAuthService(authService)

	form := tview.NewForm()
	form.AddInputField("Login", "user", 20, nil, nil)
	form.AddInputField("Password", "password", 20, nil, nil)

	client.signUpHandler(form)

	assert.Empty(t, client.authToken, "Auth token should be empty")
}

func TestSignInHandler_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authService := NewMockAuthService(mockCtrl)
	authService.EXPECT().SignIn("user", "password").Return("auth token", nil)

	app := tview.NewApplication()
	cfg := &config.ClientConfig{}
	client := &ClientImpl{
		authService: authService,
		config:      cfg,
	}
	client.SetApp(app)
	client.SetAuthService(authService)

	form := tview.NewForm()
	form.AddInputField("Login", "user", 20, nil, nil)
	form.AddInputField("Password", "password", 20, nil, nil)

	client.signInHandler(form)

	assert.Equal(t, "auth token", client.authToken, "Auth token should be set")
}

func TestSignInHandler_Failure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authService := NewMockAuthService(mockCtrl)
	authService.EXPECT().SignIn("user", "password").Return("", fmt.Errorf("authentication failed"))

	app := tview.NewApplication()
	cfg := &config.ClientConfig{}
	client := ClientImpl{
		authService: authService,
		config:      cfg,
	}
	client.SetApp(app)
	client.SetAuthService(authService)

	form := tview.NewForm()
	form.AddInputField("Login", "user", 20, nil, nil)
	form.AddInputField("Password", "password", 20, nil, nil)

	client.signInHandler(form)

	assert.Empty(t, client.authToken, "Auth token should be empty")
}
