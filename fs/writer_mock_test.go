// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zostay/go-std/fs (interfaces: WriteFileFS,MkdirFS)

// Package fs_test is a generated GoMock package.
package fs_test

import (
	fs "io/fs"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockWriteFileFS is a mock of WriteFileFS interface.
type MockWriteFileFS struct {
	ctrl     *gomock.Controller
	recorder *MockWriteFileFSMockRecorder
}

// MockWriteFileFSMockRecorder is the mock recorder for MockWriteFileFS.
type MockWriteFileFSMockRecorder struct {
	mock *MockWriteFileFS
}

// NewMockWriteFileFS creates a new mock instance.
func NewMockWriteFileFS(ctrl *gomock.Controller) *MockWriteFileFS {
	mock := &MockWriteFileFS{ctrl: ctrl}
	mock.recorder = &MockWriteFileFSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriteFileFS) EXPECT() *MockWriteFileFSMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWriteFileFS) Create(arg0 string) (fs.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(fs.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWriteFileFSMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWriteFileFS)(nil).Create), arg0)
}

// WriteFile mocks base method.
func (m *MockWriteFileFS) WriteFile(arg0 string, arg1 []byte, arg2 fs.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile.
func (mr *MockWriteFileFSMockRecorder) WriteFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockWriteFileFS)(nil).WriteFile), arg0, arg1, arg2)
}

// MockMkdirFS is a mock of MkdirFS interface.
type MockMkdirFS struct {
	ctrl     *gomock.Controller
	recorder *MockMkdirFSMockRecorder
}

// MockMkdirFSMockRecorder is the mock recorder for MockMkdirFS.
type MockMkdirFSMockRecorder struct {
	mock *MockMkdirFS
}

// NewMockMkdirFS creates a new mock instance.
func NewMockMkdirFS(ctrl *gomock.Controller) *MockMkdirFS {
	mock := &MockMkdirFS{ctrl: ctrl}
	mock.recorder = &MockMkdirFSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMkdirFS) EXPECT() *MockMkdirFSMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMkdirFS) Create(arg0 string) (fs.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(fs.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMkdirFSMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMkdirFS)(nil).Create), arg0)
}

// Mkdir mocks base method.
func (m *MockMkdirFS) Mkdir(arg0 string, arg1 fs.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mkdir", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mkdir indicates an expected call of Mkdir.
func (mr *MockMkdirFSMockRecorder) Mkdir(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mkdir", reflect.TypeOf((*MockMkdirFS)(nil).Mkdir), arg0, arg1)
}
