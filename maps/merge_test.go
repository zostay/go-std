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
