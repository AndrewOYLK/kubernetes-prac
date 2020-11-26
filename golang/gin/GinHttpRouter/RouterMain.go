package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	/* 分类处理 - GET */
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		name := context.Query("name") // http://localhost:8080/hello?name=mansi
		fmt.Println(name)

		context.Writer.Write([]byte("Hello, " + name))
	})

	/* 分类处理 - POST */
	engine.POST("/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		// Form URL Encoded 前端提交方式（表单提交形式）
		username, exist := context.GetPostForm("username")
		if exist {
			fmt.Println(username)
		}
		password, exist := context.GetPostForm("password")
		if exist {
			fmt.Println(password)
		}

		context.Writer.Write([]byte("Hello, " + username))
	})

	engine.DELETE("/user/:id", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		userID := context.Param("id") // http://localhost:8080/user/123
		fmt.Println(userID)

		context.Writer.Write([]byte("删除用户ID: " + userID))
	})

	engine.Run()
}
