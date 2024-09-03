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
func Flip[K comparable, V comparable](in map[K]V) map[V]K {
	out := make(map[V]K, len(in))
	for k, v := range in {
		out[v] = k
	}
	return out
}
