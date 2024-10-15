package strings

import (
	"strings"
	"unicode/utf8"
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

// IncrementSet is the interface used for determining your sets. In addition to
// implementing this interface, an IncrementSet must be sure that First != Last
// and that repeated calls to Next starting with First will equal Last.
type IncrementSet interface {
	// Carry returns the rune to use as the rune to prefix with. The n value is
	// the number of digits that follow this value, which is useful for creating
	// strings that switch from one set to another at a certain digit.
	Carry(n int) rune

	// First returns the first rune in the set. This is the zero value for the
	// set.
	First() rune

	// Next returns the next rune in the set after the one given. If the given
	// rune is the last rune, it will return the first rune and true. It returns
	// -1 if the given rune is not in the set.
	Next(r rune) (rune, bool)

	// Last returns the last rune in the set. Incrementing one past this will
	// result in a carry.
	Last() rune
}

type incSet struct {
	numeric bool
	set     string
}

func (s incSet) Carry(int) rune {
	if s.numeric {
		cr, _ := s.Next(s.First())
		return cr
	}
	return s.First()
}

func (s incSet) First() rune {
	r, _ := utf8.DecodeRuneInString(s.set)
	return r
}

func (s incSet) Last() rune {
	r, _ := utf8.DecodeLastRuneInString(s.set)
	return r
}

func (s incSet) Next(r rune) (rune, bool) {
	if r == s.Last() {
		return s.First(), true
	}

	if ix := strings.IndexRune(s.set, r); ix >= 0 {
		nr, _ := utf8.DecodeRuneInString(s.set[ix+utf8.RuneLen(r):])
		return nr, false
	}

	return -1, false
}

// NumericSet constructs a numeric style set for use with IncrementWithSets.
// This is different from a LetterSet in the way it handles the Carry operation.
//
// When carrying to in front of the first current digit, a NumericSet uses
// Next(First) as the carry value.
func NumericSet(set string) IncrementSet {
	return incSet{true, set}
}

// NumericSetRange constructs a numeric set from the given rune range. See
// NumericSet for more details.
func NumericSetRange(a, z rune) IncrementSet {
	return incSet{true, FromRange(a, z)}
}

// LetterSet constructs a non-numeric style set for use with IncrementWithSets.
// This is differnet from a NumericSet in the way it handles the Carry
// operation.
//
// When carrying to in front of the first current digit, a LetterSet uses First
// as the carry value.
func LetterSet(set string) IncrementSet {
	return incSet{false, set}
}

// LetterSetRange constructs a non-numeric style set for the given rune range.
// See LetterSet for more details.
func LetterSetRange(a, z rune) IncrementSet {
	return incSet{false, FromRange(a, z)}
}

// SeqSet constructs a special set, which carries to different set types after a
// certain number of digits. For example, Kansas license plates are 3 letters
// followed by 3 numbers and could be (sort of) emulate like this:
//
//	output := strings.IncrementWithSets("",
//	    strings.SeqSet(3, 'A', strings.NumericSet("0123456789")),
//	    strings.LetterSet(strings.FromRange('A', 'Z')),
//	) // output := "1"
//	output := strings.IncrementWithSets("999",
//	    strings.SeqSet(3, 'A', strings.NumericSet("0123456789")),
//	    strings.LetterSet(strings.FromRange('A', 'Z')),
//	) // output := "A000"
//	output := strings.IncrementWithSets("ZZZ999",
//	    strings.SeqSet(3, 'A', strings.NumericSet("0123456789")),
//	    strings.LetterSet(strings.FromRange('A', 'Z')),
//	) // output := "AAAA000"
func SeqSet(limit int, limitCarry rune, inner IncrementSet) IncrementSet {
	return seqSet{
		IncrementSet: inner,
		limit:        limit,
		limitCarry:   limitCarry,
	}
}

type seqSet struct {
	IncrementSet
	limit      int
	limitCarry rune
}

func (s seqSet) Carry(n int) rune {
	if n >= s.limit {
		return s.limitCarry
	}
	return s.IncrementSet.Carry(n)
}

// IncrementWithSets returns the string incremented according to th range
// strings provided in the second argument. For example,
//
//	      output := strings.IncrementWithSets("A9",
//		      strings.LetterSet("0123456789ABCDEF"),
//		  )
//
// will set output to:
//
//	"AA"
//
// You can use it for spreadsheet style increments like:
//
//	output := strings.IncrementWithSets("AZ",
//	    strings.LetterSet(strings.FromRange('A', 'Z')),
//	) // output := "BA"
//
// The increment works by increasing the last byte in the string that matches
// one of the given ranges to the next rune in that range. If it matches the
// last rune in the range, then it causes a carry operation, which will
// increment the next rune in the string according to whichever range it belongs
// to. Carries will continue until the first rune that does not belong to a
// range is found. If that happens, then a new rune is inserted into the string
// from the range encountered most recently during the carry process (i.e, the
// first in the set of ranges).
//
// A couple more examples:
//
//		output := strings.IncrementWithSets("ZZZ999",
//		    strings.LetterSet(strings.FromRange('A', 'Z')),
//		    strings.NumericSet(strings.FromRange('0', '9')),
//		) // output := "AAAA000" // add a new rune from carrying
//		output := strings.IncrementWithSets("ID:ZZ9",
//	        strings.LetterSet(strings.FromRange('A', 'Z')),
//	        strings.NumericSet(strings.FromRange('0', '9')),
//		) // output := "ID:AAA0" // because ":" is in no range
//
// If the string is empty, the first set given is used to pick a 0 value.
func IncrementWithSets(in string, sets ...IncrementSet) string {
	out := ""
	prevSet := sets[0]

Digit:
	for {
		// either the string is empty or we carried beyond the first rune
		if in == "" {
			out = string(prevSet.Carry(len(out))) + out
			break
		}

		r, size := utf8.DecodeLastRuneInString(in)
		in = in[:len(in)-size]

		for _, set := range sets {
			if nr, carry := set.Next(r); nr >= 0 {
				out = string(nr) + out
				prevSet = set
				if carry {
					continue Digit
				} else {
					break Digit
				}
			}
		}

		// next rune is not in a set, so we just need to handle the final carry
		out = string(r) + string(prevSet.Carry(len(out))) + out
		break
	}

	return in + out
}

var defaultIncs = []IncrementSet{
	LetterSet("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	LetterSet("abcdefghijklmnopqrstuvwxyz"),
	NumericSet("0123456789"),
}

// Increment returns the string with the value incremented. It is exactly
// identical to running:
//
//	output := strings.IncrementWithSets(input,
//	    strings.LetterSet(strings.FromRange('A', 'Z')),
//	    strings.LetterSet(strings.FromRange('a', 'z')),
//	    strings.NumericSet(strings.FromRange('0', '9')),
//	)
//
// (Though, the ranges are pre-calculated.
func Increment(input string) string {
	return IncrementWithSets(input, defaultIncs...)
}

// Indent returns the string with each line indented by the given string.
func Indent(s, indent string) string {
	startIndent := indent
	if len(s) > 0 && s[0] == '\n' {
		startIndent = ""
	}
	return strings.TrimRight(
		startIndent+strings.ReplaceAll(s, "\n", "\n"+indent),
		" ")
}
