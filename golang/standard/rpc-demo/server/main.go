package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// 配套写法
type Args struct {
	A, B int
}

type Test struct{}

func (this *Test) Multiply(args *Args, reply *int) error {
	fmt.Println("=====")
	*reply = args.A * args.B
	return nil
}

// 主方法
func main() {
	rpc_server := rpc.NewServer()
	rpc_server.RegisterName("test", &Test{})

	l, e := net.Listen("tcp4", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	fmt.Println("正在监听...")
	rpc_server.Accept(l)
}
