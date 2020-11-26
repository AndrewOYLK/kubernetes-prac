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
	return &message.RespMsg{
		Code: "00000",
		Data: "",
		Msg:  "服务端已经被触发了接口1",
	}, nil
}

// 服务端定义的一个接口2
func (as AuthServiceImpl) GetSome(ctx context.Context, req *message.ReqNil) (*message.RespMsg, error) {
	return &message.RespMsg{
		Code: "00000",
		Data: "",
		Msg:  "服务端已经被触发了接口2",
	}, nil
}

func main() {
	/*
		1. 使用拦截器，首先需要注册
		2. grpc框架通过grpc.UnaryInterceptor()方法设置自定义的拦截器，并返回ServerOption
		3. UnaryInterceptor()接收一个UnaryServerInterceptor类型，继续查看源码定义，可以发现UnaryServerInterceptor是一个func
		type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
		4. 如果开发者需要注册自定义拦截器，需要自定义实现UnaryServerInterceptor的定义
	*/

	// 进行拦截器的注册
	server := grpc.NewServer(grpc.UnaryInterceptor(TokenInterceptor))

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

// 自定义实现func,符合UnaryServerInterceptor的标准，在该func的定义中实现对token的验证逻辑
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	fmt.Println("被拦截器截取，正在进行检测，请遵循规则！")

	// metadata中取出请求头中携带的token认证信息
	md, exist := metadata.FromIncomingContext(ctx)

	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "无Token认证")
	}

	var appKey string
	var appSecret string
	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}
	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}

	if appKey != "hello" || appSecret != "20190812" {
		return nil, status.Errorf(codes.Unauthenticated, "Token不合法")
	}

	// （通过token验证，继续处理请求）如果token验证通过，则继续处理请求后续逻辑，后续继续处理可以由grpc.UnaryHandler进行处理
	return handler(ctx, req)
}
