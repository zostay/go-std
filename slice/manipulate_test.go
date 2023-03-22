package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	var slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice = Delete[int](slice, 7)
	slice = Delete[int](slice, 4)
	assert.Equal(t, []int{0, 1, 2, 3, 5, 6, 8, 9}, slice)

	slice = Delete[int](slice, 3)
	slice = Delete[int](slice, 6)
	assert.Equal(t, []int{0, 1, 2, 5, 6, 8}, slice)

	slice = Delete[int](slice, 0)
	assert.Equal(t, []int{1, 2, 5, 6, 8}, slice)

	assert.Panics(t, func() {
		slice = Delete[int](slice, -1)
	})

	assert.Panics(t, func() {
		slice = Delete[int](slice, 5)
	})

	slice = Delete[int](slice, 0)
	slice = Delete[int](slice, 0)
	slice = Delete[int](slice, 0)
	slice = Delete[int](slice, 0)
	slice = Delete[int](slice, 0)

	assert.Empty(t, slice)

	assert.Panics(t, func() {
		slice = Delete[int](slice, 0)
	})
}
