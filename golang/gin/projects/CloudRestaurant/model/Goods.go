package model

/**
 * 食品结构体定义
 */
type Goods struct {
	Id          int64   `xorm:"pk autoincr" json:"id"`
	Name        string  `xorm:"varchar(12)" json:"name"`
	Description string  `xorm:"varchar(32)" json:"description"`
	Icon        string  `xorm:"varchar(255)" json:"icon"`
	SellCount   int64   `xorm:"int" json:"sell_count"` // 销售分数
	Price       float32 `xorm:"float" json:"price"`    // 销售价格
	OldPrice    float32 `xorm:"" json:"old_price"`     // 原价
	ShopId      int64   `xorm:"" json:"shop_id"`       // 商品ID
}
