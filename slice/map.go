package slice

// IndexStructsByKey 将给定的结构体slice转换为map，其中key为结构体的某个字段的值，value为结构体本身。
func IndexStructsByKey[T any, K comparable](data []T, keyExtractFunc func(T) K) map[K]T {
	result := make(map[K]T, len(data))
	for _, item := range data {
		result[keyExtractFunc(item)] = item
	}
	return result
}

// CombineAndDeduplicateNestedSlicesByEqFunc 从任何类型的slice中提取并组合嵌套的slice，然后使用自定义的比较函数去重。
func CombineAndDeduplicateNestedSlicesByEqFunc[Src, Dst any](src []Src, extractFunc func(idx int, s Src) []Dst, eqFunc equalFunc[Dst]) []Dst {
	result := make([]Dst, 0, len(src))
	for i, s := range src {
		result = append(result, extractFunc(i, s)...)
	}
	return DeduplicateByEqFunc(result, eqFunc)
}

// CombineAndDeduplicateNestedSlices 从任何类型的slice中提取并组合嵌套的slice，并去重。
func CombineAndDeduplicateNestedSlices[Src any, Dst comparable](src []Src, extractFunc func(idx int, s Src) []Dst) []Dst {
	result := make([]Dst, 0, len(src))
	for i, s := range src {
		result = append(result, extractFunc(i, s)...)
	}
	return Deduplicate(result)
}

// CombineNestedSlices 从任何类型的slice中提取并组合嵌套的slice。
func CombineNestedSlices[Src, Dst any](src []Src, extractFunc func(idx int, s Src) []Dst) []Dst {
	result := make([]Dst, 0, len(src))
	for i, s := range src {
		result = append(result, extractFunc(i, s)...)
	}
	return result
}

// FilterMap 将给定的slice转换为一个新的，其中每个元素都是通过给定的函数fn转换得到的。
// 如果fn返回的bool为false，则元素将被忽略。
func FilterMap[Src, Dst any](src []Src, fn func(idx int, s Src) (Dst, bool)) []Dst {
	dst := make([]Dst, 0, len(src))
	for i, s := range src {
		if item, ok := fn(i, s); ok {
			dst = append(dst, item)
		}
	}
	return dst
}

// Map applies a function to each element of a slice, and returns a new slice with the results.
func Map[Src, Dst any](src []Src, fn func(idx int, s Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	for i, s := range src {
		dst[i] = fn(i, s)
	}
	return dst
}

// toMap converts a slice of T to a map of T.
// map中键为元素，值为 struct{}{}（一个空结构体）。
// 使用 struct{}{} 可以将 map 中的内存使用降至最小，因为它不会分配任何空间。
func toMap[T comparable](data []T) map[T]struct{} {
	result := make(map[T]struct{}, len(data))
	for _, d := range data {
		// remove duplicate elements from the given slice.
		result[d] = struct{}{}
	}
	return result
}

// mapKeyToSlice 将给定的 map 的键转换为切片。
func mapKeyToSlice[T comparable](data map[T]struct{}) []T {
	result := make([]T, 0, len(data))
	for k := range data {
		result = append(result, k)
	}
	return result
}
