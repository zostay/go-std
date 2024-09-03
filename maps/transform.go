package maps

// Keys will return all the keys of the input map. The keys returned are not
// sorted.
func Keys[K comparable, V any](in map[K]V) []K {
	out := make([]K, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

// Values will return all the values of the input map. The values returend are
// not sorted.
func Values[K comparable, V any](in map[K]V) []V {
	out := make([]V, 0, len(in))
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

// KVs will return all the keys and values in an array. The even-indexed
// elements will the keys and the element following is the value that belongs to
// that key.
func KVs[K comparable, V any](in map[K]V) []any {
	out := make([]any, 0, len(in)*2)
	for k, v := range in {
		out = append(out, k, v)
	}
	return out
}

// Flip will return a new map with the keys and values of the input map flipped.
// That is, the keys of the input map will be the values of the output map and
// the values of the input map will be the keys of the output map.
//
// This will only work if the value type is comparable and if values are
// repeated, it is indeterminate which key will be kept.
func Flip[K comparable, V comparable](in map[K]V) map[V]K {
	out := make(map[V]K, len(in))
	for k, v := range in {
		out[v] = k
	}
	return out
}

// FlipSlice will return a new map with the keys and values of the input map
// flipped. That is, the keys of the input map will be the values of the output
// map and the values of the input map will be the keys of the output map.
// However, the resulting value is a slice of keys.
//
// This will only work if the value type is comparable. If values repeat, the
// keys will be collected into the value slice. The ordering is indeterminate.
//
// You may also be interested in set.FlipMap.
func FlipSlice[K comparable, V comparable](in map[K]V) map[V][]K {
	out := make(map[V][]K, len(in))
	for k, v := range in {
		outv, alreadyExists := out[v]
		if !alreadyExists {
			outv = make([]K, 0, 1)
		}
		outv = append(outv, k)
		out[v] = outv
	}
	return out
}
