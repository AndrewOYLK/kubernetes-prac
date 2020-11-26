package main

import (
	"fmt"
	"net/http"
	"samples/projects/CloudRestaurant/controller"
	"samples/projects/CloudRestaurant/tool"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 执行读取自定义配置文件内容
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error()) // 让程序抛出一个恐慌（中断程序，并提示）
	}

	// 执行ORM实例化
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		//logger.Error(err.Error())
		fmt.Println(err.Error())
		return
	}

	// 执行实例化Redis
	tool.InitRedisStore()

	app := gin.Default() // 声明Gin引擎

	app.Use(Cors()) // 设置全局跨域访问（必须在路由注册之前）
	tool.InitSession(app)

	registerRouter(app) // 路由注册总入口

	app.Run(cfg.AppHost + ":" + cfg.AppPort) // 运行
}

// 各模块路由注册
func registerRouter(router *gin.Engine) {
	// hello处理模块
	// 方法1:
	//var hello controller.HelloController
	//hello.Router(app)
	// 方法2:
	new(controller.HelloController).Router(router)        // hello模块
	new(controller.MemberController).Router(router)       // member模块
	new(controller.FoodCategoryController).Router(router) // FoodCategory模块
	new(controller.ShopController).Router(router)         // Shop模块
	new(controller.GoodController).Router(router)         // Good模块
}

// 跨域访问: Cross Origin Resource share
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("跨域请求中间件处理")
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //// 设置返回格式是json
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}

		//处理请求
		context.Next()
	}
}
