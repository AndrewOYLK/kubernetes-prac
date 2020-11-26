package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	/* 应用于全部接口 */
	//engine.Use(RequestInfos())

	/* 应用于单接口 */
	engine.GET("/query", RequestInfos(), func(context *gin.Context) {
		fmt.Println("=================")
		context.JSON(404, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})

	engine.GET("/hello", RequestInfos(), func(context *gin.Context) {
		// todo
	})

	engine.Run(":9001")
}

// 自定义中间件
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 第一部分（执行业务逻辑之前）
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("请求路径: ", path)
		fmt.Println("请求方法: ", method)
		fmt.Println("状态码: ", context.Writer.Status())

		/*
			context.Next()函数可以将中间件代码的执行顺序一分为二，
			Next函数调用之前在处理请求之前，当程序执行到context.Next时，
			会中断向下执行，转而先去执行具体的业务逻辑，执行完业务逻辑处理函数之后，
			程序会再次回到context.Next处，继续执行中间件后续的代码
		*/
		context.Next()

		// 第二部分（执行完业务逻辑之后）
		fmt.Println("状态码: ", context.Writer.Status())
	}
}
