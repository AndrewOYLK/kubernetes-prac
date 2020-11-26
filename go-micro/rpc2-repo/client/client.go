package main

import (
	"fmt"
	"net/rpc"
	"rpc2/message"
	"rpc2/param"
	"time"
)

func main() {
	// 连接rpc服务端
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}

	var resp *float32

	// 同步方式 - 远端方法调用 - 1
	var req float32 = 3
	// err = client.Call("MathUtil.CalculateCircleArea", req, &resp) // 核心地方
	err = client.Call("TestName.CalculateCircleArea", req, &resp) // 核心地方
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("============ MathUtil.CalculateCircleArea 打印结果 ============")
	fmt.Println(*resp)

	// 同步调用 - 2 - 按照AddParam结构 “组装”
	var params param.AddParma = param.AddParma{
		Args1: 123,
		Args2: 123,
	}
	// err = client.Call("MathUtil.Add", params, &resp)
	err = client.Call("TestName.Add", params, &resp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("============ MathUtil.Add 打印结果 ============")
	fmt.Println(*resp)

	// 同步调用 - 3
	timeStamp := time.Now().Unix()
	request := message.OrderRequest{
		OrderId:   "201907310001",
		TimeStamp: timeStamp,
	}
	var response *message.OrderInfo
	err = client.Call("OrderService.GetOrderInfo", request, &response)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("============ OrderService.GetOrderInfo 打印结果 ============")
	fmt.Println(*response)

	// 异步调用
	// var respSync *float32
	// // syncCall := client.Go("MathUtil.CalculateCircleArea", req, &respSync, nil)
	// syncCall := client.Go("TestName.CalculateCircleArea", req, &respSync, nil)
	// replayDone := <-syncCall.Done
	// fmt.Println("============ MathUtil.CalculateCircleArea 打印结果（异步方式） ============")
	// fmt.Println(replayDone)
	// fmt.Println(*respSync)
}
