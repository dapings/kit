package slice

// Reverse reverses the slice.
func Reverse[T any](data []T) []T {
	result := make([]T, 0, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return result
}

// ReverseInplace reverses the slice in place.
func ReverseInplace[T any](data []T) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
