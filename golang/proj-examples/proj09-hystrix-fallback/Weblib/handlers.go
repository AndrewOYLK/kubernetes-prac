package Weblib

import (
	"context"
	"proj09-hystrix-fallback/Services"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

func defaultProds() (*Services.ProdListResponse, error) {
	// 降级方法：尽可能简单，不容易出错，尽量不要有异常
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 4; i++ {
		models = append(models, newProd(200+i, "prodname"+strconv.Itoa(100+int(i))))
	}
	res := &Services.ProdListResponse{}
	res.Data = models
	return res, nil
}

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

	// 同步方法示例 - 熔断器hystrix的使用
	// 以下都是熔断代码改造，包括降级处理

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
		// 关键部分（耗时地方）
		prodRes, err = prodService.GetProdsList(context.Background(), &prodReq)
		return err
	}, func(e error) error {
		// 降级方法
		// e：有两种情况
		// 1. 调用grpc出错
		// 2. timeout超时
		prodRes, err = defaultProds()
		return err
	})

	if err != nil {
		ctx.JSON(500, gin.H{
			"status": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{
		"data": prodRes.Data,
	})
}
