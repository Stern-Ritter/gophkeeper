package client

import (
	"context"

	"github.com/rivo/tview"
	"google.golang.org/grpc"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
)

// Client represents a client in the TUI application, managing services and the application.
// It holds references to services for authentication, account management, card management, text management,
// file management, and the TUI application instance.
type Client interface {
	SetAuthService(authService AuthService)
	SetAccountService(accountService AccountService)
	SetCardService(cardService CardService)
	SetTextService(textService TextService)
	SetFileService(fileService FileService)
	SetApp(app *tview.Application)
	AuthInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error
	AuthStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string,
		streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error)
	SelectView(view tview.Primitive)
	UpdateDraw()
	QueueUpdateDraw(render func())
	AuthView() tview.Primitive
	MainView() tview.Primitive
	AddView() tview.Primitive
	AddAccountView() tview.Primitive
	AddCardView() tview.Primitive
	AddTextView() tview.Primitive
	AddFileView() tview.Primitive
	DataView() tview.Primitive
	AccountsView(previousView tview.Primitive) tview.Primitive
	CardsView(previousView tview.Primitive) tview.Primitive
	TextsView(previousView tview.Primitive) tview.Primitive
	FilesView(previousView tview.Primitive) tview.Primitive
	VersionView() tview.Primitive
	ShowInfoModal(text string, currentView tview.Primitive) tview.Primitive
	ShowConfirmModal(text string, currentView tview.Primitive, handler func()) tview.Primitive
	ShowRetryModal(text string, currentView tview.Primitive, previousView tview.Primitive) tview.Primitive
}

type ClientImpl struct {
	authService    AuthService
	accountService AccountService
	cardService    CardService
	textService    TextService
	fileService    FileService
	app            *tview.Application
	authToken      string
	config         *config.ClientConfig
}

func NewClient(config *config.ClientConfig) Client {
	return &ClientImpl{
		config: config,
	}
}

func (c *ClientImpl) SetAuthService(authService AuthService) {
	c.authService = authService
}

func (c *ClientImpl) SetAccountService(accountService AccountService) {
	c.accountService = accountService
}

func (c *ClientImpl) SetCardService(cardService CardService) {
	c.cardService = cardService
}

func (c *ClientImpl) SetTextService(textService TextService) {
	c.textService = textService
}

func (c *ClientImpl) SetFileService(fileService FileService) {
	c.fileService = fileService
}

func (c *ClientImpl) SetApp(app *tview.Application) {
	c.app = app
}
