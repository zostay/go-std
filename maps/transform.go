package maps

func Keys[K comparable, V any](in map[K]V) []K {
	out := make([]K, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

func Values[K comparable, V any](in map[K]V) []V {
	out := make([]V, 0, len(in))
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

func KVs[K comparable, V any](in map[K]V) []any {
	out := make([]any, 0, len(in)*2)
	for k, v := range in {
		out = append(out, k, v)
	}
	return out
}
