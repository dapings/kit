package slice

// equalFunc is a function type used to compare whether two values of a given type are equal.
// If the two values are equal, return true; otherwise, return false.
// 用于比较两个给定类型的值是否相等。如果两个值相等，返回true；否则，返回false。
type equalFunc[T any] func(src, dst T) bool

// filterFunc is a function type used to determine whether an element of a given type meets a specific condition.
// If the element to be judged meets the condition, return true; otherwise, return false.
// 用于判断给定类型的元素是否符合特定条件。如果待判断元素符合条件，返回true；否则，返回false。
type filterFunc[T any] func(idx int, item T) bool
