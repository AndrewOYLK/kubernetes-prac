package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hellojson", func(context *gin.Context) {
		fullpath := "请求路径: " + context.FullPath()
		fmt.Println(fullpath)
		// map类型 - key和value形式
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg": "OK!",
			"data": fullpath,
		})
	})

	engine.GET("/jsonstruct", func(context *gin.Context) {
		fullpath := "请求路径: " + context.FullPath()
		fmt.Println(fullpath)

		// struct类型
		resp := Response{
			Code:    1,
			Message: "OK!",
			Data: fullpath,
		}
		//context.JSON(200, &resp) // 取值类型的地址
		context.JSON(400, &resp) // 取值类型的地址
	})

	engine.Run()
}

type Response struct {
	Code int
	Message string
	Data interface{}
}
