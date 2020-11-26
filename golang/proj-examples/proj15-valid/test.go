package main

import (
	"fmt"
	"log"

	"proj15-valid/AppLib"

	"gopkg.in/go-playground/validator.v9"
)

/*
	用户名的规则（需要使用到正则表达式）
	1. 首字母必须是字母
	2. 其余只能是数字、字母或下划线
	3. 长度在6，20位之间

	[a-zA-Z]\\w{5,19}


	usertags代表是用户标签
	1. 用户注册时必须填一个
	2. 里面每一个tag的内容不能相同
	3. 最多不能超过5个
	4. 标签只能是字母数字或者中文

	^[\u4e00-\u9fa5a-zA-Z0-9]{2,4}$
*/

// 具体的“动态类型” Users
// 打tag! 可以通过反射到具体动态类型来拿到
type Users struct {
	Username string   `validate:"required,min=6,max=20" vmsg:"用户名称必须6位以上"`   //
	Userpwd  string   `validate:"required,min=10,max=18" vmsg:"用户密码必须10位以上"` // 自定义vmsg，用于不通过验证时，作为返回信息
	Testname string   `validate:"required,abc" vmsg:"用户名规则不正确"`
	Usertags []string `validate:"required,min=1,max=5,unique,dive,usertag" vmsg:"用户名标签不合法"` // dive表示右边的验证规则深入到切片内的string
	// 注意：还可以分体式验证
}

func main() {
	// 具体的“动态值” user
	// userTags := []string{"a", "b", "c", "d", "e"}
	userTags := []string{"aa", "bb", "cc", "dd", "ee"}

	user := &Users{
		Username: "xxxxxxx",
		Userpwd:  "123123123123",
		// Testname: "xxxx", // "规则不正确"
		Testname: "ac213qweq",
		Usertags: userTags,
	}

	valid := validator.New() // 定义一个验证器

	// 注册一个验证tag，也就是自定义验证tag，而且支持正则
	// 方式1：
	// _ = valid.RegisterValidation("abc", func(fl validator.FieldLevel) bool {
	// 	m, _ := regexp.MatchString("[a-zA-Z]\\w{5,19}", fl.Field().String())
	// 	return m
	// 	// fmt.Println(fl.Field().String()) // 输出"xxxx"
	// 	// return true
	// }, false) // 这里的false代表这个abc没有填值就不做验证

	// 方式2：添加自定义的正则验证tag
	_ = AppLib.AddRegexTag("abc", "[a-zA-Z]\\w{5,19}", valid)
	_ = AppLib.AddRegexTag("usertag", "^[\u4e00-\u9fa5a-zA-Z0-9]{2,4}$", valid)

	err := valid.Struct(user) // 进行验证

	// 提取错误字段的vmsg tag信息
	// 方法1
	// if err != nil {
	// 	log.Fatal(err) // 报出的信息太详细和官方
	// }

	// 方法2
	// if err != nil {
	// 	if errs, ok := err.(validator.ValidationErrors); ok { // 类型断言; 代表err是validator.ValidationErrors类型
	// 		// validator.ValidationErrors是一个[]FieldError类型的实例
	// 		// 因为执行验证返回的err，可能涉及到多个field的错误验证信息
	// 		for _, e := range errs {
	// 			// fmt.Println(e.Value()) // "123"
	// 			// fmt.Println(e.Field()) // "Userpwd"
	// 			// fmt.Println(e.Tag()) // min
	// 			AppLib.GetValidMsg(user, e.Field())
	// 		}
	// 	}
	// 	log.Fatal(err)
	// }

	// 方法3
	err = AppLib.ValidErrMsg(user, err)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("验证成功")
}
