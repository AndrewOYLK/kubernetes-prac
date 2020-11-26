package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	ginRouter := gin.Default()

	ginRouter.Handle("GET", "/", func(ctx *gin.Context) {
		data := make([]interface{}, 0) // 切片，长度0

		ctx.JSON(200, gin.H{
			"data": data,
		})
	})

	server := web.NewService(
		web.Address(":8000"),
		web.Handler(ginRouter),
	)

	server.Run()
}
