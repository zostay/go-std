package fs

// ReaderWriterFS combines all the reader and writer file system interfaces
// into a single interface for convenience.
type ReaderWriterFS interface {
	ReaderFS
	WriterFS
}
