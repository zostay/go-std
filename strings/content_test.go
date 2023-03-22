package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsOnly(t *testing.T) {
	assert.True(t, ContainsOnly("abc", "abc"))
	assert.True(t, ContainsOnly("a", "abc"))
	assert.True(t, ContainsOnly("aaaaaabbbbbbccccc", "cba"))
	assert.False(t, ContainsOnly("a b c", "abc"))
}
