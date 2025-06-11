package slice

// Filter filters the slice element according to the filter function and returns a new slice.
//
// Parameters:
//   - data: A slice of type T that contains the elements to be filtered.
//   - filterFunc: A function indicating whether the element should be included in the filtered slice.
//
// Returns:
//   - A new slice of type T containing the elements that pass the filter function.
func Filter[T any](data []T, filterFunc filterFunc[T]) []T {
	result := make([]T, 0, len(data))
	for idx, item := range data {
		if filterFunc(idx, item) {
			result = append(result, item)
		}
	}
	return result
}
