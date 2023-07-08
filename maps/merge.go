package maps

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
