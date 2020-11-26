package dao

import (
	"fmt"
	"samples/projects/CloudRestaurant/model"
	"samples/projects/CloudRestaurant/tool"
)

type ShopDao struct {
	*tool.Orm
}

func NewShopDao() *ShopDao {
	return &ShopDao{tool.DbEngine}
}

const DEFUALT_RANGE = 5

/**
 * 根据shop id查询商铺服务
 */
func (sd *ShopDao) QueryServiceByShopId(shopId int64) []model.Service {
	var services []model.Service

	err := sd.Table("service").Join("INNER", "shop_service", " service.id = shop_service.service_id and shop_service.shop_id = ? ", shopId).Find(&services)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return services
}

/**
 * 操作数据库查询商铺列表
 */
func (sd *ShopDao) QueryShops(long, lati float64, keyword string) []model.Shop {
	var shops []model.Shop

	if keyword == "" {
		if err := sd.Where(" longitude > ? and longitude < ? and latitude > ? and latitude < ? and status = 1 ",
			long-DEFUALT_RANGE,
			long+DEFUALT_RANGE,
			lati-DEFUALT_RANGE,
			lati+DEFUALT_RANGE,
		).Find(&shops); err != nil {
			return nil
		}
	} else {
		if err := sd.Where(" longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like ? and status = 1 ",
			long-DEFUALT_RANGE,
			long+DEFUALT_RANGE,
			lati-DEFUALT_RANGE,
			lati+DEFUALT_RANGE,
			keyword,
		).Find(&shops); err != nil {
			return nil
		}
	}

	return shops
}
