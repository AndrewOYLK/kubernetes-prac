package main

import (
	"proj12-hystrix/server/ServiceImpl"
	"proj12-hystrix/server/Services"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// 空服务，没有Handler
	consulReg := consul.NewRegistry(
		registry.Addrs("10.1.0.13:8500"),
	)

	prodService := micro.NewService(
		micro.Name("prodservice-haha"),
		micro.Address(":8011"),
		micro.Registry(consulReg),
	)
	prodService.Init()

	// 设置Handler，并注册到consul
	Services.RegisterProdServiceHandler(prodService.Server(), new(ServiceImpl.ProdService))
	prodService.Run()
}
