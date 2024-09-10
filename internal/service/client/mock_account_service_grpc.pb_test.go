// Code generated by MockGen. DO NOT EDIT.
// Source: ./proto/gen/gophkeeper/gophkeeperapi/v1/account_service_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=./proto/gen/gophkeeper/gophkeeperapi/v1/account_service_grpc.pb.go -destination ./internal/service/client/mock_account_service_grpc.pb_test.go -package client
//

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	v1 "github.com/Stern-Ritter/gophkeeper/proto/gen/gophkeeper/gophkeeperapi/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockAccountServiceV1Client is a mock of AccountServiceV1Client interface.
type MockAccountServiceV1Client struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceV1ClientMockRecorder
}

// MockAccountServiceV1ClientMockRecorder is the mock recorder for MockAccountServiceV1Client.
type MockAccountServiceV1ClientMockRecorder struct {
	mock *MockAccountServiceV1Client
}

// NewMockAccountServiceV1Client creates a new mock instance.
func NewMockAccountServiceV1Client(ctrl *gomock.Controller) *MockAccountServiceV1Client {
	mock := &MockAccountServiceV1Client{ctrl: ctrl}
	mock.recorder = &MockAccountServiceV1ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountServiceV1Client) EXPECT() *MockAccountServiceV1ClientMockRecorder {
	return m.recorder
}

// AddAccount mocks base method.
func (m *MockAccountServiceV1Client) AddAccount(ctx context.Context, in *v1.AddAccountRequestV1, opts ...grpc.CallOption) (*v1.AddAccountResponseV1, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddAccount", varargs...)
	ret0, _ := ret[0].(*v1.AddAccountResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccount indicates an expected call of AddAccount.
func (mr *MockAccountServiceV1ClientMockRecorder) AddAccount(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccount", reflect.TypeOf((*MockAccountServiceV1Client)(nil).AddAccount), varargs...)
}

// DeleteAccount mocks base method.
func (m *MockAccountServiceV1Client) DeleteAccount(ctx context.Context, in *v1.DeleteAccountRequestV1, opts ...grpc.CallOption) (*v1.DeleteAccountResponseV1, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAccount", varargs...)
	ret0, _ := ret[0].(*v1.DeleteAccountResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockAccountServiceV1ClientMockRecorder) DeleteAccount(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockAccountServiceV1Client)(nil).DeleteAccount), varargs...)
}

// GetAccounts mocks base method.
func (m *MockAccountServiceV1Client) GetAccounts(ctx context.Context, in *v1.GetAccountsRequestV1, opts ...grpc.CallOption) (*v1.GetAccountsResponseV1, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccounts", varargs...)
	ret0, _ := ret[0].(*v1.GetAccountsResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockAccountServiceV1ClientMockRecorder) GetAccounts(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockAccountServiceV1Client)(nil).GetAccounts), varargs...)
}

// MockAccountServiceV1Server is a mock of AccountServiceV1Server interface.
type MockAccountServiceV1Server struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceV1ServerMockRecorder
}

// MockAccountServiceV1ServerMockRecorder is the mock recorder for MockAccountServiceV1Server.
type MockAccountServiceV1ServerMockRecorder struct {
	mock *MockAccountServiceV1Server
}

// NewMockAccountServiceV1Server creates a new mock instance.
func NewMockAccountServiceV1Server(ctrl *gomock.Controller) *MockAccountServiceV1Server {
	mock := &MockAccountServiceV1Server{ctrl: ctrl}
	mock.recorder = &MockAccountServiceV1ServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountServiceV1Server) EXPECT() *MockAccountServiceV1ServerMockRecorder {
	return m.recorder
}

// AddAccount mocks base method.
func (m *MockAccountServiceV1Server) AddAccount(arg0 context.Context, arg1 *v1.AddAccountRequestV1) (*v1.AddAccountResponseV1, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccount", arg0, arg1)
	ret0, _ := ret[0].(*v1.AddAccountResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccount indicates an expected call of AddAccount.
func (mr *MockAccountServiceV1ServerMockRecorder) AddAccount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccount", reflect.TypeOf((*MockAccountServiceV1Server)(nil).AddAccount), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockAccountServiceV1Server) DeleteAccount(arg0 context.Context, arg1 *v1.DeleteAccountRequestV1) (*v1.DeleteAccountResponseV1, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(*v1.DeleteAccountResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockAccountServiceV1ServerMockRecorder) DeleteAccount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockAccountServiceV1Server)(nil).DeleteAccount), arg0, arg1)
}

// GetAccounts mocks base method.
func (m *MockAccountServiceV1Server) GetAccounts(arg0 context.Context, arg1 *v1.GetAccountsRequestV1) (*v1.GetAccountsResponseV1, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetAccountsResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockAccountServiceV1ServerMockRecorder) GetAccounts(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockAccountServiceV1Server)(nil).GetAccounts), arg0, arg1)
}

// mustEmbedUnimplementedAccountServiceV1Server mocks base method.
func (m *MockAccountServiceV1Server) mustEmbedUnimplementedAccountServiceV1Server() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAccountServiceV1Server")
}

// mustEmbedUnimplementedAccountServiceV1Server indicates an expected call of mustEmbedUnimplementedAccountServiceV1Server.
func (mr *MockAccountServiceV1ServerMockRecorder) mustEmbedUnimplementedAccountServiceV1Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAccountServiceV1Server", reflect.TypeOf((*MockAccountServiceV1Server)(nil).mustEmbedUnimplementedAccountServiceV1Server))
}

// MockUnsafeAccountServiceV1Server is a mock of UnsafeAccountServiceV1Server interface.
type MockUnsafeAccountServiceV1Server struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAccountServiceV1ServerMockRecorder
}

// MockUnsafeAccountServiceV1ServerMockRecorder is the mock recorder for MockUnsafeAccountServiceV1Server.
type MockUnsafeAccountServiceV1ServerMockRecorder struct {
	mock *MockUnsafeAccountServiceV1Server
}

// NewMockUnsafeAccountServiceV1Server creates a new mock instance.
func NewMockUnsafeAccountServiceV1Server(ctrl *gomock.Controller) *MockUnsafeAccountServiceV1Server {
	mock := &MockUnsafeAccountServiceV1Server{ctrl: ctrl}
	mock.recorder = &MockUnsafeAccountServiceV1ServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAccountServiceV1Server) EXPECT() *MockUnsafeAccountServiceV1ServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAccountServiceV1Server mocks base method.
func (m *MockUnsafeAccountServiceV1Server) mustEmbedUnimplementedAccountServiceV1Server() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAccountServiceV1Server")
}

// mustEmbedUnimplementedAccountServiceV1Server indicates an expected call of mustEmbedUnimplementedAccountServiceV1Server.
func (mr *MockUnsafeAccountServiceV1ServerMockRecorder) mustEmbedUnimplementedAccountServiceV1Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAccountServiceV1Server", reflect.TypeOf((*MockUnsafeAccountServiceV1Server)(nil).mustEmbedUnimplementedAccountServiceV1Server))
}