package main

import (
	"fmt"
	"os"

	"github.com/andrewoylk/protobuf-repo/example"
	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello World. \n")

	msg_test := &example.Person{
		Name: proto.String("Andrew"),
		Age:  proto.Int(19),
		From: proto.String("China"),
	}

	// 序列化（对结构体对象指针进行序列化操作）
	msgDateEncoding, err := proto.Marshal(msg_test)
	if err != nil {
		panic(err.Error())
		return
	}
	fmt.Println("使用Protobuf协议序列化后：", msgDateEncoding)

	// 反序列化（对二进制序列进行反序列化到一个结构体对象中）
	msgEntity := example.Person{} // Go结构体对象
	err = proto.Unmarshal(msgDateEncoding, &msgEntity)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	fmt.Println("使用Protobuf协议反序列后：", msgEntity)

	fmt.Println("===========================")
	fmt.Printf("姓名: %s\n\n", msgEntity.GetName())
	fmt.Printf("年龄: %d\n\n", msgEntity.GetAge())
	fmt.Printf("国籍: %s\n\n", msgEntity.GetFrom())
}
