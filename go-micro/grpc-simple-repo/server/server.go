package main

import (
	"context"
	"errors"
	"fmt"
	"grpc/message"
	"net"
	"time"

	"google.golang.org/grpc"
)

type OrderServiceImpl struct {
	// 自定义了一个结构体类型，然后实现GetOrderInfo方法（该方法在message.pd.go文件中有说明，可以了解下）
}

// GetOrderInfo 关于OrderService服务的一个接口A
func (os OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {
	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}

	var response *message.OrderInfo
	current := time.Now().Unix()
	if request.TimeStamp > current {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			fmt.Println(result)
			return &result, nil
		} else {
			return nil, errors.New("server error")
		}
	}
	return response, nil
}

func main() {
	server := grpc.NewServer()

	// 注册OrderServiceImpl结构体对象
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)
}
