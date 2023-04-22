package slices

import "math/rand"

// Shuffle uses the built-in rand.Shuffle function to shuffle a slice in place
// without needing to define a swap function. It returns the input argument
// after shuffling that slice in place.
func Shuffle[T any](in []T) []T {
	rand.Shuffle(len(in), func(i, j int) {
		in[i], in[j] = in[j], in[i]
	})
	return in
}

// Sample selects count items from the list randomly and returns a new list with
// the selected items. If count is larger than the size of the input list, this
// works the same as Shuffle.
func Sample[T any](in []T, count int) []T {
	// TODO this is not very efficient
	out := make([]T, len(in))
	copy(out, in)
	rand.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})
	if len(out) > count {
		out = out[:count]
	}
	return out
}

// Uniq creates a new list copied from the input, but with subsequent duplicate
// items removed.
func Uniq[T comparable](in []T) []T {
	if len(in) == 0 {
		return []T{}
	}

	out := make([]T, 1, len(in))
	out[0] = in[0]

	for i := 1; i < len(in); i++ {
		if in[i] != out[len(out)-1] {
			out = append(out, in[i])
		}
	}

	return out
}

// UniqInPlace deletes subsequent duplicate items from the source list and
// returns the now modified (possibly shorter) slice.
func UniqInPlace[T comparable](in []T) []T {
	if len(in) <= 1 {
		return in
	}

	for i := len(in) - 1; i > 0; i-- {
		if in[i] == in[i-1] {
			in = Delete(in, i)
		}
	}

	return in
}
