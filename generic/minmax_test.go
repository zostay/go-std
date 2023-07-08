package generic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/generic"
)

func TestCompare(t *testing.T) {
	a := 42
	b := 13

	assert.Greater(t, generic.Compare[int](a, b), generic.Equal)
	assert.Less(t, generic.Compare[int](b, a), generic.Equal)
	assert.Equal(t, generic.Compare[int](a, a), generic.Equal)
	assert.Equal(t, generic.Compare[int](b, b), generic.Equal)
}

func TestLess(t *testing.T) {
	a := 42
	b := 13

	assert.True(t, generic.Less[int](b, a))
	assert.False(t, generic.Less[int](a, b))
}

func TestMax(t *testing.T) {
	a := 42
	b := 13

	assert.Equal(t, a, generic.Max[int](a, b))
	assert.Equal(t, a, generic.Max[int](b, a))
}

func TestMin(t *testing.T) {
	a := 42
	b := 13

	assert.Equal(t, b, generic.Min[int](a, b))
	assert.Equal(t, b, generic.Min[int](b, a))
}

type compare int

func (a compare) Compare(b generic.ComparableInterface) generic.Comparison {
	return generic.Comparison(a - b.(compare))
}

func TestLessI(t *testing.T) {
	a := compare(42)
	b := compare(13)

	assert.True(t, generic.LessI(b, a))
	assert.False(t, generic.LessI(a, b))
}

func TestMaxI(t *testing.T) {
	a := compare(42)
	b := compare(13)

	assert.Equal(t, a, generic.MaxI(a, b))
	assert.Equal(t, a, generic.MaxI(b, a))
}

func TestMinI(t *testing.T) {
	a := compare(42)
	b := compare(13)

	assert.Equal(t, b, generic.MinI(a, b))
	assert.Equal(t, b, generic.MinI(b, a))
}
