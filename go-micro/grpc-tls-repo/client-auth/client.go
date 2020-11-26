package main

import (
	"context"
	"grpc2/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// 普通连接方式（Insecure：不安全的）
	// conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())

	// 带证书的TLS连接（注意这个名字: andrew，在openssl里填写的）
	creds, err := credentials.NewClientTLSFromFile("../files/server.pem", "test")
	if err != nil {
		panic(err.Error())
	}
	conn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	client := message.NewAuthServiceClient(conn)

	_, err = client.ChkHealth(
		context.Background(),
		&message.ReqNil{
			Req: 1,
		},
	)
	if err != nil {
		grpclog.Fatal(err.Error())
	}
}
