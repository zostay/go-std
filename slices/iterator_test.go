package slices_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/slices"
)

func TestNewIterator(t *testing.T) {
	a := []int{1, 2, 3}

	iter := slices.NewIterator(a)
	assert.Equal(t, len(a), iter.Len())
	var i int
	for iter.Next() {
		assert.Equal(t, i, iter.Index())
		assert.Equal(t, i, iter.ID())
		assert.Equal(t, a[i], iter.Val())
		i++
	}

	assert.Panics(t, func() {
		iter.ID()
	})
}
