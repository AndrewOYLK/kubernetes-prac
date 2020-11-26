package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * 在需要登陆才能访问的接口处理中，调用如上函数进行用户权限判断，以此来判断用户是否已经登陆
 */

const COOKIENAME = "cookie_user"
const COOKIETIMELENGTH = 10 * 60 // 10分钟

func CookieAuth(ctx *gin.Context) (*http.Cookie, error) {
	cookie, err := ctx.Request.Cookie(COOKIENAME)
	if err == nil {
		ctx.SetCookie(
			cookie.Name,
			cookie.Value,
			cookie.MaxAge,
			cookie.Path,
			cookie.Domain,
			cookie.Secure,
			cookie.HttpOnly)
	} else {
		return nil, err
	}
	return cookie, nil
}
