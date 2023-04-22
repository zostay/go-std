package slices_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/slices"
)

func TestHead(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, slices.Head([]int{1, 2, 3, 4, 5}, 3))
	assert.Equal(t, []int{1, 2, 3}, slices.Head([]int{1, 2, 3}, 3))
	assert.Equal(t, []int{1, 2}, slices.Head([]int{1, 2}, 3))
	assert.Equal(t, []int{}, slices.Head([]int{}, 3))
	assert.Equal(t, []int{1}, slices.Head([]int{1, 2, 3, 4, 5}, 1))
	assert.Equal(t, []int{}, slices.Head([]int{1, 2, 3, 4, 5}, 0))
}

func TestTail(t *testing.T) {
	assert.Equal(t, []int{3, 4, 5}, slices.Tail([]int{1, 2, 3, 4, 5}, 3))
	assert.Equal(t, []int{1, 2, 3}, slices.Tail([]int{1, 2, 3}, 3))
	assert.Equal(t, []int{1, 2}, slices.Tail([]int{1, 2}, 3))
	assert.Equal(t, []int{}, slices.Tail([]int{}, 3))
	assert.Equal(t, []int{5}, slices.Tail([]int{1, 2, 3, 4, 5}, 1))
	assert.Equal(t, []int{}, slices.Tail([]int{1, 2, 3, 4, 5}, 0))
}
