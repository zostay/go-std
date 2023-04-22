package slices

import "github.com/zostay/go-std/generic"

// Head returns the first count items from the input list as a newly copied list.
func Head[T any](in []T, count int) []T {
	out := make([]T, generic.Min(count, len(in)))
	copy(out, in)
	return out
}

// Tail returns the last count items from the input list as a newly copied list.
func Tail[T any](in []T, count int) []T {
	out := make([]T, generic.Min(count, len(in)))
	copy(out, in[len(in)-len(out):])
	return out
}
