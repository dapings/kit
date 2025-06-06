package slice

// AppendDistinct appends elements of a given slice to the target slice.
// If src or items are nil, both are processed as empty slices.
// Return a new slice with distinct elements from the source and items slices.
func AppendDistinct[T comparable](src []T, items ...T) []T {
	result := make([]T, 0, len(src)+len(items))
	result = append(result, src...)
	result = append(result, items...)
	return Deduplicate(result)
}

// AddDistinctFunc appends elements of a given slice to the target slice.
// If src or items are nil, both are processed as empty slices.
// equal equalFunc[T] a function used to determine if two elements are equal, used when determining duplicate elements.
// Return a new slice with distinct elements from the source and items slices.
func AddDistinctFunc[T any](src []T, equal equalFunc[T], items ...T) []T {
	result := make([]T, 0, len(src)+len(items))
	result = append(result, src...)
	result = DeduplicateByEqFunc(result, equal)
	for _, item := range items {
		if !ContainsByFunc[T](result, item, equal) {
			result = append(result, item)
		}
	}
	return result
}
