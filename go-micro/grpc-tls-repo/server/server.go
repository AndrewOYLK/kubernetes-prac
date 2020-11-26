package main

import (
	"context"
	"fmt"
	"grpc2/message"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

type AuthServiceImpl struct{}

func (as AuthServiceImpl) ChkHealth(ctx context.Context, req *message.ReqNil) (*message.RespNil, error) {
	// fmt.Println("被触发了健康检测！")
	fmt.Println("被触发了健康检测（带证书认证）！")
	return &message.RespNil{
		Resp: 1,
	}, nil
}

func main() {
	// 普通的grpc服务对象
	// server := grpc.NewServer()

	// 开启tls认证的服务端
	creds, err := credentials.NewServerTLSFromFile("../files/server.pem", "../files/server.key")
	if err != nil {
		grpclog.Fatal("加载证书文件失败", err)
	}
	server := grpc.NewServer(grpc.Creds(creds))

	message.RegisterAuthServiceServer(
		server,
		new(AuthServiceImpl),
	)

	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(listen)
}
