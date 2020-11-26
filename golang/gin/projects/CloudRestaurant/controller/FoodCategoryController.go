package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"samples/projects/CloudRestaurant/service"
	"samples/projects/CloudRestaurant/tool"
)

type FoodCategoryController struct {
}

func (fc *FoodCategoryController) Router(engine *gin.Engine) {
	engine.GET("/api/food_category", fc.foodCategory)
}

func (fc *FoodCategoryController) foodCategory(ctx *gin.Context) {
	// 调用service功能获取食品种类信息
	foodCategoryService := service.FoodCategoryService{}
	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Failed(ctx, "食品种类数据获取失败")
		return
	}

	// 格式转换（重要，主要对于图片的地址拼接）
	// imgUrl：hello.png
	for _, category := range categories {
		if category.ImageUrl != "" {
			category.ImageUrl = tool.FileServerAddr() + "/" + category.ImageUrl
			fmt.Println("==============")
			fmt.Println(category.ImageUrl)
		}
	}
	tool.Success(ctx, categories)
}
