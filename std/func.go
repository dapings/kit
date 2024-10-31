package std

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"
)

func CurrentFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	funcDesc := runtime.FuncForPC(pc)
	if ok && funcDesc != nil {
		return funcDesc.Name()
	}

	return ""
}

func ParentFuncName() string {
	pc, _, _, ok := runtime.Caller(2)
	funcDesc := runtime.FuncForPC(pc)
	if ok && funcDesc != nil {
		return funcDesc.Name()
	}

	return ""
}

// ParentFuncNameAndFileLine 获取调用者的详细信息，包括函数名、文件名及行数
func ParentFuncNameAndFileLine() string {
	// 获取调用者的PC(程序计数器)地址，skip为1，+1是要跳过当前函数本身
	pc, file, line, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		_, filename := path.Split(file)
		return fmt.Sprintf("[%s:L%d]%s", filename, line, details.Name())
	}
	return ""
}

func ParentFuncShortName() string {
	pc, _, _, ok := runtime.Caller(2)
	funcDesc := runtime.FuncForPC(pc)
	if ok && funcDesc != nil {
		name := funcDesc.Name()
		if i := strings.Index(name, "."); i >= 0 {
			name = name[i+1:]
		}

		return name
	}

	return ""
}

func SplitFuncName(fn string, separator ...rune) string {
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range separator {
			if sep == s {
				return true
			}
		}

		return false
	})

	if size := len(fields); size > 0 {
		return fields[size-1]
	}

	return ""
}

func GetFuncName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}
