package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// consul注册中心部分
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.189.129:8500"), // consul地址
	)

	ginRouter := gin.Default()

	ginRouter.Handle("GET", "/user", func(ctx *gin.Context) {
		ctx.String(200, "user api")
	})

	ginRouter.Handle("GET", "/news", func(ctx *gin.Context) {
		ctx.String(200, "news api")
	})

	server := web.NewService(
		web.Name("prodservice"), // 服务名
		web.Address(":8081"),
		web.Handler(ginRouter),
		web.Registry(consulReg), // 整合注册中心
	)

	server.Run()
}
