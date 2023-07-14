package generic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/generic"
)

func TestFirstNonZero(t *testing.T) {
	assert.Equal(t, 1, generic.FirstNonZero[int](0, 1, 2))
	assert.Equal(t, 1, generic.FirstNonZero[int](1, 0, 2))
	assert.Equal(t, 1, generic.FirstNonZero[int](0, 0, 1))
	assert.Equal(t, 0, generic.FirstNonZero[int](0, 0, 0))
}

func TestFirstNonNil(t *testing.T) {
	var a *int
	b := new(int)
	c := new(int)
	*b = 1
	*c = 2

	assert.Equal(t, *b, *generic.FirstNonNil(a, b, c))
	assert.Equal(t, *b, *generic.FirstNonNil(b, a, c))
	assert.Equal(t, *b, *generic.FirstNonNil(a, a, b))
	assert.Nil(t, generic.FirstNonNil[int](nil, nil, nil))
}

func TestZero(t *testing.T) {
	assert.Equal(t, 0, generic.Zero[int]())
	assert.Equal(t, "", generic.Zero[string]())
	assert.Equal(t, rune(0), generic.Zero[rune]())
}
