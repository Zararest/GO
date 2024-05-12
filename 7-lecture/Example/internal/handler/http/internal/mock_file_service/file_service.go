// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package mock_file_service is a generated GoMock package.
package mock_file_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockfileService is a mock of fileService interface.
type MockfileService struct {
	ctrl     *gomock.Controller
	recorder *MockfileServiceMockRecorder
}

// MockfileServiceMockRecorder is the mock recorder for MockfileService.
type MockfileServiceMockRecorder struct {
	mock *MockfileService
}

// NewMockfileService creates a new mock instance.
func NewMockfileService(ctrl *gomock.Controller) *MockfileService {
	mock := &MockfileService{ctrl: ctrl}
	mock.recorder = &MockfileServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockfileService) EXPECT() *MockfileServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockfileService) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockfileServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockfileService)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockfileService) Get(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockfileServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockfileService)(nil).Get), arg0)
}

// Upload mocks base method.
func (m *MockfileService) Upload(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockfileServiceMockRecorder) Upload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockfileService)(nil).Upload), arg0, arg1)
}
