package slices

import "github.com/zostay/go-std/generic"

// FromRange returns a slice containing all the values values from a to z
// with incremented delta difference between.
func FromRange[T generic.Number](a, z, delta T) []T {
	if a > z {
		a, z = z, a
	}

	if delta < 0 {
		delta = -delta
	}

	output := make([]T, 0, generic.CountDeltas(a, z, delta)+1)
	for i := a; i <= z; i += delta {
		output = append(output, i)
	}

	return output
}
