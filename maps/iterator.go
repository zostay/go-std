package maps

type Iterator[K comparable, V any] struct {
	data   map[K]V
	index  []K
	cursor int
}

func NewIterator[K comparable, V any](data map[K]V) *Iterator[K, V] {
	index := make([]K, 0, len(data))
	for k := range data {
		index = append(index, k)
	}
	return &Iterator[K, V]{data, index, -1}
}

func (i *Iterator[K, V]) Len() int {
	return len(i.data)
}

func (i *Iterator[K, V]) Next() bool {
	i.cursor++
	return i.cursor < len(i.data)
}

func (i *Iterator[K, V]) Index() int {
	return i.cursor
}

func (i *Iterator[K, V]) ID() K {
	return i.index[i.cursor]
}

func (i *Iterator[K, V]) Val() V {
	return i.data[i.index[i.cursor]]
}
