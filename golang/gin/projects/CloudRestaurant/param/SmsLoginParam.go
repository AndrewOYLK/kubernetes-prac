package param

// 手机号 + 验证码登陆时 参数传递
// 这样做的目的主要是为了提高：解析这些参数的效率（参数不要一个一个获取）
type SmsLogin struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
