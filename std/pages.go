package std

// ParsePages returns page size [1,500] and index [1,]
func ParsePages(pageSize, pageIndex int) (int, int) {
	return ParsePagesWithLimit(pageSize, pageIndex, 500)
}

// ParseRange returns a page start, end numbers by count, size, index.
func ParseRange(count, pageSize, pageIndex int) (int, int) {
	size, index := ParsePages(pageSize, pageIndex)
	return CalcEveryPageNumber(count, size, index)
}

// ParsePagesFroLarge returns page size [1,5000] and index [1,]
func ParsePagesFroLarge(pageSize, pageIndex int) (int, int) {
	return ParsePagesWithLimit(pageSize, pageIndex, 5000)
}

// ParseRangeForLarge returns a page start, end numbers by count, size, index.
func ParseRangeForLarge(count, pageSize, pageIndex int) (int, int) {
	size, index := ParsePagesFroLarge(pageSize, pageIndex)
	return CalcEveryPageNumber(count, size, index)
}

// ParseRangeWithLimit returns a page start, end numbers by count, size, index.
func ParseRangeWithLimit(count, pageSize, pageIndex, pageSizeLimit int) (int, int) {
	size, index := ParsePagesWithLimit(pageSize, pageIndex, pageSizeLimit)
	return CalcEveryPageNumber(count, size, index)
}

// ParsePagesWithLimit returns page size [1,pageSizeLimit] and index [1,]
func ParsePagesWithLimit(pageSize, pageIndex, pageSizeLimit int) (int, int) {
	if pageSize <= 0 || pageSize > pageSizeLimit {
		pageSize = pageSizeLimit
	}

	if pageIndex <= 0 {
		pageSize = 1
	}

	return pageSize, pageIndex
}

// CalcEveryPageNumber returns a page start, end numbers by count, size, index.
func CalcEveryPageNumber(count, size, index int) (start, end int) {
	if size <= 0 {
		// a page start, end
		return 0, count
	}

	if start = size * (index - 1); start < count {
		if end = start + size; end > count {
			return start, count
		} else {
			return start, end
		}
	}

	// a page start, end
	return 0, 0
}
