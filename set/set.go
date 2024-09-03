package set

import (
	"github.com/zostay/go-std/generic"
	"github.com/zostay/go-std/maps"
)

// Set provides a generic set data type. A set is an unsorted list of objects
// where each object is unique, as defined by the equivalence operation. These
// are implemented using a map. Any value that can be used as a map key can be
// used as a set object.
//
// This is a finite set implementation. You must explicitly add values to the
// set.
type Set[T comparable] map[T]struct{}

var exists = struct{}{}

// New constructs a new set containing the given values.
func New[T comparable](vals ...T) Set[T] {
	out := make(Set[T], len(vals))
	for _, v := range vals {
		out[v] = exists
	}
	return out
}

// NewSized constructs a new set that is empty, but has the requested capacity.
func NewSized[T comparable](capacity int) Set[T] {
	return make(Set[T], capacity)
}

// Contains returns true if the given value is contained within the set.
func (s Set[T]) Contains(val T) bool {
	_, hasVal := s[val]
	return hasVal
}

// All returns true if all the given values are contained within the set.
// This is similar to running:
//
//	s.Intersection(set.New(vals...)).Len()==len(vals)
//
// This is both more efficient and easier to read in many cases.
func (s Set[T]) All(vals ...T) bool {
	for _, val := range vals {
		if _, contains := s[val]; !contains {
			return false
		}
	}
	return true
}

// Any returns true if any of the given values are contained within the
// set. This is similar to running:
//
//	s.Intersection(set.New(vals...)).Len()>0
//
// This is both more efficient and easier to read in many cases.
func (s Set[T]) Any(vals ...T) bool {
	for _, val := range vals {
		if _, contains := s[val]; contains {
			return true
		}
	}
	return false
}

// None returns true if none of the given values are contained within the
// set. This is similar to running:
//
//	s.Intersection(set.New(vals...)).Len()==0
//
// This is both more efficient and easier to read in many cases.
func (s Set[T]) None(vals ...T) bool {
	for _, val := range vals {
		if _, contains := s[val]; contains {
			return false
		}
	}
	return true
}

// Insert adds the given value to the set. If the value is already present, this
// will have no effect.
func (s Set[T]) Insert(vals ...T) {
	for _, val := range vals {
		s[val] = exists
	}
}

// Delete removes the given value from the set. If the value is not present,
// this will have no effect.
func (s Set[T]) Delete(val ...T) {
	for _, val := range val {
		delete(s, val)
	}
}

// Len returns the size of the set.
func (s Set[T]) Len() int {
	return len(s)
}

// SubsetOf returns true if this set is entirely contained within the other set
// o.
func (s Set[T]) SubsetOf(o Set[T]) bool {
	for val := range s {
		if !o.Contains(val) {
			return false
		}
	}
	return true
}

// Keys returns all the keys in the set as a slice. The return slice is
// unsorted and the keys returned could be in any order.
func (s Set[T]) Keys() []T {
	return maps.Keys(s)
}

// Intersects returns true if any value in set a is found within set b.
func Intersects[T comparable](a, b Set[T]) bool {
	for val := range a {
		if b.Contains(val) {
			return true
		}
	}
	return false
}

// Intersection constructs and returns a new set containing all the values held
// in common between sets a and b.
func Intersection[T comparable](a, b Set[T]) Set[T] {
	minLen := len(a)
	if minLen > len(b) {
		minLen = len(b)
	}

	out := make(Set[T], minLen)
	for val := range a {
		if b.Contains(val) {
			out[val] = struct{}{}
		}
	}

	return out
}

// Union constructs and returns a new set containing all the values in set a as
// well as all the values in set b.
func Union[T comparable](a, b Set[T]) Set[T] {
	maxLen := len(a)
	if maxLen < len(b) {
		maxLen = len(b)
	}

	out := make(Set[T], maxLen)
	Copy[T](out, a)
	Copy[T](out, b)
	return out
}

// Difference constructs and returns a new set containing all the values in set
// a that are not also in set b.
func Difference[T comparable](a, b Set[T]) Set[T] {
	out := make(Set[T], len(a))
	for val := range a {
		if b.Contains(val) {
			continue
		}
		out[val] = struct{}{}
	}
	return out
}

// Copy will add all values from set src to set dst. Existing values in dst will
// remain. The dst set must be allocated or this will panic.
func Copy[T comparable](dst, src Set[T]) {
	for k := range src {
		dst[k] = struct{}{}
	}
}

// CopyInit will check to see if dst is already allocated. If not, it will
// initialize dst and then perform a Copy(dst, src).
func CopyInit[T comparable](dst *Set[T], src Set[T]) {
	if *dst == nil {
		*dst = make(Set[T], len(src))
	}
	Copy(*dst, src)
}

// Diff returns three sets from two. The first set returned is items held in
// common by both sets. The second set is items found in the first set, but not
// the second set. The third contains items found in the second set, but not the
// first.
func Diff[T comparable](a, b Set[T]) (common, inFirst, inSecond Set[T]) {
	common = NewSized[T](generic.Max(len(a), len(b)))
	inFirst = NewSized[T](len(a))
	inSecond = NewSized[T](len(b))

	for akey := range a {
		if _, found := b[akey]; found {
			common.Insert(akey)
		} else {
			inFirst.Insert(akey)
		}
	}

	for bkey := range b {
		if _, found := a[bkey]; !found {
			inSecond.Insert(bkey)
		}
	}

	return
}
