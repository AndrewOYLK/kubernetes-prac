package service

import (
	"fmt"
	"samples/projects/CloudRestaurant/dao"
	"samples/projects/CloudRestaurant/model"
)

type GoodService struct {
}

func NewGoodService() *GoodService {
	return &GoodService{}
}

func (gs *GoodService) GetFoods(shop_id int64) []model.Goods {
	goodDao := dao.NewGoodDao()
	goods, err := goodDao.GetFoodsByShopId(shop_id)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return goods
}
