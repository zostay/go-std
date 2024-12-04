package strings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zostay/go-std/strings"
)

func TestContainsOnly(t *testing.T) {
	t.Parallel()

	assert.True(t, strings.ContainsOnly("abc", "abc"))
	assert.True(t, strings.ContainsOnly("a", "abc"))
	assert.True(t, strings.ContainsOnly("aaaaaabbbbbbccccc", "cba"))
	assert.False(t, strings.ContainsOnly("a b c", "abc"))
}

func TestTrimToOnly(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "abc", strings.TrimToOnly("abc", "abc"))
	assert.Equal(t, "a", strings.TrimToOnly("a", "abc"))
	assert.Equal(t, "aaaaaabbbbbbccccc", strings.TrimToOnly("aaaaaabbbbbbccccc", "cba"))
	assert.Equal(t, "abc", strings.TrimToOnly("a b c", "abc"))
	assert.Equal(t, "", strings.TrimToOnly("abc", "def"))
	assert.Equal(t, "abde", strings.TrimToOnly("abcdef", "adbe"))
}

func TestFromRange(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "abcd", strings.FromRange('a', 'd'))
	assert.Equal(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", strings.FromRange('A', 'Z'))
	assert.Equal(t, "0123456789", strings.FromRange('0', '9'))
	assert.Equal(t, "zyxwvutsrqponmlkjihgfedcba", strings.FromRange('z', 'a'))
}

func TestReverse(t *testing.T) {
	t.Parallel()

	assert.Equal(t, ".txet gnisreveR", strings.Reverse("Reversing text."))
}

func TestIncrement(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "A", strings.Increment(""))
	assert.Equal(t, "B", strings.Increment("A"))
	assert.Equal(t, "AA", strings.Increment("Z"))
	assert.Equal(t, "BA", strings.Increment("AZ"))
	assert.Equal(t, "b", strings.Increment("a"))
	assert.Equal(t, "aa", strings.Increment("z"))
	assert.Equal(t, "ba", strings.Increment("az"))
	assert.Equal(t, "1", strings.Increment("0"))
	assert.Equal(t, "2", strings.Increment("1"))
	assert.Equal(t, "10", strings.Increment("9"))
	assert.Equal(t, "19", strings.Increment("18"))
	assert.Equal(t, "20", strings.Increment("19"))
	assert.Equal(t, "B0", strings.Increment("A9"))
	assert.Equal(t, "ID:A", strings.Increment("ID:"))
	assert.Equal(t, "ID:B", strings.Increment("ID:A"))
	assert.Equal(t, "ID:AA", strings.Increment("ID:Z"))
	assert.Equal(t, "ID:BA", strings.Increment("ID:AZ"))
	assert.Equal(t, "ID:b", strings.Increment("ID:a"))
	assert.Equal(t, "ID:aa", strings.Increment("ID:z"))
	assert.Equal(t, "ID:ba", strings.Increment("ID:az"))
	assert.Equal(t, "ID:1", strings.Increment("ID:0"))
	assert.Equal(t, "ID:2", strings.Increment("ID:1"))
	assert.Equal(t, "ID:10", strings.Increment("ID:9"))
	assert.Equal(t, "ID:19", strings.Increment("ID:18"))
	assert.Equal(t, "ID:20", strings.Increment("ID:19"))
	assert.Equal(t, "ID:B0", strings.Increment("ID:A9"))

	licPlate := []strings.IncrementSet{
		strings.SeqSet(3, 'A', strings.NumericSetRange('0', '9')),
		strings.LetterSetRange('A', 'Z'),
	}
	assert.Equal(t, "1", strings.IncrementWithSets("", licPlate...))
	assert.Equal(t, "2", strings.IncrementWithSets("1", licPlate...))
	assert.Equal(t, "10", strings.IncrementWithSets("9", licPlate...))
	assert.Equal(t, "20", strings.IncrementWithSets("19", licPlate...))
	assert.Equal(t, "999", strings.IncrementWithSets("998", licPlate...))
	assert.Equal(t, "A000", strings.IncrementWithSets("999", licPlate...))
	assert.Equal(t, "A001", strings.IncrementWithSets("A000", licPlate...))
	assert.Equal(t, "B000", strings.IncrementWithSets("A999", licPlate...))
	assert.Equal(t, "AA000", strings.IncrementWithSets("Z999", licPlate...))
	assert.Equal(t, "AAA000", strings.IncrementWithSets("ZZ999", licPlate...))
	assert.Equal(t, "AAAA000", strings.IncrementWithSets("ZZZ999", licPlate...))
	assert.Equal(t, "ID:1", strings.IncrementWithSets("ID:", licPlate...))
	assert.Equal(t, "ID:2", strings.IncrementWithSets("ID:1", licPlate...))
	assert.Equal(t, "ID:10", strings.IncrementWithSets("ID:9", licPlate...))
	assert.Equal(t, "ID:20", strings.IncrementWithSets("ID:19", licPlate...))
	assert.Equal(t, "ID:999", strings.IncrementWithSets("ID:998", licPlate...))
	assert.Equal(t, "ID:A000", strings.IncrementWithSets("ID:999", licPlate...))
	assert.Equal(t, "ID:A001", strings.IncrementWithSets("ID:A000", licPlate...))
	assert.Equal(t, "ID:B000", strings.IncrementWithSets("ID:A999", licPlate...))
	assert.Equal(t, "ID:AA000", strings.IncrementWithSets("ID:Z999", licPlate...))
	assert.Equal(t, "ID:AAA000", strings.IncrementWithSets("ID:ZZ999", licPlate...))
	assert.Equal(t, "ID:AAAA000", strings.IncrementWithSets("ID:ZZZ999", licPlate...))
}

func TestIndent(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "", strings.Indent("", "    "))
	assert.Equal(t, "\n    a\n    b\n    c\n", strings.Indent("\na\nb\nc\n", "    "))
	assert.Equal(t, "    a\n    b\n    c", strings.Indent("a\nb\nc", "    "))
}
