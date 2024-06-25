package bytes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/bytes"
)

func TestContainsOnly(t *testing.T) {
	t.Parallel()

	assert.True(t, bytes.ContainsOnly([]byte("abc"), []rune("abc")))
	assert.True(t, bytes.ContainsOnly([]byte("a"), []rune("abc")))
	assert.True(t, bytes.ContainsOnly([]byte("aaaaaabbbbbbccccc"), []rune("cba")))
	assert.False(t, bytes.ContainsOnly([]byte("a b c"), []rune("abc")))
}

func TestFromRange(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []byte("abcd"), bytes.FromRange('a', 'd'))
	assert.Equal(t, []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), bytes.FromRange('A', 'Z'))
	assert.Equal(t, []byte("0123456789"), bytes.FromRange('0', '9'))
	assert.Equal(t, []byte("zyxwvutsrqponmlkjihgfedcba"), bytes.FromRange('z', 'a'))
}

func TestReverse(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []byte(".txet gnisreveR"), bytes.Reverse([]byte("Reversing text.")))
}

func TestIndent(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []byte("\n    a\n    b\n    c\n"), bytes.Indent([]byte("\na\nb\nc\n"), []byte("    ")))
	assert.Equal(t, []byte("    a\n    b\n    c"), bytes.Indent([]byte("a\nb\nc"), []byte("    ")))
}
