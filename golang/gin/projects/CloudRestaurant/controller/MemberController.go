package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"samples/projects/CloudRestaurant/model"
	"samples/projects/CloudRestaurant/param"
	"samples/projects/CloudRestaurant/service"
	"samples/projects/CloudRestaurant/tool"
	"strconv"
	"time"
)

/*
	以下结构体就提供了:
		用户登陆、找回密码、发送收集验证码等等相关内容操作的方法
*/
type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)
	engine.POST("/api/login_sms", mc.smsLogin)
	engine.GET("/api/captcha", mc.captcha)
	engine.POST("/api/verifycha", mc.verifyCaptcha) // 这只是一个方便测试的接口（实际上不会调用该接口）
	engine.POST("/api/login_pwd", mc.nameLogin)
	engine.POST("/api/upload/avator", mc.uploadAvator)
	engine.GET("api/userinfo", mc.userInfo) // 用户信息查询
}

/*
 * 对用户的信息进行查询
 */
func (mc *MemberController) userInfo(context *gin.Context) {
	// 获取cookie信息
	cookie, err := tool.CookieAuth(context)
	if err != nil {
		context.Abort()
		tool.Failed(context, "还未登陆，请先登陆")
		return
	}

	// 根据cookie信息获取用户信息
	memberService := service.MemberService{}
	member := memberService.GetUserInfo(cookie.Value)
	if member != nil {
		//tool.Success(context, member)
		// 验证后，返回需要的信息
		tool.Success(context, map[string]interface{}{
			"id":            member.Id,
			"user_name":     member.UserName,
			"mobile":        member.Mobile,
			"register_time": member.RegisterTime,
			"avatar":        member.Avatar,
			"balanc":        member.Balance,
			"city":          member.City,
		})
		return
	}
	tool.Failed(context, "获取用户信息失败")
}

// 头像上传
func (mc *MemberController) uploadAvator(context *gin.Context) {
	// 1. 解析上传的参数: file、user_id
	userId := context.PostForm("user_id")
	fmt.Println(userId)
	file, err := context.FormFile("avatar")
	if err != nil || userId == "" {
		tool.Failed(context, "参数解析失败")
		return
	}
	// 2. 判断user_id对应的用户是否已经登陆
	sess := tool.GetSess(context, "user_"+userId)
	if sess == nil {
		tool.Failed(context, "参数不合法")
		return
	}
	var member model.Member
	json.Unmarshal(sess.([]byte), &member) // 解析，使用类型进行转换interface -> []byte
	// 3. file保存到本地
	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	if err := context.SaveUploadedFile(file, fileName); err != nil {
		tool.Failed(context, "头像更新失败")
		return
	}

	// 3.1 将文件上传到fastDFS系统（新增）
	fileId := tool.UploadFile(fileName)
	if fileId != "" {
		os.Remove(fileName) // 已上传到FastDFS服务器上，就把本地的文件删除
		// 4. 将保存后的文件本地路径，保存到用户表中的头像字段
		memberService := service.MemberService{}
		path := memberService.UploadAvatar(member.Id, fileId)
		if path != "" {
			tool.Success(context, tool.FileServerAddr()+"/"+path)
			return
		}
	}

	// http://localhost:8080/static/.../xx.png
	// 4. 将保存后的文件本地路径，保存到用户表中的头像字段
	//memberService := service.MemberService{}
	//path := memberService.UploadAvatar(member.Id, fileName[1:])
	//if path != "" {
	//	tool.Success(context, "http://localhost:8080" + path)
	//	return
	//}
	// 5. 返回结果
	tool.Failed(context, "上传失败")
}

// 用户名+密码、验证码登陆
func (mc *MemberController) nameLogin(context *gin.Context) {
	// 1. 解析用户登陆传递参数
	var loginParam param.LoginParm
	if err := tool.Decode(context.Request.Body, &loginParam); err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	// 2. 验证验证码
	validate := tool.VarifyCaptcha(loginParam.Id, loginParam.Value)
	if !validate {
		tool.Failed(context, "验证码不正确，请重新验证")
		return
	}
	// 3. 登陆
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		// 将用户信息保存到session当中
		sess, _ := json.Marshal(member) // member对象进行序列化
		err := tool.SetSess(context, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "登陆失败")
			return
		}
		tool.Success(context, &member)
		return
	}
	tool.Failed(context, "登陆失败")
}

func (mc *MemberController) captcha(context *gin.Context) {
	// todo 生成验证码，并返回客户端
	tool.GenerateCaptcha(context)
}

func (mc *MemberController) verifyCaptcha(context *gin.Context) {
	// 解析客户端传递的参数
	var captcha tool.CaptchaResult
	if err := tool.Decode(context.Request.Body, &captcha); err != nil {
		tool.Failed(context, "参数解析失败")
	}
	if tool.VarifyCaptcha(captcha.Id, captcha.VerifyValue) {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证失败")
	}
	tool.Success(context, "参数解析成功")
}

// 发送验证码
// http://localhost:8090/api/sendcode?phone=13631263246
func (ms *MemberController) sendSmsCode(context *gin.Context) {
	// 获取phone值
	phone, exist := context.GetQuery("phone")
	if !exist {
		tool.Failed(context, "参数解析失败")
		//context.JSON(200, map[string]interface{}{
		//	"code": "0",
		//	"msg":  "参数解析失败",
		//})
		return
	}

	// 调用短信发送功能
	service := service.MemberService{}
	isSend := service.SendCode(phone)
	if isSend {
		tool.Success(context, "发送成功")
		//context.JSON(200, map[string]interface{}{
		//	"code": "1",
		//	"msg":  "发送成功",
		//})
		return
	}
	tool.Failed(context, "发送失败")
	//context.JSON(200, map[string]interface{}{
	//	"code": 0,
	//	"msg":  "发送失败",
	//})
}

// （手机号）短信验证登陆方法
func (ms *MemberController) smsLogin(context *gin.Context) {
	var smsLoginParam param.SmsLogin
	fmt.Println("Body: ", context.Request.Body)
	if err := tool.Decode(context.Request.Body, &smsLoginParam); err != nil { // 获取请求参数(Body包含请求参数)
		tool.Failed(context, "参数解析失败")
		//context.JSON(200, map[string]interface{}{
		//	"code": 0,
		//	"msg": "参数解析失败",
		//})
		return
	}

	// 完成手机+验证码登陆
	us := service.MemberService{}
	member := us.SmsLogin(smsLoginParam)
	if member != nil {
		sess, _ := json.Marshal(member)
		err := tool.SetSess(context, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "登陆失败")
			return
		}
		// 登陆成功
		// Cookies设置
		context.SetCookie(
			"cookie_user",
			strconv.Itoa(int(member.Id)), //time.Now().String(),
			10*60,
			"/",
			"localhost",
			true,
			true)
		tool.Success(context, member)
		return
	}
	tool.Failed(context, "登陆失败")
}
