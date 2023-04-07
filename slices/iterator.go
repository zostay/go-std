package slices

// Iterator implements an iterator over the given slice. If the undlying slice
// is modified during iteration, the results of iteration may lead to unexpected
// results or panics.
type Iterator[T any] struct {
	slice  []T
	cursor int
}

// NewIterator constructs and returns a new iterable.Interface for the given
// object.
func NewIterator[T any](slice []T) *Iterator[T] {
	return &Iterator[T]{slice, -1}
}

// Len returns the number of items in the underlying slice.
func (i *Iterator[T]) Len() int {
	return len(i.slice)
}

// Next moves the iterator to the next element of the slice.
func (i *Iterator[T]) Next() bool {
	i.cursor++
	return i.cursor < len(i.slice)
}

// Index returns the index of the current element of the slice being iterated.
func (i *Iterator[T]) Index() int {
	return i.ID()
}

// ID returns the index of the current element of the slice being iterated.
func (i *Iterator[T]) ID() int {
	if i.cursor < 0 || i.cursor >= len(i.slice) {
		panic("attempt to retrieve out of bounds iterator ID")
	}
	return i.cursor
}

// Val returns the value of the current element of the slice being itereated.
func (i *Iterator[T]) Val() T {
	return i.slice[i.cursor]
}
