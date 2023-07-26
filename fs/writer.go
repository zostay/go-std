package fs

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
)

//go:generate mockgen -package=fs_test -destination=writer_mock_test.go github.com/zostay/go-std/fs CreateFS,WriteFileFS,FileWriter
//go:generate mockgen -package=fs_test -destination=iofs_mock_test.go io/fs File,FileInfo

var ErrNotDir = errors.New("parent is not a directory")

// FileWriter is a writable file.
type FileWriter interface {
	Chmod(fs.FileMode) error   // change the file mode of the file
	Write([]byte) (int, error) // write data to the file
	Close() error              // close the file for writing
}

// CreateFS is a file system that can create files.
type CreateFS interface {
	fs.FS

	Create(name string) (FileWriter, error)    // create or append to a file
	Chmod(name string, mode fs.FileMode) error // change the mode of the file
	Mkdir(name string, perm fs.FileMode) error // create a directory
}

// WriteFileFS is a file system that can write files whole.
type WriteFileFS interface {
	CreateFS

	// WriteFile writes out a complete file, creating or updating the file in
	// place using the given bytes. The file mode will be set as given as well.
	// Returns an error on failure.
	WriteFile(name string, data []byte, perm fs.FileMode) error
}

// WriterFS combines the file system writer interfaces into a single interface
// for convenience.
type WriterFS interface {
	WriteFileFS
}

// WriteFile writes out a complete file, creating or updating the file in place
// using the given bytes. The file mode will be set as given as well. Returns an
// error on failure.
func WriteFile(fsys CreateFS, name string, data []byte, perm fs.FileMode) error {
	if wfsys, isWriteFileFS := fsys.(WriteFileFS); isWriteFileFS {
		return wfsys.WriteFile(name, data, perm)
	}

	file, err := fsys.Create(name)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	err = file.Chmod(perm)
	if err != nil {
		return err
	}

	return nil
}

// MkdirAll will ensure that a directory exists with the given name, creating
// all parent directories as needed. Returns an error on failure.
func MkdirAll(fsys CreateFS, name string, perm fs.FileMode) error {
	if info, err := fs.Stat(fsys, name); err == nil && info.IsDir() {
		return nil
	}

	need := []string{name}
	unravel := false
	for len(need) > 0 {
		curPath := need[len(need)-1]
		dir := path.Dir(curPath)
		if unravel || dir == "." {
			err := fsys.Mkdir(curPath, perm)
			if err != nil {
				return err
			}

			unravel = true
			need = need[:len(need)-1]
			continue
		}

		info, err := fs.Stat(fsys, dir)
		if errors.Is(err, fs.ErrNotExist) {
			need = append(need, dir)
			continue
		}
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return fmt.Errorf("%w: %s", ErrNotDir, dir)
		}

		// path exists, so we can unravel from here
		unravel = true
	}

	return nil
}
