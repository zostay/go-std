package maps

func Copy[K comparable, V any](dst, src map[K]V) {
	for k, v := range src {
		dst[k] = v
	}
}
