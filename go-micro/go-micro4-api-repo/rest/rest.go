package main

import (
	"context"
	"fmt"
	proto "gomicro/message"
	"log"

	"github.com/emicklei/go-restful" // go rest框架的其中一个框架，类似的由beego、gin等等
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

// （重点）全局变量 - StudentService服务（message.pb.micro.go）
var (
	cli proto.StudentService
)

type Student struct{}

func (s *Student) GetStudent(req *restful.Request, rsp *restful.Response) {
	// URL路径参数key - name
	name := req.PathParameter("name")
	fmt.Println(name)

	// （重点）这里开始rpc调用
	response, err := cli.GetStudent(context.TODO(), &proto.Request{
		Name: name,
	})

	if err != nil {
		fmt.Println(err.Error())
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}

func main() {
	// Registry组件
	consulReg := consul.NewRegistry(
		registry.Addrs("10.1.0.13:8500"),
	)

	service := web.NewService(
		web.Name("go.micro.api.student"), // 服务名，未注册到"注册中心"
		web.Address(":8081"),
		web.Registry(consulReg),
	)
	service.Init()

	// （重点）类似客户端的使用 -- 去Registry组件中查找server服务 “go.micro.srv.student”
	cli = proto.NewStudentService("go.micro.srv.student", client.DefaultClient)

	// 构建restful api
	ws := new(restful.WebService)
	ws.Path("/student")
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)

	student := new(Student)
	ws.Route(ws.GET("/{name}").To(student.GetStudent))

	wc := restful.NewContainer()
	wc.Add(ws)

	service.Handle("/", wc) // 注册web api功能

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
