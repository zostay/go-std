package slice

func Delete[T any](slice []T, i int) []T {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func Push[T any](slice []T, v T) []T {
	return append(slice, v)
}

func Pop[T any](slice []T) (T, []T) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

func Shift[T any](slice []T) (T, []T) {
	return slice[0], slice[1:]
}

func Unshift[T any](slice []T, v T) []T {
	var zero T
	slice = append(slice, zero)
	copy(slice[1:], slice[:])
	slice[0] = v
	return slice
}
