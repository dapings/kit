package std

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ConvertBasicType 根据obj的反射类型、值，将value设置给obj。
func ConvertBasicType(value string, objType reflect.Type, objValue reflect.Value) error {
	switch objType.Kind() {
	case reflect.Bool:
		if strings.ToLower(value) == "on" || strings.ToLower(value) == "1" || strings.ToLower(value) == "yes" {
			objValue.SetBool(true)
			return nil
		}
		if strings.ToLower(value) == "off" || strings.ToLower(value) == "0" || strings.ToLower(value) == "no" {
			objValue.SetBool(false)
			return nil
		}
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		objValue.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		x, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		objValue.SetInt(x)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		x, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		objValue.SetUint(x)
	case reflect.Float32, reflect.Float64:
		x, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		objValue.SetFloat(x)
	case reflect.Interface:
		objValue.Set(reflect.ValueOf(value))
	case reflect.String:
		objValue.SetString(value)
	default:
		return fmt.Errorf("%v is not supported", objType.Kind())
	}

	return nil
}
