package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		// 以Form URL encoded形式提交上来
		var register Register
		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(register.UserName)
		fmt.Println(register.Phone)

		context.Writer.Write([]byte(register.UserName + " Register "))
	})

	engine.Run()
}

type Register struct {
	UserName string `form:"name"`
	Phone    string `form:"phone"`
	PassWord string `form:"password"`
}
