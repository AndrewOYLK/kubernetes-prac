package Weblib

import (
	"proj07-grpc-client/Services"

	"github.com/gin-gonic/gin"
)

func initMiddleware(prodService Services.ProdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["prodService"] = prodService
		context.Next()
	}
}
