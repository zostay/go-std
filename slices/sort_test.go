package slices_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/generic"
	"github.com/zostay/go-std/slices"
)

func TestShuffle(t *testing.T) {
	t.Parallel()

	in := []int{1, 2, 3, 4, 5}

	slices.Shuffle(in)
	assert.Len(t, in, 5)
	assert.Contains(t, in, 1)
	assert.Contains(t, in, 2)
	assert.Contains(t, in, 3)
	assert.Contains(t, in, 4)
	assert.Contains(t, in, 5)
}

func TestSample(t *testing.T) {
	t.Parallel()

	ins := [][]int{
		{},
		{1, 2},
		{1, 2, 3, 4, 5},
	}

	for _, in := range ins {
		is := slices.Sample(in, 3)

		// all values are from the original list
		assert.Len(t, is, generic.Min(len(in), 3))
		for _, v := range is {
			assert.Contains(t, in, v)
		}

		// no duplicates
		sort.Ints(is)
		p := -1
		for _, v := range is {
			assert.NotEqual(t, v, p)
			p = v
		}
	}
}

func TestUniq(t *testing.T) {
	t.Parallel()

	in := []int{1, 1, 1, 2, 2, 3, 4, 4, 5, 6, 7, 7, 7}
	out := slices.Uniq(in)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, out)
	assert.Equal(t, []int{}, slices.Uniq([]int{}))
}

func TestUniqInPlace(t *testing.T) {
	t.Parallel()

	in := []int{1, 1, 1, 2, 2, 3, 4, 4, 5, 6, 7, 7, 7}
	in = slices.UniqInPlace(in)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, in)
	assert.Equal(t, []int{}, slices.UniqInPlace([]int{}))
}
