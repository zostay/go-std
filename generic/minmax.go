package generic

// Comparison is an integer that represents who two values compare to each
// other. If Comparison is negative, then the first value is less than the
// second. If Comparison is positive, then the first value is greater than the
// second. If Comparison is zero, then the first value is equal to the second.
type Comparison int

// The following generic Comparison constants are provided for convenience when
// defining a Compare method. Any negative value is considered less-than and any
// positive value is considered greater-than. Those are just two possible values.
const (
	LessThan    Comparison = -1 // one possible less-than value
	Equal       Comparison = 0  // the equivalent value
	GreaterThan Comparison = 1  // one possible greater-than value
)

// ComparableInterface defines a generic way of making objects comparable.
// Implementation of this interface should be very easy. Here's a quick sample:
//
//	type MyInt int
//	func (a MyInt) Compare(b generic.ComparableInterface) generic.Comparison {
//	  return generic.Comparison(a-b.(MyInt))
//	}
//
// It assumed that a ComparableInterface will never be compared to another
// ComparableInterface of a different type. If you do, it will likely result in
// a panic unless tha ComparableInterface implementation allows comparison for
// different types. The implementor is responsible for keeping the code safe in
// such a case.
type ComparableInterface interface {
	Compare(b ComparableInterface) Comparison
}

// Comparable identifies all those types that can be compared using <, >, <=,
// and >= operations.
//
// This is what "comparable" means in most languages whereas in Golang,
// "comparable" only means values of those types can be checked for equivalence.
type Comparable interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~uint | ~int
}

// Compare provides a generic comparison function for any two Comparable values
// of the same type. This method will return a negative value if a is less than
// b, a positive value if a is greater than b, and 0 if the values are
// equivalent. (Please note, that some of these types, notably float32 and
// float64, are technically comparable for equivalence, but might result in a
// non-zero value even if they are close enough to be equivalent.)
func Compare[T Comparable](a, b T) Comparison {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	}
	return 0
}

// Max returns the larger of the two values given. It will prefer the first if
// equal.
func Max[T Comparable](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

// MaxI returns the larger of the two values given. It will prefer the first if
// equal.
func MaxI(a, b ComparableInterface) ComparableInterface {
	if a.Compare(b) >= 0 {
		return a
	}
	return b
}

// Min returns the smaller of the two values given. It will prefer the first if
// the two values are equal.
func Min[T Comparable](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// MinI returns the smaller of the two values given. It will prefer the first if
// the two values are equal.
func MinI(a, b ComparableInterface) ComparableInterface {
	if a.Compare(b) <= 0 {
		return a
	}
	return b
}

// Less provides a generic Less function for any Comparable.
func Less[T Comparable](a, b T) bool {
	return a < b
}

// LessI provides a generic Less function for any ComparableInterface.
func LessI(a, b ComparableInterface) bool {
	return a.Compare(b) < 0
}

// Zero will return the zero-value of any type.
func Zero[T any]() (v T) {
	return
}
