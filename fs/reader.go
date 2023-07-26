package fs

import "io/fs"

// ReaderFS combines all teh reader file system interfaces into a single
// interface for convenience.
type ReaderFS interface {
	fs.GlobFS
	fs.ReadDirFS
	fs.ReadFileFS
	fs.StatFS
}
