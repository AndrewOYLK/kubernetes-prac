package controller

import (
	"github.com/gin-gonic/gin"
	"samples/projects/CloudRestaurant/service"
	"samples/projects/CloudRestaurant/tool"
	"strconv"
)

type GoodController struct {
}

func (gc *GoodController) Router(engine *gin.Engine) {
	engine.GET("/api/foods", gc.getGoods)
}

// 获取某个商户下面所包含的食品
func (gc *GoodController) getGoods(ctx *gin.Context) {
	shopId, exist := ctx.GetQuery("shop_id")
	if !exist {
		tool.Failed(ctx, "参数解析失败，请重试")
		return
	}

	// 实例化一个goodService，并调用对应的Service方法
	id, err := strconv.Atoi(shopId)
	if err != nil {
		tool.Failed(ctx, "参数解析失败，请重试")
	}
	var goodService = service.GoodService{}
	goods := goodService.GetFoods(int64(id))
	if len(goods) == 0 {
		tool.Failed(ctx, "未查询到相关数据")
		return
	}
	tool.Success(ctx, goods)
}
