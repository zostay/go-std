package generic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/generic"
)

func TestCountDeltas(t *testing.T) {
	t.Parallel()

	d := generic.CountDeltas(1.0, 10.0, 1.0)
	assert.Equal(t, 9, d)

	d = generic.CountDeltas(10.0, 1.0, -1.0)
	assert.Equal(t, 9, d)

	d = generic.CountDeltas(1.1, 10.0, 1.0)
	assert.Equal(t, 8, d)

	d = generic.CountDeltas(1.0, 2.0, 1.01)
	assert.Equal(t, 0, d)

	d = generic.CountDeltas(1, 10, 1)
	assert.Equal(t, 9, d)

	d = generic.CountDeltas(1, 10, 2)
	assert.Equal(t, 4, d)

	d = generic.CountDeltas(1, 2, 2)
	assert.Equal(t, 0, d)

	d = generic.CountDeltas(1.0, 10.0, 2.0)
	assert.Equal(t, 4, d)
}
