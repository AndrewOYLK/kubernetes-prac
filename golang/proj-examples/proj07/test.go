package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"proj07/Models"

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
		client.ContentType("application/json"), // 需要查看go-plugins/client支持什么Content-Type
	)

	// 使用protobuf
	req := myClient.NewRequest("prodservice", "/v1/prods",
		Models.ProdsRequest{
			Size: 4,
		},
	)

	// var rsp map[string]interface{}
	var rsp Models.ProdListResponse
	err := myClient.Call(context.TODO(), req, &rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp.Data)
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
		registry.Addrs("10.1.0.13:8500"),
	)

	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin),
	)

	// 调用
	callAPI2(mySelector)
}
