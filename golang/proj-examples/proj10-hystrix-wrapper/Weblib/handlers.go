package Weblib

import (
	"context"
	"proj10-hystrix-wrapper/Services"

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
