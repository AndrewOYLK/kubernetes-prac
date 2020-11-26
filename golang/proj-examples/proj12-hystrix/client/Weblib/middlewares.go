package Weblib

import (
	"fmt"
	"proj12-hystrix/client/Services"

	"github.com/gin-gonic/gin"
)

// 中间件机制
func initMiddleware(prodService Services.ProdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["prodService"] = prodService
		context.Next()
	}
}

// 统一异常处理
func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ctx.JSON(500, gin.H{"status": fmt.Sprintf("%s", r)})
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
