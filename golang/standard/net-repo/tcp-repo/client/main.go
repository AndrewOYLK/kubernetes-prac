package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// 客户端
	ltcpAddr := net.TCPAddr{
		IP:   net.ParseIP("192.168.189.128"),
		Port: 8082,
	}

	// 服务端
	rtcpAddr := net.TCPAddr{
		IP:   net.ParseIP("192.168.189.128"),
		Port: 8081,
	}

	tcpConn, err := net.DialTCP("tcp4", &ltcpAddr, &rtcpAddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(tcpConn)
	tcpConn.SetKeepAlive(true)
	tcpConn.SetKeepAlivePeriod(5 * time.Second)

	n := 0
	for {
		if n == 20 {
			break
		}
		tcpConn.Write([]byte("Hello everyone, i am coming."))
		time.Sleep(2 * time.Second)
		n++
	}
	fmt.Println("发送结束")
}
