package slices_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/slices"
)

func TestFromRange(t *testing.T) {
	gotI := slices.FromRange(1, 10, 1)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, gotI)

	gotI = slices.FromRange(10, 1, -1)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, gotI)

	gotF := slices.FromRange(1.0, 10.0, 1.0)
	assert.Equal(t, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}, gotF)

	gotF = slices.FromRange(10.0, 1.0, -1.0)
	assert.Equal(t, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}, gotF)

	gotF = slices.FromRange(1.1, 10.0, 1.0)
	assert.Equal(t, []float64{1.1, 2.1, 3.1, 4.1, 5.1, 6.1, 7.1, 8.1, 9.1}, gotF)
}
