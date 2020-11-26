package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	/* 通用方法 - Handle() */
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		fmt.Println("请求路径: ", context.FullPath())

		name := context.DefaultQuery("name", "none")
		fmt.Println(name)

		context.Writer.Write([]byte("Hello, " + name))
	})

	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		// 前端表单提交方式（Form URL Encoded）
		username := context.PostForm("username")
		password := context.PostForm("password")
		fmt.Println(username)
		fmt.Println(password)

		context.Writer.Write([]byte(username + " 登陆"))
	})

	/* 分类处理 - GET()、POST()、DELETE()... */
	//engine.GET("/hello", func(context *gin.Context) {
	//	fmt.Println("请求路径: ", context.FullPath())
	//	context.Writer.Write([]byte("Hello, gin\n"))
	//})

	if err := engine.Run(":8081"); err != nil {
		log.Fatal(err.Error())
	}
}
