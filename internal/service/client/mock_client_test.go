// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/client/client.go
//
// Generated by this command:
//
//	mockgen -source=./internal/service/client/client.go -destination ./internal/service/client/mock_client_test.go -package client
//

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	tcell "github.com/gdamore/tcell/v2"
	tview "github.com/rivo/tview"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// Draw mocks base method.
func (m *MockApplication) Draw() *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Draw")
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// Draw indicates an expected call of Draw.
func (mr *MockApplicationMockRecorder) Draw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Draw", reflect.TypeOf((*MockApplication)(nil).Draw))
}

// EnableMouse mocks base method.
func (m *MockApplication) EnableMouse(enable bool) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableMouse", enable)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// EnableMouse indicates an expected call of EnableMouse.
func (mr *MockApplicationMockRecorder) EnableMouse(enable any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableMouse", reflect.TypeOf((*MockApplication)(nil).EnableMouse), enable)
}

// EnablePaste mocks base method.
func (m *MockApplication) EnablePaste(enable bool) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePaste", enable)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// EnablePaste indicates an expected call of EnablePaste.
func (mr *MockApplicationMockRecorder) EnablePaste(enable any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePaste", reflect.TypeOf((*MockApplication)(nil).EnablePaste), enable)
}

// ForceDraw mocks base method.
func (m *MockApplication) ForceDraw() *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceDraw")
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// ForceDraw indicates an expected call of ForceDraw.
func (mr *MockApplicationMockRecorder) ForceDraw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceDraw", reflect.TypeOf((*MockApplication)(nil).ForceDraw))
}

// GetAfterDrawFunc mocks base method.
func (m *MockApplication) GetAfterDrawFunc() func(tcell.Screen) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAfterDrawFunc")
	ret0, _ := ret[0].(func(tcell.Screen))
	return ret0
}

// GetAfterDrawFunc indicates an expected call of GetAfterDrawFunc.
func (mr *MockApplicationMockRecorder) GetAfterDrawFunc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAfterDrawFunc", reflect.TypeOf((*MockApplication)(nil).GetAfterDrawFunc))
}

// GetBeforeDrawFunc mocks base method.
func (m *MockApplication) GetBeforeDrawFunc() func(tcell.Screen) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBeforeDrawFunc")
	ret0, _ := ret[0].(func(tcell.Screen) bool)
	return ret0
}

// GetBeforeDrawFunc indicates an expected call of GetBeforeDrawFunc.
func (mr *MockApplicationMockRecorder) GetBeforeDrawFunc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBeforeDrawFunc", reflect.TypeOf((*MockApplication)(nil).GetBeforeDrawFunc))
}

// GetFocus mocks base method.
func (m *MockApplication) GetFocus() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFocus")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// GetFocus indicates an expected call of GetFocus.
func (mr *MockApplicationMockRecorder) GetFocus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFocus", reflect.TypeOf((*MockApplication)(nil).GetFocus))
}

// GetInputCapture mocks base method.
func (m *MockApplication) GetInputCapture() func(*tcell.EventKey) *tcell.EventKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInputCapture")
	ret0, _ := ret[0].(func(*tcell.EventKey) *tcell.EventKey)
	return ret0
}

// GetInputCapture indicates an expected call of GetInputCapture.
func (mr *MockApplicationMockRecorder) GetInputCapture() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInputCapture", reflect.TypeOf((*MockApplication)(nil).GetInputCapture))
}

// GetMouseCapture mocks base method.
func (m *MockApplication) GetMouseCapture() func(*tcell.EventMouse, tview.MouseAction) (*tcell.EventMouse, tview.MouseAction) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMouseCapture")
	ret0, _ := ret[0].(func(*tcell.EventMouse, tview.MouseAction) (*tcell.EventMouse, tview.MouseAction))
	return ret0
}

// GetMouseCapture indicates an expected call of GetMouseCapture.
func (mr *MockApplicationMockRecorder) GetMouseCapture() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMouseCapture", reflect.TypeOf((*MockApplication)(nil).GetMouseCapture))
}

