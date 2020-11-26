package main

import (
	"context"
	"fmt"
	"grpc/message"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// 客户端声明一个OrderService的客户端
	orderServiceClient := message.NewOrderServiceClient(conn)

	// 组装请求信息
	orderRequest := &message.OrderRequest{
		OrderId:   "201907300001",
		TimeStamp: time.Now().Unix(),
	}

	// 接收订单信息
	// 使用OrderService的客户端调用GetOrderInfo的接口
	orderInfo, err := orderServiceClient.GetOrderInfo( // 非流式调用，调用后就会卡着，等待服务端执行完毕
		context.Background(),
		orderRequest,
	)

	if orderInfo != nil {
		fmt.Println(orderInfo.GetOrderId())
		fmt.Println(orderInfo.GetOrderName())
		fmt.Println(orderInfo.GetOrderStatus())
	}
}
