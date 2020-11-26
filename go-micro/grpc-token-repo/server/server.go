package main

import (
	"context"
	"fmt"
	"net"
	"tools/golang/grpc-token-repo/message"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthServiceImpl struct{}

// 服务端定义的一个接口1
func (as AuthServiceImpl) ChkHealth(ctx context.Context, req *message.ReqNil) (*message.RespMsg, error) {
	// 在服务端的调用方法中实现对token请求参数的判断，可以通过metadata获取token认证信息
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var appKey string
	var appSecret string

	fmt.Println("接收到的客户端token值为: ", md)

	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}

	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}

	if appKey != "hello" || appSecret != "20190812" {
		return nil, status.Errorf(codes.Unauthenticated, "Token不合法")
	}

	fmt.Println("被触发了健康检查")
	return &message.RespMsg{
		Code: "00000",
		Data: "",
		Msg:  "服务端已经触发了健康检查",
	}, nil
}

// 服务端定义的一个接口2
func (as AuthServiceImpl) GetSome(ctx context.Context, req *message.ReqNil) (*message.RespMsg, error) {
	return &message.RespMsg{
		Code: "00000",
		Data: "",
		Msg:  "你获取到了礼物",
	}, nil
}

func main() {
	server := grpc.NewServer()

	// 注册服务对象（一个服务对象包含多个接口）
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
