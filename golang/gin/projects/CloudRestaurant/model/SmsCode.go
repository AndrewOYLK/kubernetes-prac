package model

/*
	通过xorm tag的方式，来指定字段在数据库表当中的约束
*/
type SmsCode struct {
	Id         int64  `xorm:"pk autoincr" json:"id"`
	Phone      string `xorm:"varchar(11)" json:"phone"`
	BizId      string `xorm:"varchar(64)" json:"biz_id"`
	Code       string `xorm:"varchar(6)" json:"code"`
	CreateTime int64  `xorm:"bigint" json:"create_time"`
}
