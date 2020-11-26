package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	fmt.Println("xxxxx: ", engine)
	// 设置加载静态文件（设置html目录）
	engine.LoadHTMLGlob("./html/*")
	engine.Static("/img", "./img")

	engine.GET("/html", func(context *gin.Context) {
		fullPath := "请求路径: " + context.FullPath()
		fmt.Println(fullPath)

		context.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"fullPath": fullPath,
				"title": "我来学gin啦！",
			})
	})

	engine.Run()
}