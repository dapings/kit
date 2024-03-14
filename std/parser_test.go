package std

import (
	"reflect"
	"testing"
)

func TestConvertBasicType(t *testing.T) {
	testCases := []struct {
		name         string
		value        string
		v            any
		shouldHasErr bool
	}{
		{
			name:         "bool",
			value:        "1",
			v:            true,
			shouldHasErr: false,
		},
		{
			name:         "int",
			value:        "1",
			v:            1,
			shouldHasErr: false,
		},
	}

	for i, tt := range testCases {
		_ = i
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// reflect.Value.Elem() 仅适用于接口和指针类型，如果在非接口或非指针类型上调用Elem()方法，就会引发panic
			// 获取interface{}接口变量内部包含的值，或指针所指向的实际值
			// 确保利用reflect包时，对类型和方法调用时的逻辑条件有透彻的理解，一旦逻辑有误就可能引发panic。
			// 在实际情况下，对于反射的操作很少直接用于基本类型，而是用于更复杂的类型处理。
			obj := reflect.ValueOf(&tt.v).Elem()
			err := ConvertBasicType(tt.value, obj.Type(), obj)
			if err != nil {
				t.Errorf("error expected %v, but got %v", tt.shouldHasErr, err)
				return
			}
		})
	}
}
