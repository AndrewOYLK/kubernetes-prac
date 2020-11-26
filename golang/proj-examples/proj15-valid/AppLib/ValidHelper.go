package AppLib

import (
	"fmt"
	"reflect"
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

func ValidErrMsg(obj interface{}, err error) error {
	/*
		只要有一个出现验证错误就返回，直接抛出，而不是全部验证完毕，把整个信息给用户
		这种做法为了验证的性能
	*/

	getObj := reflect.TypeOf(obj) // 重点，拿到了具体的动态类型

	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok { // 这里执行了一个断言
			for _, e := range errs {
				fmt.Println(e.Field()) // 输出：Usertags[4]
				if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
					// fmt.Println(f.Tag.Lookup("vmsg"))
					if value, ok := f.Tag.Lookup("vmsg"); ok {
						// 如果有vmsg
						return fmt.Errorf("%s", value)
					} else {
						// 没有vmsg，返回原始内容
						// return fmt.Errorf("%s", e.Value)
						return fmt.Errorf("%s", e)
					}
				} else {
					return fmt.Errorf("%s", e)
				}
			}
		}
	}
	return nil
}

func GetValidMsg(obj interface{}, field string) {
	fmt.Println(obj)
	fmt.Println(field)
	// obj interface{} --> user实例
	// field string --> 验证错误的字段

	// 1. 通过反射获取类型，并拿到对应field的tag
	getObj := reflect.TypeOf(obj)

	// v, _ := getObj.Elem().FieldByName(field)
	// fmt.Printf("%#v\n", v.Tag)
	// 2. 获取到所有的属性字段
	if f, exist := getObj.Elem().FieldByName(field); exist {
		// fmt.Println(f.Tag) // 输出: validate:"required,min=10,max=18" vmsg="用户密码必须6位以上"
		fmt.Println(f.Tag.Get("vmsg"))
	}
}

// 封装个函数，专门用来处理正则类型的tag
func AddRegexTag(tagName string, patter string, v *validator.Validate) error {
	return v.RegisterValidation(tagName, func(fl validator.FieldLevel) bool {
		m, _ := regexp.MatchString(patter, fl.Field().String())
		return m
	}, false) // 该false/true，标识该字段是否必填
}
