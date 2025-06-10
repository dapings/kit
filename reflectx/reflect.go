package reflectx

import "reflect"

// IsPtr returns true if the value is a pointer.
func IsPtr(value any) bool {
	valueOf := reflect.ValueOf(value)
	return valueOf.Kind() == reflect.Ptr
}
