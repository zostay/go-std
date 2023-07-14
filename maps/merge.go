package maps

import "github.com/zostay/go-std/generic"

// Merge maps into a single map. Keys in maps later in the list will overwrite
// keys found in earlier maps.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	out := map[K]V{}
	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}

// Diff returns three maps from two. The first map returned is a map of
// key/value pairs where the keys are found in both the first and second input
// maps. The second map returned contains a map of key/value pairs for keys only
// found in the first map. The third map returned contains a map of key/value
// pairs for keys only found in the third map.
//
// If values differ for keys held in common, the value in the first input map
// are kept.
func Diff[K comparable, V any](a, b map[K]V) (common, inFirst, inSecond map[K]V) {
	common = make(map[K]V, generic.Max(len(a), len(b)))
	inFirst = make(map[K]V, len(a))
	inSecond = make(map[K]V, len(b))

	for akey, avalue := range a {
		if _, found := b[akey]; found {
			common[akey] = avalue
		} else {
			inFirst[akey] = avalue
		}
	}

	for bkey, bvalue := range b {
		if _, found := a[bkey]; !found {
			inSecond[bkey] = bvalue
		}
	}

	return
}
