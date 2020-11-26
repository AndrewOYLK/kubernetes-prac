package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

// 指定配置选项参数
var driver = &base64Captcha.DriverDigit{
	Length:   3,
	Width:    115,
	Height:   70,
	MaxSkew:  0.55,
	DotCount: 30,
}

// 验证码全局对象
//var store = base64Captcha.DefaultMemStore
//c := base64Captcha.NewCaptcha(driver, store)
var c = base64Captcha.NewCaptcha(driver, &Rds) // redis后端结构体指针对象

// 生成图形验证码
func GenerateCaptcha(ctx *gin.Context) {
	// Rds已经在main函数里提前实例化了
	id, b64s, err := c.Generate()
	if err != nil {
		Failed(ctx, "生成验证码失败")
	}
	// 方法1:
	captchaResult := CaptchaResult{
		Id:         id,
		Base64Blob: b64s,
	}
	Success(ctx, map[string]interface{}{
		"captcha_result": captchaResult,
	})
	// 方法2:
	//body := map[string]interface{}{
	//	"code": 1,
	//	"data": b64s,
	//	"captchaId": id,
	//	"msg": "success",
	//}
	//if err != nil {
	//	body = map[string]interface{}{
	//		"code": 0,
	//		"msg": err.Error(),
	//	}
	//	fmt.Println(err.Error())
	//}
	////ctx.Request.Header.Set("Content-Type", "application/json; charset=utf-8")
	//ctx.JSON(200, body)
}

// 验证验证码
func VarifyCaptcha(id string, value string) bool {
	verifyResult := c.Verify(id, value, true)
	return verifyResult
	//decoder := json.NewDecoder(ctx.Request.Body)
	//var param configJsonBody
	//err := decoder.Decode(&param)
	//if err != nil {
	//	log.Println(err)
	//}
	//defer r.Body.Close()
	////verify the captcha
	//body := map[string]interface{}{"code": 0, "msg": "failed"}
	//if store.Verify(param.Id, param.VerifyValue, true) {
	//	body = map[string]interface{}{"code": 1, "msg": "ok"}
	//}
	//
	////set json response
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//
	//json.NewEncoder(w).Encode(body)
}
