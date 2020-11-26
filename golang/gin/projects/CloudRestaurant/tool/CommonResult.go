package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 0 // 操作成功
	FAILED  int = 1 // 操作失败
)

// 普通的成功返回
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "成功",
		"data": v,
	})
}

// 普通的操作失败返回
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(200, map[string]interface{}{
		"code": FAILED,
		"msg":  "失败",
		"data": v,
	})
}
