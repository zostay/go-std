package fs

import "io/fs"

//go:generate mockgen -package=fs_test -destination=writer_mock_test.go github.com/zostay/go-std/fs WriteFileFS,MkdirFS

// CreateFS is a file system that can create files.
type CreateFS interface {
	Create(name string) (fs.File, error)
}

// WriteFileFS is a file system that can write files whole.
type WriteFileFS interface {
	CreateFS

	WriteFile(name string, data []byte, perm fs.FileMode) error
}

// MkdirFS is a file system that can create directories.
type MkdirFS interface {
	CreateFS

	Mkdir(name string, perm fs.FileMode) error
}

// WriterFS combines the file system writer interfaces into a single interface
// for convenience.
type WriterFS interface {
	WriteFileFS
	MkdirFS
}

// WriteFile writes out a complete file, creating or updating the file in place
// using the given bytes. The file mode will be set as given as well. Returns an
// error on failure.
func WriteFile(fsys WriteFileFS, name string, data []byte, perm fs.FileMode) error {
	return fsys.WriteFile(name, data, perm)
}

// Mkdir creates a directory with the given name and permission in the file
// system specified. Returns an error on failure.
func Mkdir(fsys MkdirFS, name string, perm fs.FileMode) error {
	return fsys.Mkdir(name, perm)
}
