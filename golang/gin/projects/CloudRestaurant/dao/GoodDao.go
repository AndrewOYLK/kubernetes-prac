package dao

import (
	"fmt"
	"samples/projects/CloudRestaurant/model"
	"samples/projects/CloudRestaurant/tool"
)

type GoodDao struct {
	*tool.Orm
}

func NewGoodDao() *GoodDao {
	return &GoodDao{tool.DbEngine}
}

// 根据商户的ID查询商户下的所拥有的所有食品数据
func (gd *GoodDao) GetFoodsByShopId(shop_id int64) ([]model.Goods, error) {
	var goods []model.Goods
	if err := gd.Where(" shop_id = ? ", shop_id).Find(&goods); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return goods, nil
}
