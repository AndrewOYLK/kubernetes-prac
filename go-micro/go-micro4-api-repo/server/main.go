package main

import (
	"context"
	"errors"
	"fmt"
	proto "gomicro/message"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	_ "github.com/micro/go-plugins/registry/consul" // 必须引入
)

type StudentServiceImpl struct{}

func (ss *StudentServiceImpl) GetStudent(ctx context.Context, request *proto.Request, resp *proto.Student) error {

	studentMap := map[string]proto.Student{
		"davie":  proto.Student{Name: "davie", Classes: "软件工程专业", Grade: 80},
		"steven": proto.Student{Name: "steven", Classes: "计算机科学与技术", Grade: 90},
		"tony":   proto.Student{Name: "tony", Classes: "计算机网络工程", Grade: 85},
		"jack":   proto.Student{Name: "jack", Classes: "工商管理", Grade: 96},
	}

	if request.Name == "" {
		return errors.New("请求参数错误，请重新请求")
	}

	student := studentMap[request.Name] // 根据name值来查询信息
	if student.Name != "" {
		fmt.Println(student.Name, student.Classes, student.Grade)
		*resp = student
		return nil
	}
	return errors.New("未查询到")
}

func main() {
	// Registry组件
	consulReg := consul.NewRegistry(
		registry.Addrs("10.1.0.13:8500"),
	)

	// 服务端声明
	service := micro.NewService(
		micro.Name("go.micro.srv.student"),
		micro.Registry(consulReg),
	)
	service.Init()

	// 注册服务的 “方法”
	proto.RegisterStudentServiceHandler(service.Server(), new(StudentServiceImpl))

	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
