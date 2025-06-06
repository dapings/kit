package slice

// Deduplicate removes duplicate elements from the given slice and returns a new slice.
// Because of map, the location of elements can change even if there are no duplicate elements.
// 去重后的新 slice，由于使用了 map，即使没有重复元素，元素位置也可能会发生改变。
func Deduplicate[T comparable](data []T) []T {
	// convert the given slice to a map, auto deduplicate.
	mp := toMap(data)
	result := make([]T, 0, len(mp))
	for item := range mp {
		result = append(result, item)
	}
	return result
}

// DeduplicateByEqFunc takes a generic slice "data" and an equivalence comparison function "equalFunc[T]", and
// returns a new slice with duplicate elements removed.
// Parameters:
// - data: a slice of any type "[]T" that needs to be deduplicated.
// - equalFunc: a function of type "equalFunc[T]" used to compare whether two elements are equal, where "T" can be any type.
//
// Return:
// - a new deduplicated slice of type "[]T".
func DeduplicateByEqFunc[T any](data []T, eqFunc equalFunc[T]) []T {
	result := make([]T, 0, len(data))
	for i, val := range data {
		// if no same element, append to result.
		if !ContainsByFunc[T](result[:i], val, eqFunc) {
			result = append(result, val)
		}
	}
	return result
}
