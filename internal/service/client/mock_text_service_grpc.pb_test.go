// Code generated by MockGen. DO NOT EDIT.
// Source: ./proto/gen/gophkeeper/gophkeeperapi/v1/text_service_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=./proto/gen/gophkeeper/gophkeeperapi/v1/text_service_grpc.pb.go -destination ./internal/service/client/mock_text_service_grpc.pb_test.go -package client
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

// MockTextServiceV1Client is a mock of TextServiceV1Client interface.
type MockTextServiceV1Client struct {
	ctrl     *gomock.Controller
	recorder *MockTextServiceV1ClientMockRecorder
}

// MockTextServiceV1ClientMockRecorder is the mock recorder for MockTextServiceV1Client.
type MockTextServiceV1ClientMockRecorder struct {
	mock *MockTextServiceV1Client
}

// NewMockTextServiceV1Client creates a new mock instance.
func NewMockTextServiceV1Client(ctrl *gomock.Controller) *MockTextServiceV1Client {
	mock := &MockTextServiceV1Client{ctrl: ctrl}
	mock.recorder = &MockTextServiceV1ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTextServiceV1Client) EXPECT() *MockTextServiceV1ClientMockRecorder {
	return m.recorder
}

// AddText mocks base method.
func (m *MockTextServiceV1Client) AddText(ctx context.Context, in *v1.AddTextRequestV1, opts ...grpc.CallOption) (*v1.AddTextResponseV1, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddText", varargs...)
	ret0, _ := ret[0].(*v1.AddTextResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddText indicates an expected call of AddText.
func (mr *MockTextServiceV1ClientMockRecorder) AddText(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddText", reflect.TypeOf((*MockTextServiceV1Client)(nil).AddText), varargs...)
}

// DeleteText mocks base method.
func (m *MockTextServiceV1Client) DeleteText(ctx context.Context, in *v1.DeleteTextRequestV1, opts ...grpc.CallOption) (*v1.DeleteTextResponseV1, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteText", varargs...)
	ret0, _ := ret[0].(*v1.DeleteTextResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteText indicates an expected call of DeleteText.
func (mr *MockTextServiceV1ClientMockRecorder) DeleteText(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteText", reflect.TypeOf((*MockTextServiceV1Client)(nil).DeleteText), varargs...)
}

// GetTexts mocks base method.
func (m *MockTextServiceV1Client) GetTexts(ctx context.Context, in *v1.GetTextsRequestV1, opts ...grpc.CallOption) (*v1.GetTextsResponseV1, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTexts", varargs...)
	ret0, _ := ret[0].(*v1.GetTextsResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTexts indicates an expected call of GetTexts.
func (mr *MockTextServiceV1ClientMockRecorder) GetTexts(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTexts", reflect.TypeOf((*MockTextServiceV1Client)(nil).GetTexts), varargs...)
}

// MockTextServiceV1Server is a mock of TextServiceV1Server interface.
type MockTextServiceV1Server struct {
	ctrl     *gomock.Controller
	recorder *MockTextServiceV1ServerMockRecorder
}

// MockTextServiceV1ServerMockRecorder is the mock recorder for MockTextServiceV1Server.
type MockTextServiceV1ServerMockRecorder struct {
	mock *MockTextServiceV1Server
}

// NewMockTextServiceV1Server creates a new mock instance.
func NewMockTextServiceV1Server(ctrl *gomock.Controller) *MockTextServiceV1Server {
	mock := &MockTextServiceV1Server{ctrl: ctrl}
	mock.recorder = &MockTextServiceV1ServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTextServiceV1Server) EXPECT() *MockTextServiceV1ServerMockRecorder {
	return m.recorder
}

// AddText mocks base method.
func (m *MockTextServiceV1Server) AddText(arg0 context.Context, arg1 *v1.AddTextRequestV1) (*v1.AddTextResponseV1, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddText", arg0, arg1)
	ret0, _ := ret[0].(*v1.AddTextResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddText indicates an expected call of AddText.
func (mr *MockTextServiceV1ServerMockRecorder) AddText(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddText", reflect.TypeOf((*MockTextServiceV1Server)(nil).AddText), arg0, arg1)
}

// DeleteText mocks base method.
func (m *MockTextServiceV1Server) DeleteText(arg0 context.Context, arg1 *v1.DeleteTextRequestV1) (*v1.DeleteTextResponseV1, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteText", arg0, arg1)
	ret0, _ := ret[0].(*v1.DeleteTextResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteText indicates an expected call of DeleteText.
func (mr *MockTextServiceV1ServerMockRecorder) DeleteText(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteText", reflect.TypeOf((*MockTextServiceV1Server)(nil).DeleteText), arg0, arg1)
}

// GetTexts mocks base method.
func (m *MockTextServiceV1Server) GetTexts(arg0 context.Context, arg1 *v1.GetTextsRequestV1) (*v1.GetTextsResponseV1, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTexts", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetTextsResponseV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTexts indicates an expected call of GetTexts.
func (mr *MockTextServiceV1ServerMockRecorder) GetTexts(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTexts", reflect.TypeOf((*MockTextServiceV1Server)(nil).GetTexts), arg0, arg1)
}

// mustEmbedUnimplementedTextServiceV1Server mocks base method.
func (m *MockTextServiceV1Server) mustEmbedUnimplementedTextServiceV1Server() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTextServiceV1Server")
}

// mustEmbedUnimplementedTextServiceV1Server indicates an expected call of mustEmbedUnimplementedTextServiceV1Server.
func (mr *MockTextServiceV1ServerMockRecorder) mustEmbedUnimplementedTextServiceV1Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTextServiceV1Server", reflect.TypeOf((*MockTextServiceV1Server)(nil).mustEmbedUnimplementedTextServiceV1Server))
}

// MockUnsafeTextServiceV1Server is a mock of UnsafeTextServiceV1Server interface.
type MockUnsafeTextServiceV1Server struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeTextServiceV1ServerMockRecorder
}

// MockUnsafeTextServiceV1ServerMockRecorder is the mock recorder for MockUnsafeTextServiceV1Server.
type MockUnsafeTextServiceV1ServerMockRecorder struct {
	mock *MockUnsafeTextServiceV1Server
}

// NewMockUnsafeTextServiceV1Server creates a new mock instance.
func NewMockUnsafeTextServiceV1Server(ctrl *gomock.Controller) *MockUnsafeTextServiceV1Server {
	mock := &MockUnsafeTextServiceV1Server{ctrl: ctrl}
	mock.recorder = &MockUnsafeTextServiceV1ServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeTextServiceV1Server) EXPECT() *MockUnsafeTextServiceV1ServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedTextServiceV1Server mocks base method.
func (m *MockUnsafeTextServiceV1Server) mustEmbedUnimplementedTextServiceV1Server() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTextServiceV1Server")
}

// mustEmbedUnimplementedTextServiceV1Server indicates an expected call of mustEmbedUnimplementedTextServiceV1Server.
func (mr *MockUnsafeTextServiceV1ServerMockRecorder) mustEmbedUnimplementedTextServiceV1Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTextServiceV1Server", reflect.TypeOf((*MockUnsafeTextServiceV1Server)(nil).mustEmbedUnimplementedTextServiceV1Server))
}
