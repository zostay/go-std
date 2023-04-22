package slices_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/slices"
)

func TestHead(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, slices.Head(3, []int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{1, 2, 3}, slices.Head(3, []int{1, 2, 3}))
	assert.Equal(t, []int{1, 2}, slices.Head(3, []int{1, 2}))
	assert.Equal(t, []int{}, slices.Head(3, []int{}))
	assert.Equal(t, []int{1}, slices.Head(1, []int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{}, slices.Head(0, []int{1, 2, 3, 4, 5}))
}

func TestTail(t *testing.T) {
	assert.Equal(t, []int{3, 4, 5}, slices.Tail(3, []int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{1, 2, 3}, slices.Tail(3, []int{1, 2, 3}))
	assert.Equal(t, []int{1, 2}, slices.Tail(3, []int{1, 2}))
	assert.Equal(t, []int{}, slices.Tail(3, []int{}))
	assert.Equal(t, []int{5}, slices.Tail(1, []int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{}, slices.Tail(0, []int{1, 2, 3, 4, 5}))
}
