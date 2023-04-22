package slices

// Delete will remove the data at index i from the given slice. This works by
// copying the data after it back and then shortening the slice by one element.
func Delete[T any](slice []T, i int) []T {
	if i != len(slice)-1 {
		copy(slice[i:], slice[i+1:])
	}
	return slice[:len(slice)-1]
}

// DeleteValue will find the first element in the slice with value s. It will
// remove that value from the slice by copying all elements after it back one
// index and then shortening the slice by one element.
func DeleteValue[T comparable](slice []T, s T) []T {
	for i, v := range slice {
		if v == s {
			return Delete[T](slice, i)
		}
	}
	return slice
}

// DeleteAllValues will delete every value s in the given slice. It works the
// same way as DeleteValue, but performs the deletion task for every element
// matching the given s, not just the first.
func DeleteAllValues[T comparable](slice []T, s T) []T {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == s {
			slice = Delete[T](slice, i)
		}
	}
	return slice
}

// Push will add the given values to the end of the slice and return the newly
// transformed slice. The slice must be allocated or this function will panic.
func Push[T any](slice []T, v ...T) []T {
	return append(slice, v...)
}

// Pop will remove one value from the end of the slice and return it and the
// newly transform slice (shorter by one value). The slice must be allocated and
// be non-empty or this operation will panic.
func Pop[T any](slice []T) (T, []T) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

// Shift will remove the first item from the slice and return it and the newly
// transformed slice (which is one shorter). The slice must be allocated and be
// non-empty or this function will panic.
func Shift[T any](slice []T) (T, []T) {
	return slice[0], slice[1:]
}

// Unshift will add the given values to the front of the slice and return the
// newly transformed slice. The slice must be allocated or this function will
// panic. The values will be added in the order given.
func Unshift[T any](slice []T, vs ...T) []T {
	var zero T
	for range vs {
		slice = append(slice, zero)
	}
	copy(slice[len(vs):], slice[:])
	for i, v := range vs {
		slice[i] = v
	}
	return slice
}

// Concat will concatenate any set of 0 or more slices into a single slice
// containing the values of all the passed slices.
func Concat[T any](slices ...[]T) []T {
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}

	out := make([]T, totalLen)
	start := 0
	for _, s := range slices {
		copy(out[start:], s)
		start += len(s)
	}

	return out
}
