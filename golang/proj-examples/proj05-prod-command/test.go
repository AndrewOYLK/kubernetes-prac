package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// 连接consul注册中心
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.189.129:8500"),
	)

	// 查找服务“prodservice”
	getService, err := consulReg.GetService("prodservice")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getService)

	// 随机选择服务“prodservice”的节点信息
	// next := selector.Random(getService)
	next := selector.RoundRobin(getService)
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(node.Id, node.Address, node.Metadata)
}
