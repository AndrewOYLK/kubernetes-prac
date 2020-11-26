package main

import (
	"encoding/json"
	"fmt"
	"gomicro/message"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/mqtt"
)

func main() {
	service := micro.NewService(
		micro.Name("student_client"),
		micro.Broker(mqtt.NewBroker()),
	)
	service.Init()

	// 连接消息组件
	brok := service.Server().Options().Broker // 返回：接口实例
	if err := brok.Connect(); err != nil {
		log.Fatal(" broker connection failed, error : ", err.Error())
	}

	fmt.Println(brok.Address())

	// 定义消息体
	student := &message.Student{Name: "davie", Classes: "软件工程专业", Grade: 80}
	msgBody, err := json.Marshal(student) // 类型：[]byte
	if err != nil {
		log.Fatal(err.Error())
	}
	msg := &broker.Message{
		Header: map[string]string{
			"name": student.Name,
		},
		Body: msgBody,
	}

	// 发布消息
	err = brok.Publish("go.micro.srv.message", msg)
	if err != nil {
		log.Fatal(" 消息发布失败：%s\n", err.Error())
	} else {
		log.Print("消息发布成功")
	}
}
