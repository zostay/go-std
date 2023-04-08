package strings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/strings"
)

func TestContainsOnly(t *testing.T) {
	assert.True(t, strings.ContainsOnly("abc", "abc"))
	assert.True(t, strings.ContainsOnly("a", "abc"))
	assert.True(t, strings.ContainsOnly("aaaaaabbbbbbccccc", "cba"))
	assert.False(t, strings.ContainsOnly("a b c", "abc"))
}

func TestFromRange(t *testing.T) {
	assert.Equal(t, "abcd", strings.FromRange('a', 'd'))
	assert.Equal(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", strings.FromRange('A', 'Z'))
	assert.Equal(t, "0123456789", strings.FromRange('0', '9'))
	assert.Equal(t, "zyxwvutsrqponmlkjihgfedcba", strings.FromRange('z', 'a'))
}

func TestReverse(t *testing.T) {
	assert.Equal(t, ".txet gnisreveR", strings.Reverse("Reversing text."))
}
