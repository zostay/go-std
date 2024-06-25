package bytes

import "bytes"

// ContainsOnly returns true fi the given byte slice only contains the given
// letters.
func ContainsOnly(b []byte, chars []rune) bool {
	idx := make(map[rune]struct{}, len(chars))
	for _, c := range chars {
		idx[c] = struct{}{}
	}
	return bytes.IndexFunc(b, func(c rune) bool {
		_, expectedChar := idx[c]
		return !expectedChar
	}) == -1
}

// FromRange returns a byte slice containing all the characters between two given
// bytes (inclusive).
func FromRange(a, z byte) []byte {
	reverse := false
	if a > z {
		reverse = true
		a, z = z, a
	}

	out := make([]byte, z-a+1)
	for i := a; i <= z; i++ {
		out[i-a] = i
	}

	if reverse {
		return Reverse(out)
	}

	return out
}

// Reverse returns the byte slice reversed. That is, passing the byte slice
// []byte("Reversing text.") will result in this function returning
// []byte(".txet gnisreveR").
func Reverse(in []byte) []byte {
	out := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = in[len(in)-i-1]
	}
	return out
}

// Indent returns a new byte slice with each line in the input byte slice
// indented by the given byte slice.
func Indent(b, indent []byte) []byte {
	startIndent := indent
	if b[0] == '\n' {
		startIndent = []byte{}
	}
	newB := bytes.ReplaceAll(b, []byte("\n"), []byte("\n"+string(indent)))
	if len(startIndent) > 0 {
		newB = append(newB, startIndent...)
		copy(newB[len(startIndent):], newB)
		copy(newB, startIndent)
	}
	return bytes.TrimRight(newB, " ")
}
