// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/storage/server/user_storage.go
//
// Generated by this command:
//
//	mockgen -source=./internal/storage/server/user_storage.go -destination ./internal/service/server/mock_user_storage_test.go -package server
//

// Package server is a generated GoMock package.
package server

import (
	context "context"
	reflect "reflect"

	model "github.com/Stern-Ritter/gophkeeper/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockUserStorage is a mock of UserStorage interface.
type MockUserStorage struct {
	ctrl     *gomock.Controller
	recorder *MockUserStorageMockRecorder
}

// MockUserStorageMockRecorder is the mock recorder for MockUserStorage.
type MockUserStorageMockRecorder struct {
	mock *MockUserStorage
}

// NewMockUserStorage creates a new mock instance.
func NewMockUserStorage(ctrl *gomock.Controller) *MockUserStorage {
	mock := &MockUserStorage{ctrl: ctrl}
	mock.recorder = &MockUserStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStorage) EXPECT() *MockUserStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserStorage) Create(ctx context.Context, user model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserStorageMockRecorder) Create(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserStorage)(nil).Create), ctx, user)
}

// GetOneByLogin mocks base method.
func (m *MockUserStorage) GetOneByLogin(ctx context.Context, login string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneByLogin", ctx, login)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneByLogin indicates an expected call of GetOneByLogin.
func (mr *MockUserStorageMockRecorder) GetOneByLogin(ctx, login any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneByLogin", reflect.TypeOf((*MockUserStorage)(nil).GetOneByLogin), ctx, login)
}