package set

// FlipMap will return a new map with the keys and values of the input map
// flipped. That is, the keys of the input map will be the values of the output
// map and the values of the input map will be the keys of the output map.
// However, the resulting value is a Set of keys.
//
// This will only work if the value type is comparable. If values repeat, the
// keys will be collected into the value set. The ordering is indeterminate.
//
// You might also be interested in maps.Flip and maps.FlipSlice.
func FlipMap[K comparable, V comparable](in map[K]V) map[V]Set[K] {
	out := make(map[V]Set[K], len(in))
	for k, v := range in {
		outv, alreadyExists := out[v]
		if !alreadyExists {
			outv = New[K]()
		}
		outv.Insert(k)
		out[v] = outv
	}
	return out
}
