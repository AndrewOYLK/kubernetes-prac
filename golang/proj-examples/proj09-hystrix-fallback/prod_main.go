package main

import (
	"context"
	"fmt"
	"proj09-hystrix-fallback/Services"
	"proj09-hystrix-fallback/Weblib"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

// 嵌套的方式（类似继承）
type logWrapper struct {
	// go里面没有java那样的继承，我们使用嵌套struct的方式来做
	client.Client
}

// 包装器
// 装饰 - 拦截
// 重写 call方法
func (this *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("调用接口")
	// 日志处理
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return this.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	// 这里的client不需要主动传入，他初始化的时候自动会传入，底层机制
	return &logWrapper{
		c,
	}
}

func main() {
	// consul注册中心部分
	consulReg := consul.NewRegistry(
		registry.Addrs("10.1.0.13:8500"), // consul地址
	)

	// 定义rpc服务实例
	myService := micro.NewService(
		micro.Name("prodservice.client"),
		micro.WrapClient(NewLogWrapper), // 意思：把客户端包起来
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
