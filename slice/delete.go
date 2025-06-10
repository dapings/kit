package slice

import "github.com/dapings/kit/internal/errors"

// DeleteByIndex deletes a specific index of a given type in a slice.
func DeleteByIndex[T any](data []T, index int) ([]T, error) {
	length := len(data)
	if index < 0 || index >= length {
		return nil, errors.NewIndexOutOfRange(length, index)
	}

	return DeleteByFilterFunc[T](data, func(idx int, item T) bool {
		return idx == index
	}), nil
}

// DeleteByItem deletes the specified elements in a slice of a given type that are equal to a specific item.
// Returns a new slice containing only the elements that were not equal to the specified item.
func DeleteByItem[T comparable](data []T, dstItem T) []T {
	return DeleteByFilterFunc[T](data, func(idx int, srcItem T) bool {
		return srcItem == dstItem
	})
}

// DeleteByFilterFunc deletes elements in a slice of a given type that meet a specific condition by a filter function.
// Returns a new slice containing only the elements that were ont filtered out by the filter function.
func DeleteByFilterFunc[T any](data []T, filterFunc func(idx int, item T) bool) []T {
	pos := 0
	for idx := range data {
		if !filterFunc(idx, data[idx]) {
			data[pos] = data[idx]
			pos++
		}
	}
	return data[:pos]
}
