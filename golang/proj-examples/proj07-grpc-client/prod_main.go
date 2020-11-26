package main

import (
	"proj07-grpc-client/Services"
	"proj07-grpc-client/Weblib"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// consul注册中心部分
	consulReg := consul.NewRegistry(
		registry.Addrs("10.1.0.13:8500"), // consul地址
	)

	// 定义rpc服务实例
	myService := micro.NewService(
		micro.Name("prodservice.client"),
	)
	prodService := Services.NewProdService(
		"prodservice-haha", // 具体在consul上的"服务名"
		myService.Client(),
	)

	// web服务
	httpServer := web.NewService(
		web.Name("httpprodservice"), // web服务
		web.Address(":8001"),
		web.Handler(Weblib.NewGinRouter(prodService)),
		web.Registry(consulReg),
	)
	httpServer.Init()
	httpServer.Run()
}
