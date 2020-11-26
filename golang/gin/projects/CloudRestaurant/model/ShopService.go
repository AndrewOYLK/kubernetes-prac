package model

/*
 * 定义商户-服务关系表（1:n）
 */
type ShopService struct {
	ShopId    int64 `xorm:"pk not null" json:"shop_id"`
	ServiceId int64 `xorm:"pk not null" json:"service_id"`
}
