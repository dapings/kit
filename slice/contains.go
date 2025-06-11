package slice

// Contains checks if the given slice contains the given element.
// If the given element is found in the slice, returns true, otherwise returns false.
// 检查给定的 slice 是否包含给定的元素。如果给定的元素在 slice 中找到了，返回 true，否则返回 false。
func Contains[T comparable](data []T, dst T) bool {
	return ContainsByFunc(data, dst, func(src, dst T) bool {
		return src == dst
	})
}

// ContainsByFunc checks if the given slice contains the given element using the provided equality function.
// If the given element is found in the slice, returns true, otherwise returns false.
// 通过提供的相等函数，检查给定的 slice 中是否包含指定元素。如果找到了指定的元素，返回 true，否则返回 false。
func ContainsByFunc[T any](data []T, dst T, eqFunc equalFunc[T]) bool {
	if eqFunc == nil {
		return false
	}
	for _, src := range data {
		if eqFunc(src, dst) {
			return true
		}
	}
	return false
}
