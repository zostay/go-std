package generic

// Comparable identifies all those types that can be compared using <, >, <=,
// and >= operations.
//
// This is what "comparable" means in most languages whereas in Golang,
// "comparable" only means values of those types can be checked for equivalence.
type Comparable interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~uint | ~int | ~string
}

// Number identifies all those types that are numeric and can be added,
// subtracted, multiplied, and divided.
type Number interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~uint | ~int
}

// Summable identifies all those types that have an add operator.
type Summable interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~uint | ~int | ~string
}
