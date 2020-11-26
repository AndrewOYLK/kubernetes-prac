package main

import (
	"net/http"

	"github.com/micro/go-micro/web"
)

func main() {
	addr := func(op *web.Options) {
		op.Address = ":8081"
	}

	// server := web.NewService(web.Address(":8081"))
	server := web.NewService(addr)

	// 路由处理
	server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 比较原始的方式
		// 不管是参数的获取，包括输出的内容，都需要手动完成，这个不太方便，一般都会使用比较流行的框架
		writer.Write([]byte("Hello World"))
	})

	server.Run()
}
