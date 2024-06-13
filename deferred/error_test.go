package deferred_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/deferred"
)

func TestError(t *testing.T) {
	t.Parallel()

	errors := []error{
		fmt.Errorf("error 1"),
		fmt.Errorf("error 2"),
	}

	cases := []struct {
		first, second error
		name          string
	}{
		{nil, nil, "nil-nil"},
		{errors[0], nil, "error-nil"},
		{nil, errors[0], "nil-error"},
		{errors[0], errors[1], "error-error"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			err := c.first
			deferred.Error(&err, c.second)

			if c.first == nil && c.second == nil {
				assert.NoError(t, err, c.name)
				return
			}

			assert.Error(t, err)
			if c.first != nil {
				assert.ErrorIs(t, err, c.first)
			}
			if c.second != nil {
				assert.ErrorIs(t, err, c.second)
			}
		})
	}
}
