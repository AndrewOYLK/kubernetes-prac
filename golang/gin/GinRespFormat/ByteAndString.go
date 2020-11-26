package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hellobyte", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		fullpath := "请求路径:" + context.FullPath()
		fmt.Println(fullpath)
		context.Writer.Write([]byte(fullpath))
	})
	
	engine.GET("/hellostring", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		fullpath := "请求路径:" + context.FullPath()
		fmt.Println(fullpath)
		context.Writer.WriteString(fullpath)
	})

	engine.Run(":8081")
}
