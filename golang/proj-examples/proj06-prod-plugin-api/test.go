package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
)

/*
	服务发现，获取服务信息
*/

func callAPI2(s selector.Selector) {
	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)

	// request对象; 注册在consul的"prodservice"服务是http api服务，非rpc服务
	req := myClient.NewRequest("prodservice", "/v1/prods", map[string]string{})

	var rsp map[string]interface{}
	err := myClient.Call(context.TODO(), req, &rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp["data"])
}

// 原始调用方式
func callAPI(addr string, path string, method string) (string, error) {
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
	// 步骤1
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.189.129:8500"),
	)

	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin),
	)

	// 调用
	callAPI2(mySelector)
}
