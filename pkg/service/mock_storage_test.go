// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kamilsk/click/pkg/service (interfaces: Storage)

// Package service_test is a generated GoMock package.
package service_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	domain "github.com/kamilsk/click/pkg/domain"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Link mocks base method
func (m *MockStorage) Link(arg0 context.Context, arg1 domain.ID) (domain.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Link", arg0, arg1)
	ret0, _ := ret[0].(domain.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Link indicates an expected call of Link
func (mr *MockStorageMockRecorder) Link(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Link", reflect.TypeOf((*MockStorage)(nil).Link), arg0, arg1)
}

// LinkByAlias mocks base method
func (m *MockStorage) LinkByAlias(arg0 context.Context, arg1 domain.ID, arg2 string) (domain.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LinkByAlias", arg0, arg1, arg2)
	ret0, _ := ret[0].(domain.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkByAlias indicates an expected call of LinkByAlias
func (mr *MockStorageMockRecorder) LinkByAlias(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkByAlias", reflect.TypeOf((*MockStorage)(nil).LinkByAlias), arg0, arg1, arg2)
}
