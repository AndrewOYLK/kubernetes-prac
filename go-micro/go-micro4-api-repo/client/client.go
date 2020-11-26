package main

import (
	"context"
	"fmt"
	proto "gomicro/message"

	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul" // 必须引入
)

func main() {
	service := micro.NewService(
		micro.Name("student_client"),
	)
	service.Init()

	// 注意：这里需要依赖pb.go文件中的NewStudentServiceClient方法
	studentService := proto.NewStudentService("go.micro.srv.student", service.Client())

	/*
		注意：
			这里调用的GetStudent方法是服务端定义的，这里不太解耦！！！
			如果服务名称更改了怎么办？虽然服务端可以重新注册一个服务到注册中心
			但是客户端已经固定写了调用那个方法名称！！！
			不可能一次服务更新，就需要连同客户端也一起更新！！！
	*/
	res, err := studentService.GetStudent(context.TODO(), &proto.Request{Name: "davie"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("==================")
	fmt.Println(res)
	fmt.Println(res.Name)
	fmt.Println(res.Classes)
	fmt.Println(res.Grade)
}
