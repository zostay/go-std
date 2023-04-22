package slices

import "github.com/zostay/go-std/generic"

// Map transforms the input into the output value using the mapper function.
func Map[In, Out any](in []In, mapper func(in In) Out) []Out {
	out := make([]Out, len(in))
	for i, t := range in {
		out[i] = mapper(t)
	}
	return out
}

// Reduce collapses the input slice into a single value using the reducer
// function.
func Reduce[In, Out any](in []In, reducer func(u Out, t In) Out) Out {
	var out Out
	for _, t := range in {
		out = reducer(out, t)
	}
	return out
}

// ReduceAcc collapses the input slice into a single value using the reducer
// function, but uses the given starting value as the starting point (rather
// than the type's natural zero value).
func ReduceAcc[In, Out any](in []In, acc Out, reducer func(u Out, t In) Out) Out {
	for _, t := range in {
		acc = reducer(acc, t)
	}
	return acc
}

// Reductions is identical to Reduce, but it returns all the intermediate steps.
func Reductions[In, Out any](in []In, reducer func(u Out, t In) Out) []Out {
	out := make([]Out, len(in))
	var prev Out
	for i, t := range in {
		out[i] = reducer(prev, t)
		prev = out[i]
	}
	return out
}

// ReductionsAcc is identical to ReduceAcc, but it returns all the intermediate
// steps.
func ReductionsAcc[In, Out any](in []In, acc Out, reducer func(u Out, t In) Out) []Out {
	out := make([]Out, len(in))
	prev := acc
	for i, t := range in {
		out[i] = reducer(prev, t)
		prev = out[i]
	}
	return out
}

// Sum adds all the values that are in the input slice together and returns the
// result. If there are no values, this returns the zero value.
func Sum[T generic.Summable](in []T) T {
	var acc T
	for _, val := range in {
		acc += val
	}
	return acc
}

// Product multiples all the values that are in the input slice together and
// returns the result. If there are no values, this will return 1.
func Product[T generic.Number](in []T) T {
	acc := T(1)
	for _, val := range in {
		acc *= val
	}
	return acc
}

// Grep applies the predicate to each item and returns a new list of items
// containing only those that match the predicate.
func Grep[T any](in []T, pred func(t T) bool) []T {
	out := make([]T, 0, len(in))
	for _, t := range in {
		if pred(t) {
			out = append(out, t)
		}
	}
	return out
}

// Any returns true if any item in the input slice matches the predicate.
// Returns false if the slice is empty.
func Any[T any](in []T, pred func(t T) bool) bool {
	for _, t := range in {
		if pred(t) {
			return true
		}
	}
	return false
}

// All returns true if all items in the input slice matches the predicate.
// Returns true if the slice is empty.
func All[T any](in []T, pred func(t T) bool) bool {
	for _, t := range in {
		if !pred(t) {
			return false
		}
	}
	return true
}

// None is the opposite of All. It returns true if none of the input items
// matches the predicate. Returns true if the slice is empty.
func None[T any](in []T, pred func(t T) bool) bool {
	for _, t := range in {
		if pred(t) {
			return false
		}
	}
	return true
}

// NotAll is the opposite of Any. It returns true if any of the input items does
// not match the predicate. Returns false if the slice is empty.
func NotAll[T any](in []T, pred func(t T) bool) bool {
	for _, t := range in {
		if !pred(t) {
			return true
		}
	}
	return false
}

// First returns the first item that matches the predicate. If a value is found,
// the second value returned will be true. If no match is found, this returns a
// zero value and false.
func First[T any](in []T, pred func(t T) bool) (first T, found bool) {
	for _, t := range in {
		if pred(t) {
			return t, true
		}
	}
	return
}

// FirstOr returns the first tiem that matches the predicate. If nothing
// matches, the missed value is returned.
func FirstOr[T any](in []T, missed T, pred func(t T) bool) T {
	for _, t := range in {
		if pred(t) {
			return t
		}
	}
	return missed
}
