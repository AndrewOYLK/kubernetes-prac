package model

// 用户数据结构体定义
type Member struct {
	Id           int64   `orm:"pk autoincr" json:"id"`
	UserName     string  `orm:"varchar(20)" json:"user_name"`
	Mobile       string  `orm:"varchar(11)" json:"mobile"`
	Password     string  `orm:"varchar(255)" json:"password"`
	RegisterTime int64   `orm:"bigint" json:"register_time"`
	Avatar       string  `orm:"varchar(255)" json:"avatar"`
	Balance      float64 `orm:"double" json:"balance"`
	IsActive     int8    `orm:"tinyint" json:"is_active"`
	City         string  `orm:"varchar(10)" json:"city"`
}
