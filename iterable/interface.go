package iterable

// Interface is a generic iterator that can be used to iterate over just about
// any interface. The generic I component is the index type and the T is the
// value type.
type Interface[I comparable, T any] interface {
	// Len returns the number of items that will be iterated.
	Len() int

	// Next increases index by one moves the iterator to the next item to
	// iterate over.
	Next() bool

	// Index returns a numeric index of the value being iterated. Iterators must
	// return zero for the first item and increment it by exactly one with each
	// call to Next.
	Index() int

	// ID is the identifier value for current item.
	ID() I

	// Val is the value of the current item.
	Val() T
}
