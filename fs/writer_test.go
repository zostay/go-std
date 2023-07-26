package fs_test

import (
	iofs "io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/zostay/go-std/fs"
)

func TestMkdir(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mkdir := NewMockMkdirFS(ctrl)

	mkdir.EXPECT().Mkdir("foo", iofs.FileMode(0o755)).Return(nil)

	err := fs.Mkdir(mkdir, "foo", 0o755)

	assert.NoError(t, err)
}

func TestWriteFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wf := NewMockWriteFileFS(ctrl)

	wf.EXPECT().WriteFile("foo", []byte("bar"), iofs.FileMode(0o644)).Return(nil)

	err := fs.WriteFile(wf, "foo", []byte("bar"), 0o644)

	assert.NoError(t, err)
}
