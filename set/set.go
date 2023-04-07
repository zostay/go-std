package set

type Set[T comparable] map[T]struct{}

var exists = struct{}{}

func New[T comparable](vals ...T) Set[T] {
	out := make(Set[T], len(vals))
	for _, v := range vals {
		out[v] = exists
	}
	return out
}

func (s Set[T]) Contains(val T) bool {
	_, hasVal := s[val]
	return hasVal
}

func (s Set[T]) Insert(val T) {
	s[val] = exists
}

func (s Set[T]) Delete(val T) {
	delete(s, val)
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) SubsetOf(o Set[T]) bool {
	for val := range s {
		if !o.Contains(val) {
			return false
		}
	}
	return true
}

func Intersects[T comparable](a, b Set[T]) bool {
	for val := range a {
		if b.Contains(val) {
			return true
		}
	}
	return false
}

func Intersection[T comparable](a, b Set[T]) Set[T] {
	minLen := len(a)
	if minLen > len(b) {
		minLen = len(b)
	}

	out := make(Set[T], minLen)
	for val := range a {
		if b.Contains(val) {
			out[val] = struct{}{}
		}
	}

	return out
}

func Union[T comparable](a, b Set[T]) Set[T] {
	maxLen := len(a)
	if maxLen < len(b) {
		maxLen = len(b)
	}

	out := make(Set[T], maxLen)
	Copy[T](out, a)
	Copy[T](out, b)
	return out
}

func Difference[T comparable](a, b Set[T]) Set[T] {
	out := make(Set[T], len(a))
	for val := range a {
		if b.Contains(val) {
			continue
		}
		out[val] = struct{}{}
	}
	return out
}

func Copy[T comparable](dst, src Set[T]) {
	for k := range src {
		dst[k] = struct{}{}
	}
}
