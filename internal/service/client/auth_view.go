package client

import (
	"fmt"

	"github.com/rivo/tview"
	"google.golang.org/grpc/status"
)

// AuthView displays an authentication form for the user to sign in or sign up.
func (c *Client) AuthView() {
	form := tview.NewForm()
	form.AddInputField("Login", "", 20, nil, nil).
		AddInputField("Password", "", 20, nil, nil).
		AddButton("Sign in", func() { c.signInHandler(form) }).
		AddButton("Sign up", func() { c.signUpHandler(form) }).
		AddButton("Quit", func() { stopApp(c.app) })

	selectView(c.app, form)
}

// signUpHandler handles the sign-up process.
// It retrieves the login and password from the form, sends a sign-up request to the authentication service,
// and handles the response.
// If sign-up is successful, the user's authentication token is stored and the main view is displayed.
// If sign-up fails, a retry modal is shown with options to try again or cancel.
func (c *Client) signUpHandler(form *tview.Form) {
	login := form.GetFormItemByLabel("Login").(*tview.InputField).GetText()
	password := form.GetFormItemByLabel("Password").(*tview.InputField).GetText()

	token, err := c.authService.SignUp(login, password)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowRetryModal(fmt.Sprintf("Registration failed: %s. Try again or cancel.", err.Error()), form, form)
		} else {
			errMsg := st.Message()
			c.ShowRetryModal(fmt.Sprintf("Registration failed: %s. Try again or cancel.", errMsg), form, form)
		}
		return
	}

	c.authToken = token
	c.MainView()
}

// signInHandler handles the sign-in process.
// It retrieves the login and password from the form, sends a sign-in request to the authentication service,
// and handles the response.
// If sign-in is successful, the user's authentication token is stored and the main view is displayed.
// If sign-in fails, a retry modal is shown with options to try again or cancel.
func (c *Client) signInHandler(form *tview.Form) {
	login := form.GetFormItemByLabel("Login").(*tview.InputField).GetText()
	password := form.GetFormItemByLabel("Password").(*tview.InputField).GetText()

	token, err := c.authService.SignIn(login, password)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			c.ShowRetryModal(fmt.Sprintf("Login failed: %s. Try again or cancel.", err.Error()), form, form)
		} else {
			errMsg := st.Message()
			c.ShowRetryModal(fmt.Sprintf("Login failed: %s. Try again or cancel.", errMsg), form, form)
		}
		return
	}

	c.authToken = token
	c.MainView()
}
