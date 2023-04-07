package iterate

type Interface[I comparable, T any] interface {
	Len() int
	Next() bool
	Index() int
	ID() I
	Val() T
}
