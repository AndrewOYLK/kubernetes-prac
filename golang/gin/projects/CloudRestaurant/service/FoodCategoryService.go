package service

import (
	dao2 "samples/projects/CloudRestaurant/dao"
	"samples/projects/CloudRestaurant/model"
)

type FoodCategoryService struct {
}

/**
 * 获取美食种类
 */
func (fcs *FoodCategoryService) Categories() ([]model.FoodCategory, error) {
	// 数据库操作层
	foodCategoryDao := dao2.NewFoodCategoryDao()
	return foodCategoryDao.QueryCategories()
}
