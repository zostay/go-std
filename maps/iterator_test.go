package maps_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/maps"
)

func TestNewIterator(t *testing.T) {
	t.Parallel()

	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	var expect map[string]int
	maps.CopyInit(&expect, a)

	iter := maps.NewIterator(a)
	assert.Equal(t, len(a), iter.Len())
	var i int
	for iter.Next() {
		assert.Equal(t, i, iter.Index())
		assert.Contains(t, expect, iter.ID())
		assert.Equal(t, expect[iter.ID()], iter.Val())
		delete(expect, iter.ID())
		i++
	}
	assert.Empty(t, expect)
}
