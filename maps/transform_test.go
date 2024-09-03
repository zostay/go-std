package maps_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/maps"
)

func TestKeys(t *testing.T) {
	t.Parallel()

	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	keys := maps.Keys(a)
	assert.Len(t, keys, len(a))
	for k := range a {
		assert.Contains(t, keys, k)
	}
}

func TestValues(t *testing.T) {
	t.Parallel()

	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	values := maps.Values(a)
	assert.Len(t, values, len(a))
	for _, v := range a {
		assert.Contains(t, values, v)
	}
}

func TestKVs(t *testing.T) {
	t.Parallel()

	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	kvs := maps.KVs(a)
	assert.Len(t, kvs, len(a)*2)
	for k, v := range a {
		assert.Contains(t, kvs, k)
		assert.Contains(t, kvs, v)

		for i := range kvs {
			if kvs[i] == k {
				assert.Equal(t, v, kvs[i+1])
			}
		}
	}

	for i, x := range kvs {
		if i%2 == 0 {
			assert.IsType(t, "", x)
		} else {
			assert.IsType(t, 0, x)
		}
	}
}

func TestFlip(t *testing.T) {
	t.Parallel()

	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	vks := maps.Flip(a)
	assert.Len(t, vks, len(a))
	for k, v := range a {
		assert.Contains(t, vks, v)
		assert.Equal(t, k, vks[v])
	}

	for k, v := range vks {
		assert.Contains(t, a, v)
		assert.Equal(t, k, a[v])
	}
}
