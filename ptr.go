package kit

import (
	"reflect"

	"github.com/dapings/kit/log"
)

func ToPtr[T any](v T) *T {
	return &v
}

// GetValueOrDefault returns the value of src if src is not nil, otherwise it returns the default value of T.
// Note: only the value pointed to by the first layer of pointers is taken.
func GetValueOrDefault[T any](src *T) (t T) {
	if src != nil {
		t = *src
	}

	return
}

// IsAllPointerNil 判断结构体内所有指针字段是否都为nil
func IsAllPointerNil(s any) bool {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			return false
		}
	}
	return true
}

// copyStruct 拷贝结构体
// 仅仅对相同名称相同类型的字段进行拷贝
// 为了防止拷贝出错引起bug只能在controller包内使用
func copyStruct(srcPtr interface{}, dstPtr interface{}) {
	src := reflect.ValueOf(srcPtr)
	dst := reflect.ValueOf(dstPtr)
	srcT := reflect.TypeOf(srcPtr)
	dstT := reflect.TypeOf(dstPtr)
	if srcT.Kind() != reflect.Ptr || dstT.Kind() != reflect.Ptr ||
		srcT.Elem().Kind() == reflect.Ptr || dstT.Elem().Kind() == reflect.Ptr {
		log.Error("Fatal error:type of parameters must be Ptr of value")
		return
	}
	if src.IsNil() || dst.IsNil() {
		log.Error("Fatal error:value of parameters should not be nil")
		return
	}
	srcV := src.Elem()
	dstV := dst.Elem()
	srcFields := deepFields(reflect.ValueOf(srcPtr).Elem().Type())
	for _, v := range srcFields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && src.IsNil() {
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
			log.ErrorExt(3, "copy failed of", v.Name, src.Type(), dst.Type())
		}
	}
}

func deepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, deepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}
