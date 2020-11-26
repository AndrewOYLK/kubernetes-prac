package tool

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// 初始化session操作
func InitSession(engine *gin.Engine) {
	config := GetConfig().RedisConfig
	store, err := redis.NewStore(
		10,
		"tcp",
		config.Addr+":"+config.Port,
		config.Password,
		[]byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
	}
	engine.Use(sessions.Sessions("mysession", store))
}

// Set操作
func SetSess(ctx *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(ctx)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

// Get操作
func GetSess(ctx *gin.Context, key interface{}) interface{} {
	session := sessions.Default(ctx)
	return session.Get(key)
}
