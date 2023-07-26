package fs_test

import (
	iofs "io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/zostay/go-std/fs"
)

func TestMkdirAllExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mkdir := NewMockCreateFS(ctrl)
	file := NewMockFile(ctrl)
	info := NewMockFileInfo(ctrl)

	mkdir.EXPECT().Open("foo/bar/baz").Return(file, nil)
	file.EXPECT().Stat().Return(info, nil)
	file.EXPECT().Close().Return(nil)
	info.EXPECT().IsDir().Return(true)

	err := fs.MkdirAll(mkdir, "foo/bar/baz", 0o755)

	assert.NoError(t, err)
}

func TestMkdirAllMkdirFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mkdir := NewMockCreateFS(ctrl)

	mkdir.EXPECT().Open("foo/bar/baz").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Open("foo/bar").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Open("foo").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Mkdir("foo", iofs.FileMode(0o755)).Return(iofs.ErrPermission)

	err := fs.MkdirAll(mkdir, "foo/bar/baz", 0o755)

	assert.Error(t, err)
}

func TestMkdirAllStatFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mkdir := NewMockCreateFS(ctrl)

	mkdir.EXPECT().Open("foo/bar/baz").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Open("foo/bar").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Open("foo").Return(nil, iofs.ErrPermission)

	err := fs.MkdirAll(mkdir, "foo/bar/baz", 0o755)

	assert.Error(t, err)
}

func TestMkdirAllParentIsFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mkdir := NewMockCreateFS(ctrl)
	file := NewMockFile(ctrl)
	info := NewMockFileInfo(ctrl)

	mkdir.EXPECT().Open("foo/bar/baz").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Open("foo/bar").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Open("foo").Return(file, nil)
	file.EXPECT().Stat().Return(info, nil)
	info.EXPECT().IsDir().Return(false)
	file.EXPECT().Close().Return(nil)

	err := fs.MkdirAll(mkdir, "foo/bar/baz", 0o755)

	assert.ErrorIs(t, err, fs.ErrNotDir)
}

func TestMkdirAllPartialExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mkdir := NewMockCreateFS(ctrl)
	file := NewMockFile(ctrl)
	info := NewMockFileInfo(ctrl)

	mkdir.EXPECT().Open("foo/bar/baz").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Open("foo/bar").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Open("foo").Return(file, nil)
	file.EXPECT().Stat().Return(info, nil)
	info.EXPECT().IsDir().Return(true)
	file.EXPECT().Close().Return(nil)
	mkdir.EXPECT().Mkdir("foo/bar", iofs.FileMode(0o755)).Return(nil)
	mkdir.EXPECT().Mkdir("foo/bar/baz", iofs.FileMode(0o755)).Return(nil)

	err := fs.MkdirAll(mkdir, "foo/bar/baz", 0o755)

	assert.NoError(t, err)
}

func TestMkdirAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mkdir := NewMockCreateFS(ctrl)

	mkdir.EXPECT().Open("foo/bar/baz").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Open("foo/bar").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Open("foo").Return(nil, iofs.ErrNotExist)
	mkdir.EXPECT().Mkdir("foo", iofs.FileMode(0o755)).Return(nil)
	mkdir.EXPECT().Mkdir("foo/bar", iofs.FileMode(0o755)).Return(nil)
	mkdir.EXPECT().Mkdir("foo/bar/baz", iofs.FileMode(0o755)).Return(nil)

	err := fs.MkdirAll(mkdir, "foo/bar/baz", 0o755)

	assert.NoError(t, err)
}

func TestWriteFileFS(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wf := NewMockWriteFileFS(ctrl)

	wf.EXPECT().WriteFile("foo", []byte("bar"), iofs.FileMode(0o644)).Return(nil)

	err := fs.WriteFile(wf, "foo", []byte("bar"), 0o644)

	assert.NoError(t, err)
}

func TestWriteFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfs := NewMockCreateFS(ctrl)
	file := NewMockFileWriter(ctrl)

	cfs.EXPECT().Create("foo").Return(file, nil)
	file.EXPECT().Write([]byte("bar")).Return(3, nil)
	file.EXPECT().Chmod(iofs.FileMode(0o644)).Return(nil)
	file.EXPECT().Close().Return(nil)

	err := fs.WriteFile(cfs, "foo", []byte("bar"), 0o644)

	assert.NoError(t, err)
}

func TestWriteFileCreateFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfs := NewMockCreateFS(ctrl)

	cfs.EXPECT().Create("foo").Return(nil, iofs.ErrPermission)

	err := fs.WriteFile(cfs, "foo", []byte("bar"), 0o644)

	assert.Error(t, err)
}

func TestWriteFileWriteFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfs := NewMockCreateFS(ctrl)
	file := NewMockFileWriter(ctrl)

	cfs.EXPECT().Create("foo").Return(file, nil)
	file.EXPECT().Write([]byte("bar")).Return(0, iofs.ErrPermission)
	file.EXPECT().Close().Return(nil)

	err := fs.WriteFile(cfs, "foo", []byte("bar"), 0o644)

	assert.Error(t, err)
}

func TestWriteFileChmodFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfs := NewMockCreateFS(ctrl)
	file := NewMockFileWriter(ctrl)

	cfs.EXPECT().Create("foo").Return(file, nil)
	file.EXPECT().Write([]byte("bar")).Return(3, nil)
	file.EXPECT().Chmod(iofs.FileMode(0o644)).Return(iofs.ErrPermission)
	file.EXPECT().Close().Return(nil)

	err := fs.WriteFile(cfs, "foo", []byte("bar"), 0o644)

	assert.Error(t, err)
}
