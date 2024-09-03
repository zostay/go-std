package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/set"
)

func TestFlipMap(t *testing.T) {
	t.Parallel()

	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 1,
		"e": 1,
	}

	vks := set.FlipMap(a)
	assert.Len(t, vks, 3)
	for k, v := range a {
		assert.Contains(t, vks, v)
		assert.True(t, vks[v].Contains(k))
	}

	for k, vs := range vks {
		for _, v := range vs.Keys() {
			assert.Contains(t, a, v)
			assert.Equal(t, k, a[v])
		}
	}
}
