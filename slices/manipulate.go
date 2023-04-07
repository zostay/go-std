package slices

func Delete[T any](slice []T, i int) []T {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func DeleteValue[T comparable](slice []T, s T) []T {
	for i, v := range slice {
		if v == s {
			return Delete[T](slice, i)
		}
	}
	return slice
}

func DeleteAllValues[T comparable](slice []T, s T) []T {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == s {
			slice = Delete[T](slice, i)
		}
	}
	return slice
}

func Push[T any](slice []T, v ...T) []T {
	return append(slice, v...)
}

func Pop[T any](slice []T) (T, []T) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

func Shift[T any](slice []T) (T, []T) {
	return slice[0], slice[1:]
}

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
