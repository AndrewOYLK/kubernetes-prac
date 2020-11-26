package main

import (
	"context"
	"errors"
	"fmt"
	"gomicro/message"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
	_ "github.com/micro/go-plugins/registry/consul" // 必须引入
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
	// 创建一个服务对象实例（go-micro在创建服务时提供了很多可选项配置，其中就包含服务组件的指定）
	service := micro.NewService(
		micro.Name("student_service"),
		micro.Version("v1.0.0"),
		micro.Registry(consul.NewRegistry()), // 通过micro.Registry可以指定要注册的发现组件
	)

	// 服务初始化
	service.Init()

	// 注册
	message.RegisterStudentServiceHandler(service.Server(), new(StudentManager))

	// 运行
	err := service.Run()
	if err != nil {
		log.Fatal(err)
	}
}