// QueueEvent mocks base method.
func (m *MockApplication) QueueEvent(event tcell.Event) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueEvent", event)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// QueueEvent indicates an expected call of QueueEvent.
func (mr *MockApplicationMockRecorder) QueueEvent(event any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueEvent", reflect.TypeOf((*MockApplication)(nil).QueueEvent), event)
}

// QueueUpdate mocks base method.
func (m *MockApplication) QueueUpdate(f func()) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueUpdate", f)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// QueueUpdate indicates an expected call of QueueUpdate.
func (mr *MockApplicationMockRecorder) QueueUpdate(f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueUpdate", reflect.TypeOf((*MockApplication)(nil).QueueUpdate), f)
}

// QueueUpdateDraw mocks base method.
func (m *MockApplication) QueueUpdateDraw(f func()) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueUpdateDraw", f)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// QueueUpdateDraw indicates an expected call of QueueUpdateDraw.
func (mr *MockApplicationMockRecorder) QueueUpdateDraw(f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueUpdateDraw", reflect.TypeOf((*MockApplication)(nil).QueueUpdateDraw), f)
}

// ResizeToFullScreen mocks base method.
func (m *MockApplication) ResizeToFullScreen(p tview.Primitive) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeToFullScreen", p)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// ResizeToFullScreen indicates an expected call of ResizeToFullScreen.
func (mr *MockApplicationMockRecorder) ResizeToFullScreen(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizeToFullScreen", reflect.TypeOf((*MockApplication)(nil).ResizeToFullScreen), p)
}

// Run mocks base method.
func (m *MockApplication) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockApplicationMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockApplication)(nil).Run))
}

// SetAfterDrawFunc mocks base method.
func (m *MockApplication) SetAfterDrawFunc(handler func(tcell.Screen)) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAfterDrawFunc", handler)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// SetAfterDrawFunc indicates an expected call of SetAfterDrawFunc.
func (mr *MockApplicationMockRecorder) SetAfterDrawFunc(handler any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAfterDrawFunc", reflect.TypeOf((*MockApplication)(nil).SetAfterDrawFunc), handler)
}

// SetBeforeDrawFunc mocks base method.
func (m *MockApplication) SetBeforeDrawFunc(handler func(tcell.Screen) bool) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBeforeDrawFunc", handler)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// SetBeforeDrawFunc indicates an expected call of SetBeforeDrawFunc.
func (mr *MockApplicationMockRecorder) SetBeforeDrawFunc(handler any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBeforeDrawFunc", reflect.TypeOf((*MockApplication)(nil).SetBeforeDrawFunc), handler)
}

// SetFocus mocks base method.
func (m *MockApplication) SetFocus(p tview.Primitive) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFocus", p)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// SetFocus indicates an expected call of SetFocus.
func (mr *MockApplicationMockRecorder) SetFocus(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFocus", reflect.TypeOf((*MockApplication)(nil).SetFocus), p)
}

// SetInputCapture mocks base method.
func (m *MockApplication) SetInputCapture(capture func(*tcell.EventKey) *tcell.EventKey) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInputCapture", capture)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// SetInputCapture indicates an expected call of SetInputCapture.
func (mr *MockApplicationMockRecorder) SetInputCapture(capture any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInputCapture", reflect.TypeOf((*MockApplication)(nil).SetInputCapture), capture)
}

// SetMouseCapture mocks base method.
func (m *MockApplication) SetMouseCapture(capture func(*tcell.EventMouse, tview.MouseAction) (*tcell.EventMouse, tview.MouseAction)) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMouseCapture", capture)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// SetMouseCapture indicates an expected call of SetMouseCapture.
func (mr *MockApplicationMockRecorder) SetMouseCapture(capture any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMouseCapture", reflect.TypeOf((*MockApplication)(nil).SetMouseCapture), capture)
}

// SetRoot mocks base method.
func (m *MockApplication) SetRoot(root tview.Primitive, fullscreen bool) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRoot", root, fullscreen)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// SetRoot indicates an expected call of SetRoot.
func (mr *MockApplicationMockRecorder) SetRoot(root, fullscreen any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRoot", reflect.TypeOf((*MockApplication)(nil).SetRoot), root, fullscreen)
}

