package slice

// Union returns the union of all slices.
// 计算多个切片的并集，只支持 comparable 类型的切片元素。
// 返回的新切片，已去重且顺序不确定。
func Union[T comparable](slices ...[]T) []T {
	unionMap := make(map[T]struct{})
	for _, s := range slices {
		for _, v := range s {
			unionMap[v] = struct{}{}
		}
	}
	return mapKeyToSlice(unionMap)
}

// UnionByEqFunc returns the union of all slices by equalFunc.
// 计算多个切片的并集，支持任意类型的切片元素，由 equal 函数判断两个元素是否相等。。
// 返回的新切片已去重。
func UnionByEqFunc[T any](equal equalFunc[T], slices ...[]T) []T {
	totalElementNum := 0
	for _, s := range slices {
		totalElementNum += len(s)
	}
	result := make([]T, 0, totalElementNum)
	for _, s := range slices {
		result = append(result, s...)
	}
	return DeduplicateByEqFunc(result, equal)
}
