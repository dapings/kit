package set

// set a generic interface that defines basic operations for set-like data structures.
type set[T comparable] interface {
	// Add adds one element to the set.
	Add(T)

	// Remove removes one element from the set.
	Remove(T)

	// Contains checks if the set contains the given element.
	Contains(T) bool

	// Clear removes all elements from the set.
	Clear()

	// IsEmpty checks if the set is empty.
	IsEmpty() bool

	// Size returns the number of elements in the set.
	Size() int
}
