package strings

import (
	"sort"
	"strings"
)

// ContainsOnly returns true if the given string only contains the given
// letters.
func ContainsOnly(s, chars string) bool {
	idx := make(map[rune]struct{}, len(chars))
	for _, c := range chars {
		idx[c] = struct{}{}
	}
	return strings.IndexFunc(s, func(c rune) bool {
		_, expectedChar := idx[c]
		return !expectedChar
	}) == -1
}

type runeSlice []rune

func (rs runeSlice) Less(i, j int) bool {
	return rs[i] < rs[j]
}

func (rs runeSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs runeSlice) Len() int {
	return len(rs)
}

// FromRange returns a string containing all the characters between two given
// characters (inclusive).
func FromRange(a, z rune) string {
	reverse := false
	if a > z {
		reverse = true
		a, z = z, a
	}

	out := make([]rune, z-a)
	for i := a; i < z; i++ {
		out[i-a] = i
	}

	if reverse {
		sort.Reverse(runeSlice(out))
	}

	return string(out)
}
