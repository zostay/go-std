package maps

// Copy will copy values from the src to the dst map. For this to work, dst must
// be initialized. It will panic if dst is not initialized. Also, this does not
// clear values from dst, so existing values will be kept.
func Copy[K comparable, V any](dst, src map[K]V) {
	for k, v := range src {
		dst[k] = v
	}
}

// CopyInit will initialize dst if it is not already initialized. If it is
// already initialized, no new allocation will be performed. It then performs a
// Copy operation from src to dst.
func CopyInit[K comparable, V any](dst *map[K]V, src map[K]V) {
	if *dst == nil {
		*dst = make(map[K]V, len(src))
	}
	Copy(*dst, src)
}

// Clear will clear all values from an existing map.
func Clear[K comparable, V any](val map[K]V) {
	for k := range val {
		delete(val, k)
	}
}
