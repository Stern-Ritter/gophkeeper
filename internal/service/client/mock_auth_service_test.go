// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/client/auth_service.go
//
// Generated by this command:
//
//	mockgen -source=./internal/service/client/auth_service.go -destination ./internal/service/client/mock_auth_service_test.go -package client
//

// Package client is a generated GoMock package.
package client

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// SignIn mocks base method.
func (m *MockAuthService) SignIn(login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockAuthServiceMockRecorder) SignIn(login, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockAuthService)(nil).SignIn), login, password)
}

// SignUp mocks base method.
func (m *MockAuthService) SignUp(login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockAuthServiceMockRecorder) SignUp(login, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthService)(nil).SignUp), login, password)
}
