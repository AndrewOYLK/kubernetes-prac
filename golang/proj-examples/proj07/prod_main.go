package main

import (
	"fmt"
	"proj07/Helper"
	"proj07/ProdService"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// consul注册中心部分
	consulReg := consul.NewRegistry(
		registry.Addrs("10.1.0.13:8500"), // consul地址
	)

	ginRouter := gin.Default()
	v1Group := ginRouter.Group("/v1")
	{
		// API的输出（统一）
		v1Group.Handle("POST", "/prods", func(ctx *gin.Context) {
			fmt.Printf("POST请求提交数据形式: %s\n", ctx.ContentType())

			// 获取参数方式1
			// ctx.PostForm

			// 获取参数方式2 - 通过Bind绑定结构体
			// 使用绑定结构体，最好在结构体中做好验证方法
			// ctx.Bind
			var pr Helper.ProdsRequest
			// err := ctx.Bind(&pr) // 返回的err与验证有关
			err := ctx.BindJSON(&pr) // 返回的err与验证有关
			fmt.Println(pr)
			if err != nil || pr.Size <= 0 {
				pr = Helper.ProdsRequest{Size: 2}
			}

			ctx.JSON(
				200,
				gin.H{"data": ProdService.NewProdList(pr.Size)},
			)
		})
	}

	server := web.NewService(
		web.Name("prodservice"), // 服务名
		// web.Address(":8001"),
		web.Address(":8002"),
		web.Handler(ginRouter),
		web.Registry(consulReg), // 整合注册中心
	)
	server.Run()
}
