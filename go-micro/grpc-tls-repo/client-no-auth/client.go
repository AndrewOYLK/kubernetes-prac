package main

import (
	"context"
	"grpc2/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
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
