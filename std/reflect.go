package std

import "reflect"

func deepFields(interfaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < interfaceType.NumField(); i++ {
		v := interfaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, deepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}

// 仅对结构体中相同名称、相同类型的字段进行拷贝
func copyStruct(srcPtr interface{}, dstPtr interface{}) {
	srcVal := reflect.ValueOf(srcPtr)
	srcType := reflect.TypeOf(srcPtr)
	dstVal := reflect.ValueOf(dstPtr)
	dstType := reflect.TypeOf(dstPtr)

	// type kind must ptr
	if srcType.Kind() != reflect.Ptr || dstType.Kind() != reflect.Ptr {
		return
	}

	// value should not be nil
	if srcVal.IsNil() || dstVal.IsNil() {
		return
	}

	srcFields := deepFields(srcVal.Elem().Type())
	for _, v := range srcFields {
		// skip embedded field
		if v.Anonymous {
			continue
		}

		dst := dstVal.FieldByName(v.Name)
		src := srcVal.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
		if src.Type() != dst.Type() {
			// failed to copy of v.Name, src.Type, dst.Type
		}
	}
}

func getStructFieldValue(src interface{}, fieldName string) interface{} {
	val := reflect.ValueOf(src)

	// if val is ptr, pointer the elem.
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// ensure src is struct and not nil.
	if val.Kind() != reflect.Struct {
		return nil
	}

	fieldVal := val.FieldByName(fieldName)
	if fieldVal.IsValid() && fieldVal.CanInterface() {
		return fieldVal.Interface()
	}

	return nil
}
