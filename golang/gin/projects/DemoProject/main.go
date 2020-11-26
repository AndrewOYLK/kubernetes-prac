package main

import (
	"bufio"
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"samples/projects/DemoProject/param"
	"samples/projects/DemoProject/tool"
)

func main() {
	engine := gin.Default()

	// 读取配置文件
	file, err := os.Open("./config/app.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var cfg param.AppConfig
	reader := bufio.NewReader(file)
	tool.Decode(reader, &cfg)

	engine.POST("/hello", func(context *gin.Context) {
		var loginData param.LoginData
		// 方法1:
		//if err := context.BindJSON(&loginData); err != nil {
		//	fmt.Println(err.Error())
		//}
		// 方法2:
		fmt.Println(context.Request.Body)
		//if err := json.NewDecoder(context.Request.Body).Decode(&loginData); err != nil {
		//	fmt.Println(err.Error())
		//  return
		//}
		// 方法3:
		if err := tool.Decode(context.Request.Body, &loginData); err != nil {
			fmt.Println(err.Error())
			return
		}
		context.JSON(200, map[string]interface{}{
			"code": "0",
			"msg":  "OK",
			//"data": loginData,
			"data1": cfg,
			"data2": loginData,
		})
	})

	engine.Run()
}
