package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"samples/projects/CloudRestaurant/tool"
)

type HelloController struct {
}

// 路由配置
func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hello.Hello)
	engine.GET("/test", hello.ForTest)
}

func (hello *HelloController) ForTest(context *gin.Context) {
	data := map[string]interface{}{
		"message": "ouyanglongkun",
	}
	sess, _ := json.Marshal(data)
	tool.SetSess(context, "demo1", sess)
	tool.SetSess(context, "demo2", sess)
	tool.SetSess(context, "demo3", sess)
	tool.SetSess(context, "demo4", sess)
	//context.Writer.Write([]byte("OK"))

	ss := tool.GetSess(context, "demo1")
	fmt.Println("=========")
	fmt.Println(ss)
	fmt.Println(json.Unmarshal(ss.([]byte), &data))
	fmt.Printf("类型: %T", ss)

	context.JSON(200, map[string]interface{}{
		"code": "1",
		"msg":  "OK",
		"data": "",
	})
}

// 业务处理方法
func (hello *HelloController) Hello(context *gin.Context) {
	fmt.Println(context.FullPath())

	// JSON数据格式返回 - 方法1
	//context.JSON(200, map[string]interface{}{
	//	"code": "1",
	//	"msg": "OK!",
	//	"data": "Hello",
	//})

	// JSON数据格式返回 - 方法2
	resp := Resp{
		Code: "10000",
		Msg:  "OK啦!",
		Data: "Gin我来啦!",
	}
	//tool.Test()
	context.JSON(200, &resp)
}

// 定义所有请求函数使用的统一返回数据格式 - 结构体
type Resp struct {
	Code string
	Msg  string
	Data interface{}
}
