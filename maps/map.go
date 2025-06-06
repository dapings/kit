package maps

// Keys extracts all the keys from the given map and returns them as a slice.
// returns a slice containing all the keys in the given map.
// Note: the order of the keys is random.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values extracts all the values from the given map and returns them as a slice.
// returns a slice containing all the values in the given map.
// Note: the order of the values is random.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// KeyValues extracts all the keys and values from the given map and returns them as separate slices.
// returns a slice containing all the keys in the given map and a slice containing all the values in the given map.
// Note: the order of the keys and values is random.
func KeyValues[K comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
