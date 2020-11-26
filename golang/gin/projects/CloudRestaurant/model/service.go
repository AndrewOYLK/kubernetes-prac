package model

/**
 * 商户对应的功能服务基础表
 */
type Service struct {
	Id          int64  `xorm:"pk autoincr" json:"id"`         // Id
	Name        string `xorm:"varchar(20)" json"name"`        // 服务名称
	Description string `xorm:"varchar(30)" json"description"` // 服务描述
	IconName    string `xorm:"varchar(3)" json:"icon_name"`
	IconColor   string `xorm:"varchar(6)" json:"icon_color"`
}
