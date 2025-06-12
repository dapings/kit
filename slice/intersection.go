package slice

// IntersectionSet returns the intersection of two slices.
// 计算两个切片的交集，只支持 comparable 类型的切片元素。
// 返回的新切片，已去重且顺序不确定。
func IntersectionSet[T comparable](s1, s2 []T) []T {
	m1 := toMap(s1)
	result := make([]T, 0, len(s1))
	for _, v2 := range s2 {
		if _, ok := m1[v2]; ok {
			result = append(result, v2)
		}
	}
	return Deduplicate[T](result)
}

// IntersectionSetByEqFunc returns the intersection of two slices by equalFunc.
// 计算两个切片的交集，由 equal 函数判断两个元素是否相等。
// 返回的新切片已去重。
func IntersectionSetByEqFunc[T any](s1, s2 []T, equal equalFunc[T]) []T {
	result := make([]T, 0, len(s1))
	for _, v1 := range s1 {
		if ContainsByFunc[T](s2, v1, equal) {
			result = append(result, v1)
		}
	}
	return DeduplicateByEqFunc(result, equal)
}
