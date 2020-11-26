package main

import (
	"fmt"
	"time"
)

var msg = make(chan string)

func sample() {
	time.Sleep(2 * time.Second)
	str := <-msg
	str += "I 'm goroutine!"
	msg <- str
}

func main() {

	go func() {
		msg <- "Hello goroutine!"
		// fmt.Println("Hello goroutine!") // 标准输入、标准输出（goroutine争抢）
	}()

	go sample()
	// go func() {
	// 	time.Sleep(2*time.Second)
	// 	str := <-msg
	// 	str += "I 'm goroutine!"
	// 	msg <- str
	// }()

	time.Sleep(3 * time.Second)
	fmt.Println(<-msg)
	fmt.Println("Hello World!")
}
