package Weblib

import (
	"context"
	"proj07-grpc-client/Services"

	"github.com/gin-gonic/gin"
)

func GetProdsList(ctx *gin.Context) {
	prodService := ctx.Keys["prodService"].(Services.ProdService) // 没有类型，所以断言成Services.ProdSercice

	var prodReq Services.ProdsRequest
	err := ctx.Bind(&prodReq)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status": err.Error(),
		})
	}

	prodRes, _ := prodService.GetProdsList(context.Background(), &prodReq)
	ctx.JSON(200, gin.H{
		"data": prodRes.Data,
	})
}
