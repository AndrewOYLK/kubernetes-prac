package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
	功能：给定任意一个复杂类型x
		1. 打印这个值对应的完整结构
		2. 同时标记每个元素的发现路径

	例子：Display("e", e) // e是一个复杂类型实例
	打印：
		Display e (eval.call):
		e.fn = "sqrt"
		e.args[0].type = eval.binary
		e.args[0].value.op = 47
		e.args[0].value.x.type = eval.Var
		e.args[0].value.x.value = "A"
		e.args[0].value.y.type = eval.Var
		e.args[0].value.y.value = "pi"
*/

func main() {

}

func Display(name string, x interface{}) {
	fmt.Println("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {

}

// 格式化所有类型（简单类型+复合类型）的值为字符串
func Any(v interface{}) string {
	return formatAtom(reflect.ValueOf(v))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "Invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}
