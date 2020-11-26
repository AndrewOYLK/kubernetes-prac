package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

// 测试
var tick = time.Tick(1 * time.Second)
var abort = make(chan struct{})
var tmp = make(chan struct{}, 10) // 限制getService的并发数

// 连接consul注册中心
var consulReg = consul.NewRegistry(
	registry.Addrs("192.168.189.129:8500"),
)

func main() {
	var n sync.WaitGroup

	// 用来监听键盘动作
	go func() {
		// 操作系统的功能操作，比如监听标准输入和输出
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	// 不合适的位置
	// go func() {
	// 	n.Wait()
	// 	abort <- struct{}{}
	// }()

	for {
		select {
		case <-tick:
			// 这个goroutine的数目还是无限的增加，上面的tmp只能限制goroutine内的实际执行功能的次数
			n.Add(1)
			fmt.Println(n)
			go getService(&n)
		case <-abort:
			fmt.Printf("aborted!\n")
			return
		}
	}
}

func getService(n *sync.WaitGroup) {
	tmp <- struct{}{}        // 获得令牌; 如果sema channel容量满了，就会发生发送阻塞
	defer func() { <-tmp }() // 释放令牌; 如果没有发送操作，那么这里会发生接收阻塞
	defer n.Done()

	// 测试并发数量和goroutine实际数量
	time.Sleep(20 * time.Second)

	// 查找服务“prodservice”
	getService, err := consulReg.GetService("prodservice")
	if err != nil {
		log.Fatal(err)
	}

	// 轮询选择服务“prodservice”的节点信息
	next := selector.RoundRobin(getService)
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(node.Id, node.Address, node.Metadata)
}
