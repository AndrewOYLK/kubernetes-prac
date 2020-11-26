package main

import (
	"errors"
	"math"
	"net"
	"net/http"
	"net/rpc"
	"rpc2/message"
	"rpc2/param"
	"time"
)

// 服务定义1 =======================================================================

type MathUtil struct {
}

func (mu MathUtil) CalculateCircleArea(req float32, resp *float32) error {
	// 单个数值的返回
	*resp = math.Pi * req * req
	return nil
}

func (mu MathUtil) Add(param param.AddParma, resp *float32) error {
	// 按照AddParam结构 “接收“
	// 单个数值的返回
	*resp = param.Args1 + param.Args2
	return nil
}

// 服务定义2 =======================================================================

type OrderService struct{}

// GetOrderInfo 这里的 “输出方法” 使用了protobuf定义的结构体（Protobuf格式数据与RPC结合）
func (os OrderService) GetOrderInfo(request message.OrderRequest, response *message.OrderInfo) error {

	// 订单信息集合
	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}

	current := time.Now().Unix()
	if request.TimeStamp > current {
		*response = message.OrderInfo{
			OrderId:     "0",
			OrderName:   "",
			OrderStatus: "订单信息异常",
		}
	} else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			*response = orderMap[request.OrderId]
		} else {
			return errors.New("server error")
		}
	}
	return nil
}

func main() {
	// 1. 初始化指针数据类型
	mathUtil := new(MathUtil)

	// 2. （方法1）调用net/rpc包的功能将服务对象进行注册（服务注册）
	err := rpc.Register(mathUtil)
	if err != nil {
		panic(err.Error())
	}

	// 2. （方法2）服务注册 - 注册结构体对象
	err = rpc.RegisterName("TestName", mathUtil)
	if err != nil {
		panic(err.Error())
	}

	// 2. 服务注册
	orderService := new(OrderService)
	err = rpc.Register(orderService) // 注册结构体对象
	if err != nil {
		panic(err.Error())
	}

	// 3. 通过该函数把mathUtil中提供的服务注册到HTTP协议上，方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()

	// 4. 在特定的端口进行监听
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	http.Serve(listen, nil)
}
