package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	// gin部分可以封装到外部文件完成，涉及Gin的验证、自定义验证器、日志等等
	ginRouter := gin.Default()

	ginRouter.Handle("GET", "/user", func(ctx *gin.Context) {
		ctx.String(200, "网络请求1")
	})

	ginRouter.Handle("GET", "/news", func(ctx *gin.Context) {
		ctx.String(200, "网络请求2")
	})

	// gin整合进micro server
	server := web.NewService(
		web.Address(":8081"),
		/*
			扩展使用第三方web框架的web服务能力
			接口的作用！gin的engine类型实现了http.Handler的接口，所以可以把gin传进去
		*/
		web.Handler(ginRouter),
	)

	server.Run()
}
