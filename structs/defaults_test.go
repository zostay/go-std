package structs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/structs"
)

func TestApplyDefaults(t *testing.T) {
	type test struct {
		A string
		B int
		C *float64
	}

	var (
		f   = float64(13.0)
		g   = float64(14.0)
		one = test{
			A: "foo",
			B: 42,
			C: &f,
		}
		two = test{
			A: "",
			B: 43,
			C: &g,
		}
		three = test{
			A: "bar",
			B: 0,
			C: nil,
		}
	)

	v, d := one, two
	structs.ApplyDefaults(&v, &d)
	assert.Equal(t, one, v)

	v, d = one, three
	structs.ApplyDefaults(&v, &d)
	assert.Equal(t, one, v)

	v, d = two, three
	structs.ApplyDefaults(&v, &d)
	assert.Equal(t, test{A: "bar", B: 43, C: &g}, v)

	v, d = two, one
	structs.ApplyDefaults(&v, &d)
	assert.Equal(t, test{A: "foo", B: 43, C: &g}, v)

	v, d = two, three
	structs.ApplyDefaults(&v, &d)
	assert.Equal(t, test{A: "bar", B: 43, C: &g}, v)

	v, d = three, one
	structs.ApplyDefaults(&v, &d)
	assert.Equal(t, test{A: "bar", B: 42, C: &f}, v)

	v, d = three, two
	structs.ApplyDefaults(&v, &d)
	assert.Equal(t, test{A: "bar", B: 43, C: &g}, v)
}
