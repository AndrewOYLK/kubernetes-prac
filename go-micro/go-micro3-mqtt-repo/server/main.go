package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gomicro/message"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/mqtt" // mqtt包实现了go-micro/broker定义的接口
)

type StudentManager struct {
}

func (st *StudentManager) GetStudent(ctx context.Context, request *message.StudentRequest, response *message.Student) error {
	studentMap := map[string]message.Student{
		"davie":  message.Student{Name: "davie", Classes: "软件工程专业", Grade: 80},
		"steven": message.Student{Name: "steven", Classes: "计算机科学与技术", Grade: 90},
		"tony":   message.Student{Name: "tony", Classes: "计算机网络工程", Grade: 85},
		"jack":   message.Student{Name: "jack", Classes: "工商管理", Grade: 96},
	}

	fmt.Println(request.Name)

	if request.Name == "" {
		return errors.New("请求参数错误，请重新登陆")
	}

	student := studentMap[request.Name]

	if student.Name != "" {
		*response = student
		fmt.Println(response)
		return nil
	} else {
		return errors.New("未查询到相关学生信息")
	}

}

func main() {
	// 创建一个服务对象实例
	service := micro.NewService(
		micro.Name("go.micro.srv"),
		micro.Version("latest"),
		micro.Broker(mqtt.NewBroker()), // 使用micro.Broker来指定特定的消息组件，并通过mqtt.NewBroker初始化一个mqtt实例对象,作为broker参数
	)
	// 服务初始化
	service.Init()

	// 订阅
	pubSub := service.Server().Options().Broker
	_, err := pubSub.Subscribe("go.micro.srv.message", func(event broker.Event) error {
		var req *message.StudentRequest
		if err := json.Unmarshal(event.Message().Body, &req); err != nil {
			return err
		}
		fmt.Println(" 接收到信息：", req)
		//去执行其他操作

		return nil
	})

	// 注册
	// message.RegisterStudentServiceHandler(service.Server(), new(StudentManager))

	// 运行
	err = service.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func test(service micro.Service) {
	pubSub := service.Server().Options().Broker

	for {
		_, err := pubSub.Subscribe("go.micro.srv.message", func(event broker.Event) error {
			var req *message.StudentRequest
			if err := json.Unmarshal(event.Message().Body, &req); err != nil {
				return err
			}
			fmt.Println(" 接收到信息：", req)
			//去执行其他操作
			return nil
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
