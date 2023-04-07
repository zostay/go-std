package slices

type Iterator[T any] struct {
	slice  []T
	cursor int
}

func NewIterator[T any](slice []T) *Iterator[T] {
	return &Iterator[T]{slice, -1}
}

func (i *Iterator[T]) Len() int {
	return len(i.slice)
}

func (i *Iterator[T]) Next() bool {
	i.cursor++
	return i.cursor < len(i.slice)
}

func (i *Iterator[T]) Index() int {
	return i.ID()
}

func (i *Iterator[T]) ID() int {
	if i.cursor < 0 || i.cursor >= len(i.slice) {
		panic("attempt to retrieve out of bounds iterator ID")
	}
	return i.cursor
}

func (i *Iterator[T]) Val() T {
	return i.slice[i.cursor]
}
