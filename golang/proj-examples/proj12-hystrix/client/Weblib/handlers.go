package Weblib

import (
	"context"
	"fmt"
	"proj12-hystrix/client/Services"

	"github.com/gin-gonic/gin"
)

// gin的方法部分
func GetProdsList(ctx *gin.Context) {
	prodService := ctx.Keys["prodService"].(Services.ProdService) // 没有类型，所以断言成Services.ProdSercice

	var prodReq Services.ProdsRequest
	err := ctx.Bind(&prodReq)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status": err.Error(),
		})
	}

	// 一般熔断处理在这里开始
	prodRes, err := prodService.GetProdsList(context.Background(), &prodReq)

	// 因为有降级方法存在，所以这里就注释掉了，直接调用返回就好
	// if err != nil {
	// 	ctx.JSON(500, gin.H{
	// 		"status": err.Error(),
	// 	})
	// }
	ctx.JSON(200, gin.H{
		"data": prodRes.Data,
	})
}

// 不同的错误可以抛出不同的JSON结果
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// 商品详细
func GetProdDetail(ctx *gin.Context) {
	fmt.Println("================")
	fmt.Println(ctx.Request.UserAgent()) // ctx实例重有个Request属性，它是http.Request的实例，那么这个实例有许多可以使用的方法

	var prodReq Services.ProdsRequest

	PanicIfError(ctx.BindUri(&prodReq)) // 绑定传参值
	// PanicIfError(ctx.BindJSON(&prodReq))

	prodService := ctx.Keys["prodService"].(Services.ProdService)

	resp, _ := prodService.GetProdDetail(context.Background(), &prodReq)

	ctx.JSON(200, gin.H{"data": resp.Data})
}
