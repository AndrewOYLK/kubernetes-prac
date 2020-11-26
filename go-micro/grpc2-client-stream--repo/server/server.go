package main

import (
	"fmt"
	"grpc2/message"
	"io"
	"net"

	"google.golang.org/grpc"
)

var orderMap = map[string]message.OrderInfo{
	"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
	"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
	"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
}

type OrderServiceImpl struct{}

func (os *OrderServiceImpl) AddOrderList(stream message.OrderService_AddOrderListServer) error {
	fmt.Println("客户端流 RPC 模式")

	for {
		// 从流中读取数据信息
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("读取数据结束")
			result := message.OrderInfo{ // 返回的字段不一定都要协商！
				OrderStatus: "读取数据结束",
			}
			return stream.SendAndClose(&result)
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		fmt.Println(orderMap[orderRequest.OrderId])
	}
}

func main() {
	server := grpc.NewServer()

	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
