package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	tcpAddr := net.TCPAddr{
		IP:   net.ParseIP("192.168.189.128"),
		Port: 8081,
	}

	tcplistener, err := net.ListenTCP("tcp4", &tcpAddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		fmt.Println("正在等待TCP连接")
		conn, _ := tcplistener.Accept()
		go handler(conn)
	}

}

func handler(conn net.Conn) {
	// data, err := ioutil.ReadAll(conn)

	data := make([]byte, 1024)
	r := bufio.NewReader(conn)
	for {
		n, err := r.Read(data)
		if err != nil && err.Error() == "EOF" {
			break
		}
		fmt.Println(n)
		fmt.Println(string(data))
	}
	fmt.Println("接收结束")
}