// SetScreen mocks base method.
func (m *MockApplication) SetScreen(screen tcell.Screen) *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetScreen", screen)
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// SetScreen indicates an expected call of SetScreen.
func (mr *MockApplicationMockRecorder) SetScreen(screen any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetScreen", reflect.TypeOf((*MockApplication)(nil).SetScreen), screen)
}

// Stop mocks base method.
func (m *MockApplication) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockApplicationMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockApplication)(nil).Stop))
}

// Suspend mocks base method.
func (m *MockApplication) Suspend(f func()) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suspend", f)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Suspend indicates an expected call of Suspend.
func (mr *MockApplicationMockRecorder) Suspend(f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suspend", reflect.TypeOf((*MockApplication)(nil).Suspend), f)
}

// Sync mocks base method.
func (m *MockApplication) Sync() *tview.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync")
	ret0, _ := ret[0].(*tview.Application)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockApplicationMockRecorder) Sync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockApplication)(nil).Sync))
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AccountsView mocks base method.
func (m *MockClient) AccountsView(previousView tview.Primitive) tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountsView", previousView)
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// AccountsView indicates an expected call of AccountsView.
func (mr *MockClientMockRecorder) AccountsView(previousView any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountsView", reflect.TypeOf((*MockClient)(nil).AccountsView), previousView)
}

// AddAccountView mocks base method.
func (m *MockClient) AddAccountView() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountView")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// AddAccountView indicates an expected call of AddAccountView.
func (mr *MockClientMockRecorder) AddAccountView() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountView", reflect.TypeOf((*MockClient)(nil).AddAccountView))
}

// AddCardView mocks base method.
func (m *MockClient) AddCardView() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCardView")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// AddCardView indicates an expected call of AddCardView.
func (mr *MockClientMockRecorder) AddCardView() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCardView", reflect.TypeOf((*MockClient)(nil).AddCardView))
}

// AddFileView mocks base method.
func (m *MockClient) AddFileView() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFileView")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// AddFileView indicates an expected call of AddFileView.
func (mr *MockClientMockRecorder) AddFileView() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFileView", reflect.TypeOf((*MockClient)(nil).AddFileView))
}

// AddTextView mocks base method.
func (m *MockClient) AddTextView() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTextView")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// AddTextView indicates an expected call of AddTextView.
func (mr *MockClientMockRecorder) AddTextView() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTextView", reflect.TypeOf((*MockClient)(nil).AddTextView))
}

// AddView mocks base method.
func (m *MockClient) AddView() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddView")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// AddView indicates an expected call of AddView.
func (mr *MockClientMockRecorder) AddView() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddView", reflect.TypeOf((*MockClient)(nil).AddView))
}

// AuthInterceptor mocks base method.
func (m *MockClient) AuthInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, method, req, reply, cc, invoker}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthInterceptor", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AuthInterceptor indicates an expected call of AuthInterceptor.
func (mr *MockClientMockRecorder) AuthInterceptor(ctx, method, req, reply, cc, invoker any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, method, req, reply, cc, invoker}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthInterceptor", reflect.TypeOf((*MockClient)(nil).AuthInterceptor), varargs...)
}

// AuthStreamInterceptor mocks base method.
func (m *MockClient) AuthStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, desc, cc, method, streamer}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthStreamInterceptor", varargs...)
	ret0, _ := ret[0].(grpc.ClientStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthStreamInterceptor indicates an expected call of AuthStreamInterceptor.
func (mr *MockClientMockRecorder) AuthStreamInterceptor(ctx, desc, cc, method, streamer any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, desc, cc, method, streamer}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthStreamInterceptor", reflect.TypeOf((*MockClient)(nil).AuthStreamInterceptor), varargs...)
}

// AuthView mocks base method.
func (m *MockClient) AuthView() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthView")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// AuthView indicates an expected call of AuthView.
func (mr *MockClientMockRecorder) AuthView() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthView", reflect.TypeOf((*MockClient)(nil).AuthView))
}

