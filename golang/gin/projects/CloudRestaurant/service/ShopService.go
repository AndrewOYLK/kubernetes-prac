package service

import (
	"samples/projects/CloudRestaurant/dao"
	"samples/projects/CloudRestaurant/model"
	"strconv"
)

type ShopService struct {
}

/**
 * 根据shop id获取商家的服务
 */
func (shopService *ShopService) GetService(shopId int64) []model.Service {
	shopDao := dao.NewShopDao()
	return shopDao.QueryServiceByShopId(shopId)
}

/*
 * 根据关键词查询对应的商家信息
 */
func (shopService *ShopService) SearchShops(long, lati, keyword string) []model.Shop {
	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lati, 10)
	if err != nil {
		return nil
	}

	shopDao := dao.NewShopDao()
	return shopDao.QueryShops(longitude, latitude, "")
}

/*
 * 查询商铺列表数据
 */
func (shopService *ShopService) ShopList(long, lati string) []model.Shop {
	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lati, 10)
	if err != nil {
		return nil
	}

	shopDao := dao.NewShopDao()
	return shopDao.QueryShops(longitude, latitude, "")
}
