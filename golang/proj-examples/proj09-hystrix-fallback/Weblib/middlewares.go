package Weblib

import (
	"proj09-hystrix-fallback/Services"

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
