package main

import (
	"proj15-valid/Services"
	"proj15-valid/ServicesImpl"

	_ "proj15-valid/AppInit" // （重要）只引用，当引用的时候，它会执行这个包里的所有init函数

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
)

func main() {
	// consulReg := consul.NewRegistry(
	// 	registry.Addrs("192.168.189.128:8500"),
	// )
	etcdReg := etcd.NewRegistry(
		registry.Addrs("192.168.189.128:23791"),
	)

	myService := micro.NewService(
		micro.Name("api.jtthink.com.myapp"), // api.jtthink.com是命名空间，test是服务名
		micro.Address(":8001"),
		// micro.Registry(consulReg),
		micro.Registry(etcdReg),
	)

	Services.RegisterUserServiceHandler(
		myService.Server(),
		new(ServicesImpl.UserService),
	)

	myService.Run()
}
