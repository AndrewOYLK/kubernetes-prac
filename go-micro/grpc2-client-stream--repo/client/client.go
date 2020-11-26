package main

import (
	"context"
	"fmt"
	"grpc2/message"
	"io"
	"time"

	"google.golang.org/grpc"
)

func main() {
	orderMap := []message.OrderRequest{
		message.OrderRequest{OrderId: "201907300001", TimeStamp: time.Now().Unix()},
		message.OrderRequest{OrderId: "201907310001", TimeStamp: time.Now().Unix()},
		message.OrderRequest{OrderId: "201907310002", TimeStamp: time.Now().Unix()},
	}

	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)

	addOrderListClient, err := orderServiceClient.AddOrderList(context.Background()) // 客户端调用服务端接口AddOrderList
	if err != nil {
		panic(err.Error())
	}

	for _, info := range orderMap {
		time.Sleep(5 * time.Second)
		err = addOrderListClient.Send(&info)
		if err != nil {
			panic(err.Error())
		}
	}

	for {
		orderInfo, err := addOrderListClient.CloseAndRecv()
		if err == io.EOF {
			fmt.Println("读取数据结束了")
			return
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(orderInfo.GetOrderStatus())
		fmt.Println(orderInfo.GetOrderId())
		fmt.Println(orderInfo.GetOrderName())
	}
}
