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
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)

	// 组装请求信息
	orderRequest := &message.OrderRequest{
		OrderId:   "201907300001",
		TimeStamp: time.Now().Unix(),
	}

	// 接收订单信息
	orderInfoClient, err := orderServiceClient.GetOrderInfos(
		context.Background(),
		orderRequest,
	)

	// 数据流的读取！
	for {
		orderInfo, err := orderInfoClient.Recv()
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("读取到的信息: ", orderInfo)
	}
}
