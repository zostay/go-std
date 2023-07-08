package generic

// FirstNonZero returns the first non-zero value in the list.
func FirstNonZero[T Comparable](values ...T) T {
	z := Zero[T]()
	for _, v := range values {
		if v != z {
			return v
		}
	}
	return z
}

// FirstNonNil returns the first non-nil value in the list.
func FirstNonNil[T any](values ...*T) *T {
	for _, v := range values {
		if v != nil {
			return v
		}
	}
	return nil
}

// Zero will return the zero-value of any type.
func Zero[T any]() (v T) {
	return
}
