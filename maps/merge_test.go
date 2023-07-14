package maps_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/maps"
)

func TestMerge(t *testing.T) {
	a := map[string]int{"a": 1, "b": 2}
	b := map[string]int{"c": 3, "d": 4}
	c := map[string]int{"a": 5, "b": 6, "c": 7}

	assert.Equal(t, map[string]int{}, maps.Merge[string, int]())
	assert.Equal(t, map[string]int{
		"a": 1, "b": 2,
	}, maps.Merge(a))
	assert.Equal(t, map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4,
	}, maps.Merge(a, b))
	assert.Equal(t, map[string]int{
		"a": 5, "b": 6, "c": 7,
	}, maps.Merge(a, c))
	assert.Equal(t, map[string]int{
		"a": 5, "b": 6, "c": 7, "d": 4,
	}, maps.Merge(a, b, c))
}

func TestDiff(t *testing.T) {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	b := map[string]int{"one": 5, "three": 6, "five": 7, "seven": 8}

	c, one, two := maps.Diff(a, b)
	assert.Equal(t, map[string]int{"one": 1, "three": 3}, c)
	assert.Equal(t, map[string]int{"two": 2, "four": 4}, one)
	assert.Equal(t, map[string]int{"five": 7, "seven": 8}, two)
}
