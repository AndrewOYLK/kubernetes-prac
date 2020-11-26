package Weblib

import (
	"context"
	"proj09-hystrix/Services"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

// gin的方法部分
func GetProdsList(ctx *gin.Context) {
	// 业务代码 - 函数

	prodService := ctx.Keys["prodService"].(Services.ProdService) // 没有类型，所以断言成Services.ProdSercice

	var prodReq Services.ProdsRequest
	err := ctx.Bind(&prodReq)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status": err.Error(),
		})
	}

	// 同步方法示例
	// 关键部分（耗时地方）
	// 熔断代码改造
	// 1. 配置config
	configA := hystrix.CommandConfig{
		// Timeout: 5000, // 允许5秒超时
		Timeout: 1000, // 允许1秒超时时间
	}
	// 2. 配置command
	hystrix.ConfigureCommand("getprods", configA)
	// 3. 执行使用Do方法
	var prodRes *Services.ProdListResponse
	err = hystrix.Do("getprods", func() error { // 执行函数
		// 具体业务代码（rpc服务请求）
		prodRes, err = prodService.GetProdsList(context.Background(), &prodReq)
		return err
	}, nil) // nil这里是个降级方法

	if err != nil {
		ctx.JSON(500, gin.H{
			"status": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{
		"data": prodRes.Data,
	})
}
