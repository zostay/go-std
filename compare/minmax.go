package compare

type AbleInterface interface {
	Less(b AbleInterface) bool
}

type Able interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~uint | ~int
}

func Max[T Able](a, b T) T {
	if b < a {
		return a
	}
	return b
}

func MaxI(a, b AbleInterface) AbleInterface {
	if b.Less(a) {
		return a
	}
	return b
}

func Min[T Able](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MinI(a, b AbleInterface) AbleInterface {
	if b.Less(a) {
		return a
	}
	return b
}

func Less[T Able](a, b T) bool {
	return a < b
}

func LessI(a, b AbleInterface) bool {
	return a.Less(b)
}
