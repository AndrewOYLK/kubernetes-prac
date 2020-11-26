package main

import (
	"proj05-prod-command/ProdService"

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
	v1Group := ginRouter.Group("/v1")
	{
		// 使用代码块方式（没有为什么，仅仅为了看的爽，眼球爽）
		v1Group.Handle("GET", "/prods", func(ctx *gin.Context) {
			ctx.JSON(200, ProdService.NewProdList(5))
		})
	}

	server := web.NewService(
		web.Name("prodservice"), // 服务名
		web.Handler(ginRouter),
		web.Registry(consulReg), // 整合注册中心
	)

	/*
		使用命令行参数来启动服务，这时候必须调用Init方法
		注意：github.com/micro/go-micro/web/service.go里的Init方法
		$ go run prodmain.go --server_address :8002
	*/
	server.Init()
	server.Run()
}
