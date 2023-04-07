package maps_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/maps"
)

func TestCopyInit(t *testing.T) {
	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	var b map[string]int
	maps.CopyInit(&b, a)

	assert.Equal(t, a, b)
}

func TestCopy(t *testing.T) {
	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	b := map[string]int{}
	maps.Copy(b, a)

	assert.Equal(t, a, b)
}

func TestCopy_SadNilDstPanics(t *testing.T) {
	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	var b map[string]int
	assert.Panics(t, func() {
		maps.Copy(b, a)
	})
}

func TestClear(t *testing.T) {
	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	maps.Clear(a)
	assert.Equal(t, map[string]int{}, a)
}
