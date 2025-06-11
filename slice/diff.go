package slice

// Diff returns the difference between two slices.
// 计算两个切片的差集，只支持 comparable 类型的切片元素。
// 返回的新切片，已去重且顺序不确定。
func Diff[T comparable](s1, s2 []T) []T {
	m1 := toMap[T](s1)
	for _, v := range s2 {
		delete(m1, v)
	}
	result := make([]T, 0, len(m1))
	for k := range m1 {
		result = append(result, k)
	}
	return result
}

// DiffFunc returns the difference between two slices by equalFunc.
// 计算两个切片的差集，由 equal 函数判断两个元素是否相等。
// 返回的新切片已去重。
func DiffFunc[T any](s1, s2 []T, equal equalFunc[T]) []T {
	result := make([]T, 0, len(s1))
	for _, v := range s1 {
		if !ContainsByFunc[T](s2, v, equal) {
			result = append(result, v)
		}
	}
	return DeduplicateByEqFunc(result, equal)
}
