// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zostay/go-std/fs (interfaces: CreateFS,WriteFileFS,FileWriter)

// Package fs_test is a generated GoMock package.
package fs_test

import (
	fs0 "io/fs"
	reflect "reflect"

	fs "github.com/zostay/go-std/fs"
	gomock "go.uber.org/mock/gomock"
)

// MockCreateFS is a mock of CreateFS interface.
type MockCreateFS struct {
	ctrl     *gomock.Controller
	recorder *MockCreateFSMockRecorder
}

// MockCreateFSMockRecorder is the mock recorder for MockCreateFS.
type MockCreateFSMockRecorder struct {
	mock *MockCreateFS
}

// NewMockCreateFS creates a new mock instance.
func NewMockCreateFS(ctrl *gomock.Controller) *MockCreateFS {
	mock := &MockCreateFS{ctrl: ctrl}
	mock.recorder = &MockCreateFSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateFS) EXPECT() *MockCreateFSMockRecorder {
	return m.recorder
}

// Chmod mocks base method.
func (m *MockCreateFS) Chmod(arg0 string, arg1 fs0.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chmod", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chmod indicates an expected call of Chmod.
func (mr *MockCreateFSMockRecorder) Chmod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chmod", reflect.TypeOf((*MockCreateFS)(nil).Chmod), arg0, arg1)
}

// Create mocks base method.
func (m *MockCreateFS) Create(arg0 string) (fs.FileWriter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(fs.FileWriter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCreateFSMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCreateFS)(nil).Create), arg0)
}

// Mkdir mocks base method.
func (m *MockCreateFS) Mkdir(arg0 string, arg1 fs0.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mkdir", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mkdir indicates an expected call of Mkdir.
func (mr *MockCreateFSMockRecorder) Mkdir(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mkdir", reflect.TypeOf((*MockCreateFS)(nil).Mkdir), arg0, arg1)
}

// Open mocks base method.
func (m *MockCreateFS) Open(arg0 string) (fs0.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(fs0.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockCreateFSMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockCreateFS)(nil).Open), arg0)
}

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

// Chmod mocks base method.
func (m *MockWriteFileFS) Chmod(arg0 string, arg1 fs0.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chmod", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chmod indicates an expected call of Chmod.
func (mr *MockWriteFileFSMockRecorder) Chmod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chmod", reflect.TypeOf((*MockWriteFileFS)(nil).Chmod), arg0, arg1)
}

// Create mocks base method.
func (m *MockWriteFileFS) Create(arg0 string) (fs.FileWriter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(fs.FileWriter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWriteFileFSMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWriteFileFS)(nil).Create), arg0)
}

// Mkdir mocks base method.
func (m *MockWriteFileFS) Mkdir(arg0 string, arg1 fs0.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mkdir", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mkdir indicates an expected call of Mkdir.
func (mr *MockWriteFileFSMockRecorder) Mkdir(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mkdir", reflect.TypeOf((*MockWriteFileFS)(nil).Mkdir), arg0, arg1)
}

// Open mocks base method.
func (m *MockWriteFileFS) Open(arg0 string) (fs0.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(fs0.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockWriteFileFSMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockWriteFileFS)(nil).Open), arg0)
}

// WriteFile mocks base method.
func (m *MockWriteFileFS) WriteFile(arg0 string, arg1 []byte, arg2 fs0.FileMode) error {
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

// MockFileWriter is a mock of FileWriter interface.
type MockFileWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFileWriterMockRecorder
}

// MockFileWriterMockRecorder is the mock recorder for MockFileWriter.
type MockFileWriterMockRecorder struct {
	mock *MockFileWriter
}

// NewMockFileWriter creates a new mock instance.
func NewMockFileWriter(ctrl *gomock.Controller) *MockFileWriter {
	mock := &MockFileWriter{ctrl: ctrl}
	mock.recorder = &MockFileWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileWriter) EXPECT() *MockFileWriterMockRecorder {
	return m.recorder
}

// Chmod mocks base method.
func (m *MockFileWriter) Chmod(arg0 fs0.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chmod", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chmod indicates an expected call of Chmod.
func (mr *MockFileWriterMockRecorder) Chmod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chmod", reflect.TypeOf((*MockFileWriter)(nil).Chmod), arg0)
}

// Close mocks base method.
func (m *MockFileWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockFileWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFileWriter)(nil).Close))
}

// Write mocks base method.
func (m *MockFileWriter) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockFileWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFileWriter)(nil).Write), arg0)
}
