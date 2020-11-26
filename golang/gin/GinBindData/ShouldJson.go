package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	engine := gin.Default()

	engine.POST("/addstudent", func(context *gin.Context) {
		buf := bytes.Buffer{}
		bs, _ := ioutil.ReadAll(context.Request.Body)
		json.Marshal()
		buf.Write(bs)
		fmt.Println(buf.String())
		// tmp, _ :=
		//fmt.Println(string(tmp))
		context.Writer.Write([]byte("添加记录: "))
	})


	//engine.POST("/addstudent", func(context *gin.Context) {
	//	fmt.Println(context.FullPath())
	//
	//	// 以JSON形式提交上来
	//	var person Person
	//	if err := context.BindJSON(&person); err != nil {
	//		log.Fatal(err.Error())
	//	}
	//	fmt.Println("姓名: ", person.Name)
	//	fmt.Println("性别: ", person.Sex)
	//	fmt.Println("年龄: ", person.Age)
	//
	//	context.Writer.Write([]byte("添加记录: " + person.Name))
	//})

	engine.Run()
}

type Person struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
	Sex  string `form:"sex"`
}
