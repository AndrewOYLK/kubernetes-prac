package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	var result int
	client, err := rpc.Dial("tcp4", ":1234")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(client)

	client.Call("test.Multiply", &Args{100, 200}, &result)
	fmt.Println("结果：", result)
}
