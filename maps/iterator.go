package maps

// Iterator implements iterable.Interface for maps.
type Iterator[K comparable, V any] struct {
	data   map[K]V
	index  []K
	cursor int
}

// NewIterator returns a fresh iterator for any map. If the map is modified
// while iteration is being performed, you may encounter unexpected results.
func NewIterator[K comparable, V any](data map[K]V) *Iterator[K, V] {
	index := make([]K, 0, len(data))
	for k := range data {
		index = append(index, k)
	}
	return &Iterator[K, V]{data, index, -1}
}

// Len returns the number of elements that will be iterated.
func (i *Iterator[K, V]) Len() int {
	return len(i.data)
}

// Next moves the iterator cursor to the next item and increments Index.
func (i *Iterator[K, V]) Next() bool {
	i.cursor++
	return i.cursor < len(i.data)
}

// Index returns the index of the current item being iterated over.
func (i *Iterator[K, V]) Index() int {
	return i.cursor
}

// ID returns the map key for the current item.
func (i *Iterator[K, V]) ID() K {
	return i.index[i.cursor]
}

// Val returns the map value for the current item.
func (i *Iterator[K, V]) Val() V {
	return i.data[i.index[i.cursor]]
}
