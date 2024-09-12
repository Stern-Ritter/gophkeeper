package client

import (
	"context"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"google.golang.org/grpc"

	config "github.com/Stern-Ritter/gophkeeper/internal/config/client"
)

// Application represents an interface for interacting with a tview application.
// It provides methods for managing input, mouse events, screen, drawing, and focus.
type Application interface {
	SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *tview.Application
	GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey
	SetMouseCapture(capture func(event *tcell.EventMouse, action tview.MouseAction) (*tcell.EventMouse, tview.MouseAction)) *tview.Application
	GetMouseCapture() func(event *tcell.EventMouse, action tview.MouseAction) (*tcell.EventMouse, tview.MouseAction)
	SetScreen(screen tcell.Screen) *tview.Application
	EnableMouse(enable bool) *tview.Application
	EnablePaste(enable bool) *tview.Application
	Run() error
	Stop()
	Suspend(f func()) bool
	Draw() *tview.Application
	ForceDraw() *tview.Application
	Sync() *tview.Application
	SetBeforeDrawFunc(handler func(screen tcell.Screen) bool) *tview.Application
	GetBeforeDrawFunc() func(screen tcell.Screen) bool
	SetAfterDrawFunc(handler func(screen tcell.Screen)) *tview.Application
	GetAfterDrawFunc() func(screen tcell.Screen)
	SetRoot(root tview.Primitive, fullscreen bool) *tview.Application
	ResizeToFullScreen(p tview.Primitive) *tview.Application
	SetFocus(p tview.Primitive) *tview.Application
	GetFocus() tview.Primitive
	QueueUpdate(f func()) *tview.Application
	QueueUpdateDraw(f func()) *tview.Application
	QueueEvent(event tcell.Event) *tview.Application
}

// Client represents a client in the TUI application, managing services and the application.
// It holds references to services for authentication, account management, card management, text management,
// file management, and the TUI application instance.
type Client interface {
	SetAuthService(authService AuthService)
	SetAccountService(accountService AccountService)
	SetCardService(cardService CardService)
	SetTextService(textService TextService)
	SetFileService(fileService FileService)
	SetApp(app Application)
	GetAuthService() AuthService
	GetAccountService() AccountService
	GetCardService() CardService
	GetTextService() TextService
	GetFileService() FileService
	GetApp() Application
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
	app            Application
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

func (c *ClientImpl) SetApp(app Application) {
	c.app = app
}

func (c *ClientImpl) GetAuthService() AuthService {
	return c.authService
}

func (c *ClientImpl) GetAccountService() AccountService {
	return c.accountService
}

func (c *ClientImpl) GetCardService() CardService {
	return c.cardService
}

func (c *ClientImpl) GetTextService() TextService {
	return c.textService
}

func (c *ClientImpl) GetFileService() FileService {
	return c.fileService
}

func (c *ClientImpl) GetApp() Application {
	return c.app
}
