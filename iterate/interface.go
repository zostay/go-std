package iterate

type Interface[I comparable, T any] interface {
	Len() int
	Next() bool
	ID() I
	Val() T
}
