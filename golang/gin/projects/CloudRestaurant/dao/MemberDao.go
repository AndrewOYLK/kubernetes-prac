package dao

import (
	"fmt"
	"samples/projects/CloudRestaurant/model"
	"samples/projects/CloudRestaurant/tool"
)

type MemberDao struct {
	*tool.Orm
}

func (md *MemberDao) QueryMemberById(id int64) *model.Member {
	member := model.Member{Id: id}
	if _, err := md.Where(" id = ? ", id).Get(&member); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &member
}

// 更新member记录，头像属性
func (md *MemberDao) UpdateMember(userId int64, fileName string) int64 {
	member := model.Member{Avatar: fileName}
	result, err := md.Where(" id = ? ", userId).Update(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

// 通过用户名和密码查询
func (md *MemberDao) Query(name string, password string) *model.Member {
	var member model.Member

	if _, err := md.Where(" user_name = ? and password = ? ", name, tool.EncoderSha256(password)).Get(&member); err != nil {
		fmt.Println(err.Error())
	}
	return &member
}

// 验证手机号和验证码是否存在
func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode

	if _, err := md.Where(" phone = ? and code = ? ", phone, code).Get(&sms); err != nil {
		fmt.Println(err.Error())
	}
	return &sms
}

// 会员新用户的插入
func (md *MemberDao) InsertMember(member model.Member) int64 {
	// 为什么返回int64?因为操作数据库操作的时候，会返回一个int64的值
	result, err := md.InsertOne(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

// 函授接收值为: MemberDao
func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member

	if _, err := md.Where(" mobile = ? ", phone).Get(&member); err != nil {
		fmt.Println(err.Error())
	}
	return &member
}

func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		//logger.Error(err.Error())
		fmt.Println(err.Error())
	}
	return result
}
