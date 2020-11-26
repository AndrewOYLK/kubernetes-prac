package service

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"math/rand"
	"samples/projects/CloudRestaurant/dao"
	"samples/projects/CloudRestaurant/model"
	"samples/projects/CloudRestaurant/param"
	"samples/projects/CloudRestaurant/tool"
	"strconv"

	//"github.com/goes/logger"
	"fmt"
	"time"
)

/*
	这个结构体用于完成短信发送的服务（功能）
	这个服务当中会具体的完成一个发送验证码的过程
*/
type MemberService struct {
}

func (ms MemberService) GetUserInfo(userId string) *model.Member {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil
	}
	memberDao := dao.MemberDao{tool.DbEngine}
	return memberDao.QueryMemberById(int64(id))
}

func (ms MemberService) UploadAvatar(userId int64, fileName string) string {
	memberDao := dao.MemberDao{tool.DbEngine}
	result := memberDao.UpdateMember(userId, fileName)
	if result == 0 {
		return ""
	}
	return fileName
}

// 用户名 + 密码 登陆
func (ms MemberService) Login(name string, password string) *model.Member {
	// 1. 使用用户名 + 密码 查询用户信息， 如果存在，直接返回
	md := dao.MemberDao{Orm: tool.DbEngine}
	member := md.Query(name, password)
	if member.Id != 0 {
		return member
	}
	// 2. 用户信息不存在，作为新用户来插入数据（插入新的用户数据）
	user := model.Member{}
	user.UserName = name
	user.Password = tool.EncoderSha256(password) // 哈希计算
	user.RegisterTime = time.Now().Unix()
	result := md.InsertMember(user)
	user.Id = result
	return &user
}

// 用户手机号 + 验证码的登陆逻辑
func (ms MemberService) SmsLogin(loginparam param.SmsLogin) *model.Member {
	// 1. 获取到手机号和验证码
	// 2. 验证手机号+验证码是否正确
	md := dao.MemberDao{Orm: tool.DbEngine}
	sms := md.ValidateSmsCode(loginparam.Phone, loginparam.Code)
	if sms.Id == 0 {
		return nil
	}
	// 3. 证据手机号在member表中查询记录
	member := md.QueryByPhone(loginparam.Phone)
	if member.Id != 0 {
		return member
	}
	// 4. 如果member表中无记录，则创建一个member记录，并保存
	user := model.Member{} // 新实例化的user对象
	user.UserName = loginparam.Phone
	user.Mobile = loginparam.Phone
	user.RegisterTime = time.Now().Unix()
	user.Id = md.InsertMember(user)

	// 5. 返回用户对象
	return &user
}

func (ms MemberService) SendCode(phone string) bool {
	// 1. 产生一个验证码（产生一个四位数的纯数字验证码）
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	fmt.Println("随机码: ", code)
	// 2. 调用阿里云sdk完成发送
	config := tool.GetConfig().Sms
	fmt.Println(tool.GetConfig().Sms)

	// 使用阿里云短信服务的SDK
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	if err != nil {
		//logger.Error(err.Error())
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	// 3. 接收返回结果，并判断发送状态
	response, err := client.SendSms(request)
	if err != nil {
		//logger.Error(err.Error())
		return false
	}
	if response.Code == "OK" {
		// （数据库）将验证码保存到数据库中
		smsCode := model.SmsCode{Phone: phone, Code: code, BizId: response.BizId, CreateTime: time.Now().Unix()}
		memberDao := dao.MemberDao{tool.DbEngine}
		result := memberDao.InsertCode(smsCode)
		return result > 0
	}
	fmt.Printf("response is %#v\n", response)
	return false
}
