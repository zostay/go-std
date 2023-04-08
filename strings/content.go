package strings

import (
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

// func (rs runeSlice) Less(i, j int) bool {
// 	return rs[i] < rs[j]
// }

func (rs runeSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

// func (rs runeSlice) Len() int {
// 	return len(rs)
// }

// FromRange returns a string containing all the characters between two given
// characters (inclusive).
func FromRange(a, z rune) string {
	reverse := false
	if a > z {
		reverse = true
		a, z = z, a
	}

	out := make([]rune, z-a+1)
	for i := a; i <= z; i++ {
		out[i-a] = i
	}

	if reverse {
		return Reverse(string(out))
	}

	return string(out)
}

// Reverse returns the string reversed. That is, passing the string "Reversing
// text." will result in this function returning ".txet gnisreveR".
func Reverse(in string) string {
	out := runeSlice(in)
	for i := 0; i < len(in)-i-1; i++ {
		out.Swap(i, len(in)-i-1)
	}
	return string(out)
}
