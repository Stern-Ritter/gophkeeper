package client

import (
	"github.com/rivo/tview"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
)

// Client represents a client in the TUI application, managing services and the application.
// It holds references to services for authentication, account management, card management, text management,
// file management, and the TUI application instance.
type Client struct {
	authService    AuthService
	accountService AccountService
	cardService    CardService
	textService    TextService
	fileService    FileService
	app            *tview.Application
	authToken      string
	config         *config.ClientConfig
}

func NewClient(config *config.ClientConfig) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) SetAuthService(authService AuthService) {
	c.authService = authService
}

func (c *Client) SetAccountService(accountService AccountService) {
	c.accountService = accountService
}

func (c *Client) SetCardService(cardService CardService) {
	c.cardService = cardService
}

func (c *Client) SetTextService(textService TextService) {
	c.textService = textService
}

func (c *Client) SetFileService(fileService FileService) {
	c.fileService = fileService
}

func (c *Client) SetApp(app *tview.Application) {
	c.app = app
}
