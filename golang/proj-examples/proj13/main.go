package main

import (
	"proj13/Services"
	"proj13/ServicesImpl"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("10.1.0.13:8500"),
	)
	myService := micro.NewService(
		micro.Name("api.jtthink.com.test"), // api.jtthink.com是命名空间，test是服务名
		micro.Address(":8001"),
		micro.Registry(consulReg),
	)

	// 服务的注册
	Services.RegisterTestServiceHandler(
		myService.Server(),
		new(ServicesImpl.TestService),
	)

	myService.Run()
}
