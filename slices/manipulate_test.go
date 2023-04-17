package slices_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/slices"
)

func TestDelete(t *testing.T) {
	var slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice = slices.Delete[int](slice, 7)
	slice = slices.Delete[int](slice, 4)
	assert.Equal(t, []int{0, 1, 2, 3, 5, 6, 8, 9}, slice)

	slice = slices.Delete[int](slice, 3)
	slice = slices.Delete[int](slice, 6)
	assert.Equal(t, []int{0, 1, 2, 5, 6, 8}, slice)

	slice = slices.Delete[int](slice, 0)
	assert.Equal(t, []int{1, 2, 5, 6, 8}, slice)

	assert.Panics(t, func() {
		slice = slices.Delete[int](slice, -1)
	})

	assert.Panics(t, func() {
		slice = slices.Delete[int](slice, 5)
	})

	slice = slices.Delete[int](slice, 0)
	slice = slices.Delete[int](slice, 0)
	slice = slices.Delete[int](slice, 0)
	slice = slices.Delete[int](slice, 0)
	slice = slices.Delete[int](slice, 0)

	assert.Empty(t, slice)

	assert.Panics(t, func() {
		slice = slices.Delete[int](slice, 0)
	})
}

func TestDeleteValue(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 6}

	s = slices.DeleteValue(s, 0)
	s = slices.DeleteValue(s, 3)
	s = slices.DeleteValue(s, 6)
	s = slices.DeleteValue(s, 9)

	assert.Equal(t, []int{1, 2, 4, 5, 7, 8, 6}, s)

	s = slices.DeleteValue(s, 1)
	s = slices.DeleteValue(s, 6)

	assert.Equal(t, []int{2, 4, 5, 7, 8}, s)

	s = slices.DeleteValue(s, 8)
	s = slices.DeleteValue(s, 5)
	s = slices.DeleteValue(s, 2)
	s = slices.DeleteValue(s, 4)
	s = slices.DeleteValue(s, 7)

	assert.Empty(t, s)

	assert.NotPanics(t, func() {
		s = slices.DeleteValue(s, 0)
	})
}

func TestDeleteAllValues(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 6}

	s = slices.DeleteAllValues(s, 0)
	s = slices.DeleteAllValues(s, 3)
	s = slices.DeleteAllValues(s, 6)
	s = slices.DeleteAllValues(s, 9)

	assert.Equal(t, []int{1, 2, 4, 5, 7, 8}, s)

	s = slices.DeleteAllValues(s, 1)
	s = slices.DeleteAllValues(s, 6)

	assert.Equal(t, []int{2, 4, 5, 7, 8}, s)

	s = slices.DeleteAllValues(s, 8)
	s = slices.DeleteAllValues(s, 5)
	s = slices.DeleteAllValues(s, 2)
	s = slices.DeleteAllValues(s, 4)
	s = slices.DeleteAllValues(s, 7)

	assert.Empty(t, s)

	assert.NotPanics(t, func() {
		s = slices.DeleteAllValues(s, 0)
	})
}

func TestPop(t *testing.T) {
	s := []int{1, 2, 3}

	v, s := slices.Pop(s)
	assert.Equal(t, []int{1, 2}, s)
	assert.Equal(t, 3, v)

	v, s = slices.Pop(s)
	assert.Equal(t, []int{1}, s)
	assert.Equal(t, 2, v)

	v, s = slices.Pop(s)
	assert.Equal(t, []int{}, s)
	assert.Equal(t, 1, v)

	assert.Panics(t, func() {
		_, _ = slices.Pop(s)
	})
}

func TestPush(t *testing.T) {
	s := []int{}

	s = slices.Push(s, 1)
	s = slices.Push(s, 2, 3)

	assert.Equal(t, []int{1, 2, 3}, s)
}

func TestShift(t *testing.T) {
	s := []int{1, 2, 3}

	v, s := slices.Shift(s)
	assert.Equal(t, []int{2, 3}, s)
	assert.Equal(t, 1, v)

	v, s = slices.Shift(s)
	assert.Equal(t, []int{3}, s)
	assert.Equal(t, 2, v)

	v, s = slices.Shift(s)
	assert.Equal(t, []int{}, s)
	assert.Equal(t, 3, v)

	assert.Panics(t, func() {
		_, _ = slices.Shift(s)
	})
}

func TestUnshift(t *testing.T) {
	s := []int{}

	s = slices.Unshift(s, 1)
	s = slices.Unshift(s, 2, 3)

	assert.Equal(t, []int{2, 3, 1}, s)
}

func TestConcat(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5}
	c := []int{}
	d := []int{6, 7, 8}

	e := slices.Concat(a, b, c, d)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, e)
}