// CardsView mocks base method.
func (m *MockClient) CardsView(previousView tview.Primitive) tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CardsView", previousView)
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// CardsView indicates an expected call of CardsView.
func (mr *MockClientMockRecorder) CardsView(previousView any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CardsView", reflect.TypeOf((*MockClient)(nil).CardsView), previousView)
}

// DataView mocks base method.
func (m *MockClient) DataView() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataView")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// DataView indicates an expected call of DataView.
func (mr *MockClientMockRecorder) DataView() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataView", reflect.TypeOf((*MockClient)(nil).DataView))
}

// FilesView mocks base method.
func (m *MockClient) FilesView(previousView tview.Primitive) tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilesView", previousView)
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// FilesView indicates an expected call of FilesView.
func (mr *MockClientMockRecorder) FilesView(previousView any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesView", reflect.TypeOf((*MockClient)(nil).FilesView), previousView)
}

// GetAccountService mocks base method.
func (m *MockClient) GetAccountService() AccountService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountService")
	ret0, _ := ret[0].(AccountService)
	return ret0
}

// GetAccountService indicates an expected call of GetAccountService.
func (mr *MockClientMockRecorder) GetAccountService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountService", reflect.TypeOf((*MockClient)(nil).GetAccountService))
}

// GetApp mocks base method.
func (m *MockClient) GetApp() Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApp")
	ret0, _ := ret[0].(Application)
	return ret0
}

// GetApp indicates an expected call of GetApp.
func (mr *MockClientMockRecorder) GetApp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApp", reflect.TypeOf((*MockClient)(nil).GetApp))
}

// GetAuthService mocks base method.
func (m *MockClient) GetAuthService() AuthService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthService")
	ret0, _ := ret[0].(AuthService)
	return ret0
}

// GetAuthService indicates an expected call of GetAuthService.
func (mr *MockClientMockRecorder) GetAuthService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthService", reflect.TypeOf((*MockClient)(nil).GetAuthService))
}

// GetCardService mocks base method.
func (m *MockClient) GetCardService() CardService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCardService")
	ret0, _ := ret[0].(CardService)
	return ret0
}

// GetCardService indicates an expected call of GetCardService.
func (mr *MockClientMockRecorder) GetCardService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCardService", reflect.TypeOf((*MockClient)(nil).GetCardService))
}

// GetFileService mocks base method.
func (m *MockClient) GetFileService() FileService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileService")
	ret0, _ := ret[0].(FileService)
	return ret0
}

// GetFileService indicates an expected call of GetFileService.
func (mr *MockClientMockRecorder) GetFileService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileService", reflect.TypeOf((*MockClient)(nil).GetFileService))
}

// GetTextService mocks base method.
func (m *MockClient) GetTextService() TextService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTextService")
	ret0, _ := ret[0].(TextService)
	return ret0
}

// GetTextService indicates an expected call of GetTextService.
func (mr *MockClientMockRecorder) GetTextService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTextService", reflect.TypeOf((*MockClient)(nil).GetTextService))
}

// MainView mocks base method.
func (m *MockClient) MainView() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MainView")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// MainView indicates an expected call of MainView.
func (mr *MockClientMockRecorder) MainView() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MainView", reflect.TypeOf((*MockClient)(nil).MainView))
}

// QueueUpdateDraw mocks base method.
func (m *MockClient) QueueUpdateDraw(render func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "QueueUpdateDraw", render)
}

// QueueUpdateDraw indicates an expected call of QueueUpdateDraw.
func (mr *MockClientMockRecorder) QueueUpdateDraw(render any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueUpdateDraw", reflect.TypeOf((*MockClient)(nil).QueueUpdateDraw), render)
}

// SelectView mocks base method.
func (m *MockClient) SelectView(view tview.Primitive) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SelectView", view)
}

// SelectView indicates an expected call of SelectView.
func (mr *MockClientMockRecorder) SelectView(view any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectView", reflect.TypeOf((*MockClient)(nil).SelectView), view)
}

// SetAccountService mocks base method.
func (m *MockClient) SetAccountService(accountService AccountService) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccountService", accountService)
}

