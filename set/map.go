package set

var _ set[int] = &MapSet[int]{}

// MapSet a generic set implementation based on map[T]struct{} to store the set elements.
type MapSet[T comparable] struct {
	m map[T]struct{}
}

// NewMapSet creates a new MapSet instance with the given initial capacity.
func NewMapSet[T comparable](cap int) MapSet[T] {
	return MapSet[T]{
		m: make(map[T]struct{}, cap),
	}
}

// Add adds an element to the MapSet.
func (ms *MapSet[T]) Add(val T) {
	ms.m[val] = struct{}{}
}

// Remove removes an element from the MapSet.
func (ms *MapSet[T]) Remove(val T) {
	delete(ms.m, val)
}

// Contains checks if an element is present in the MapSet.
func (ms *MapSet[T]) Contains(val T) bool {
	_, ok := ms.m[val]
	return ok
}

// Clear removes all elements from the MapSet.
func (ms *MapSet[T]) Clear() {
	if len(ms.m) != 0 {
		ms.m = make(map[T]struct{})
	}
}

// IsEmpty checks if the MapSet is empty.
func (ms *MapSet[T]) IsEmpty() bool {
	return len(ms.m) == 0
}

// Size returns the number of elements in the MapSet.
func (ms *MapSet[T]) Size() int {
	return len(ms.m)
}
