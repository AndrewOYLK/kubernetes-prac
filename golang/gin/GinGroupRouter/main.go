package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	
	// (user)多路由分组 - 多模块开发
	routerGroup := engine.Group("/user")
	routerGroup.POST("/register", registerHandle)
	routerGroup.POST("/login", loginHandle)
	routerGroup.DELETE("/:id", deleteHandle)
	
	engine.Run()
}

type Register struct {
	Username string `form:username`
	Password string	`form:password`
}

func registerHandle(context *gin.Context) {
	fullPath := "用户注册: " + context.FullPath()
	fmt.Println(fullPath)

	var register Register
	if err := context.BindJSON(&register); err != nil {
		log.Fatal(err.Error())
	}
	context.Writer.WriteString(fullPath + "-" + register.Username)
}

func loginHandle(context *gin.Context) {
	fullPath := "用户登陆: " + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}

func deleteHandle(context *gin.Context) {
	userID := context.Param("id")
	fullPath := "用户(" + userID + ")删除: " + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}