package main

import (
	"context"
	"fmt"
	"tools/golang/grpc-token-repo/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// 自定义的token认证结构体
type TokenAuthentication struct {
	AppKey    string
	AppSecret string
}

// 组织token信息
func (ta *TokenAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	// 注意: appid和appkey字段均需要保持小写，不能大写
	return map[string]string{
		"appid":  ta.AppKey,
		"appkey": ta.AppSecret,
	}, nil
}

// 是否基于TLS认证进行安全传输
func (ta *TokenAuthentication) RequireTransportSecurity() bool {
	return false
}

func main() {
	auth := TokenAuthentication{
		AppKey:    "hello",
		AppSecret: "20190813",
	}

	// conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())  // 不带认证规则的方法
	/*
		在gRPC中，允许开发者自定义自己的认证规则，通过grpc.WithPerRPCCredentials()
		设置自定义的认证规则。WithPerRPCCredentials方法接收一个PerRPCCredentials类型的参数 --> 这是一个接口

		type PerRPCCredentials interface {
		    GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
		    RequireTransportSecurity() bool
		}
	*/
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	client := message.NewAuthServiceClient(conn)

	// 调用服务端加了token认证的方法1
	respMsg, err := client.ChkHealth(
		context.Background(),
		&message.ReqNil{
			Req: 1,
		},
	)
	fmt.Println(respMsg.GetMsg())
	if err != nil {
		grpclog.Fatal(err.Error())
	}

	// 调用服务端没有加token认证的方法2
	respMsg, err = client.GetSome(
		context.Background(),
		&message.ReqNil{
			Req: 1,
		},
	)
	fmt.Println(respMsg.GetMsg())
	if err != nil {
		grpclog.Fatal(err.Error())
	}

}
