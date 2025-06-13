package slice

// IndexOf searches the given target element in the provided slice and returns the index of the first occurrence.
// If the target element is not found, - 1 is returned.
// Note: This function requires the element type to support the == operator.
func IndexOf[T comparable](data []T, dst T) int {
	return IndexOfByFunc[T](data, dst, func(src, dst T) bool {
		return src == dst
	})
}

// IndexOfByFunc looks up the index of the first occurrence of a given target element
// in the provided slice based on the equality function.
// If the target element is found, its index in the slice is returned; Otherwise, it returns - 1.
func IndexOfByFunc[T any](data []T, dst T, equal equalFunc[T]) int {
	for idx, item := range data {
		if equal(item, dst) {
			return idx
		}
	}
	return -1
}

// LastIndexOf searches the given target element in the provided slice and returns the index of the last occurrence.
// If the target element is not found, - 1 is returned.
// Note: This function requires the element type to support the == operator.
func LastIndexOf[T comparable](data []T, dst T) int {
	return LastIndexOfByFunc[T](data, dst, func(src, dst T) bool {
		return src == dst
	})
}

// LastIndexOfByFunc looks up the index of the last occurrence of a given target element
// in the provided slice according to the equality function.
// If the target element is found, its index in the slice is returned; Otherwise, it returns - 1.
func LastIndexOfByFunc[T any](data []T, dst T, equal equalFunc[T]) int {
	for idx := len(data) - 1; idx >= 0; idx-- {
		if equal(data[idx], dst) {
			return idx
		}
	}
	return -1
}
