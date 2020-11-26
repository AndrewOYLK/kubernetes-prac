package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

/*
	服务发现，获取服务信息
*/

// 服务调用：基本方式调用API！
func callAPI(addr string, path string, method string) (string, error) {
	// get请求
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}

func main() {
	// 【获取注册中心】连接consul注册中心
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.189.129:8500"),
	)

	// 【得到服务列表】查找服务“prodservice”列表
	getService, err := consulReg.GetService("prodservice")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getService)

	// 【选择服务实例】轮询选择服务“prodservice”的节点信息
	next := selector.RoundRobin(getService)
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(node.Id, node.Address, node.Metadata)

	// 【调用服务实例提供的http api】测试
	callRes, err := callAPI(node.Address, "/v1/prods", "GET")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(callRes)
}