// SetAccountService indicates an expected call of SetAccountService.
func (mr *MockClientMockRecorder) SetAccountService(accountService any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccountService", reflect.TypeOf((*MockClient)(nil).SetAccountService), accountService)
}

// SetApp mocks base method.
func (m *MockClient) SetApp(app Application) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetApp", app)
}

// SetApp indicates an expected call of SetApp.
func (mr *MockClientMockRecorder) SetApp(app any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetApp", reflect.TypeOf((*MockClient)(nil).SetApp), app)
}

// SetAuthService mocks base method.
func (m *MockClient) SetAuthService(authService AuthService) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAuthService", authService)
}

// SetAuthService indicates an expected call of SetAuthService.
func (mr *MockClientMockRecorder) SetAuthService(authService any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAuthService", reflect.TypeOf((*MockClient)(nil).SetAuthService), authService)
}

// SetCardService mocks base method.
func (m *MockClient) SetCardService(cardService CardService) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCardService", cardService)
}

// SetCardService indicates an expected call of SetCardService.
func (mr *MockClientMockRecorder) SetCardService(cardService any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCardService", reflect.TypeOf((*MockClient)(nil).SetCardService), cardService)
}

// SetFileService mocks base method.
func (m *MockClient) SetFileService(fileService FileService) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFileService", fileService)
}

// SetFileService indicates an expected call of SetFileService.
func (mr *MockClientMockRecorder) SetFileService(fileService any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFileService", reflect.TypeOf((*MockClient)(nil).SetFileService), fileService)
}

// SetTextService mocks base method.
func (m *MockClient) SetTextService(textService TextService) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTextService", textService)
}

// SetTextService indicates an expected call of SetTextService.
func (mr *MockClientMockRecorder) SetTextService(textService any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTextService", reflect.TypeOf((*MockClient)(nil).SetTextService), textService)
}

// ShowConfirmModal mocks base method.
func (m *MockClient) ShowConfirmModal(text string, currentView tview.Primitive, handler func()) tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowConfirmModal", text, currentView, handler)
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// ShowConfirmModal indicates an expected call of ShowConfirmModal.
func (mr *MockClientMockRecorder) ShowConfirmModal(text, currentView, handler any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowConfirmModal", reflect.TypeOf((*MockClient)(nil).ShowConfirmModal), text, currentView, handler)
}

// ShowInfoModal mocks base method.
func (m *MockClient) ShowInfoModal(text string, currentView tview.Primitive) tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowInfoModal", text, currentView)
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// ShowInfoModal indicates an expected call of ShowInfoModal.
func (mr *MockClientMockRecorder) ShowInfoModal(text, currentView any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowInfoModal", reflect.TypeOf((*MockClient)(nil).ShowInfoModal), text, currentView)
}

// ShowRetryModal mocks base method.
func (m *MockClient) ShowRetryModal(text string, currentView, previousView tview.Primitive) tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowRetryModal", text, currentView, previousView)
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// ShowRetryModal indicates an expected call of ShowRetryModal.
func (mr *MockClientMockRecorder) ShowRetryModal(text, currentView, previousView any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowRetryModal", reflect.TypeOf((*MockClient)(nil).ShowRetryModal), text, currentView, previousView)
}

// TextsView mocks base method.
func (m *MockClient) TextsView(previousView tview.Primitive) tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TextsView", previousView)
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// TextsView indicates an expected call of TextsView.
func (mr *MockClientMockRecorder) TextsView(previousView any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TextsView", reflect.TypeOf((*MockClient)(nil).TextsView), previousView)
}

// UpdateDraw mocks base method.
func (m *MockClient) UpdateDraw() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDraw")
}

// UpdateDraw indicates an expected call of UpdateDraw.
func (mr *MockClientMockRecorder) UpdateDraw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDraw", reflect.TypeOf((*MockClient)(nil).UpdateDraw))
}

// VersionView mocks base method.
func (m *MockClient) VersionView() tview.Primitive {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VersionView")
	ret0, _ := ret[0].(tview.Primitive)
	return ret0
}

// VersionView indicates an expected call of VersionView.
func (mr *MockClientMockRecorder) VersionView() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VersionView", reflect.TypeOf((*MockClient)(nil).VersionView))
}
