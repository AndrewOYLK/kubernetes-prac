package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"samples/projects/CloudRestaurant/service"
	"samples/projects/CloudRestaurant/tool"
)

type ShopController struct {
}

/*
 * shop模块的路由解析
 */
func (sc *ShopController) Router(engine *gin.Engine) {
	engine.GET("/api/shops", sc.GetShopList)
	engine.GET("/api/search_shops", sc.SearchShop)
}

/*
 * 关键词搜索商铺信息
 */
func (sc *ShopController) SearchShop(ctx *gin.Context) {
	longitude := ctx.Query("longitude")
	latitude := ctx.Query("latitude")
	keyword := ctx.Query("keyword")

	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34"
		latitude = "40.34"
	}

	fmt.Println("===", keyword)
	if keyword == "" {
		tool.Failed(ctx, "查询错误，重新输入商铺名称")
		return
	}

	// 执行真实的搜索逻辑
	shopService := service.ShopService{}
	shops := shopService.SearchShops(longitude, latitude, keyword)
	if len(shops) != 0 {
		tool.Success(ctx, shops)
	}
	tool.Failed(ctx, "获取商铺列表失败")
}

/*
 * 获取商铺列表
 */
func (sc *ShopController) GetShopList(ctx *gin.Context) {
	longitude := ctx.Query("longitude")
	latitude := ctx.Query("latitude")

	fmt.Println("=================")
	fmt.Println(longitude, latitude)
	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34"
		latitude = "40.34"
	}

	var shopService service.ShopService
	shops := shopService.ShopList(longitude, latitude)
	if len(shops) == 0 {
		tool.Failed(ctx, "获取商铺列表失败")
		return
	}

	for _, shop := range shops {
		shopServices := shopService.GetService(shop.Id)
		if len(shopServices) == 0 {
			shop.Supports = nil
		} else {
			//fmt.Println("商户服务:" , shopServices)
			shop.Supports = shopServices
		}
	}
	tool.Success(ctx, shops)
}
