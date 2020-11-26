package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("test.txt") // 返回*File类型
	defer f.Close()

	r := bufio.NewReader(f) // 定义一个用户空间缓冲区（大小4096）
	fmt.Println(r.Size())

	buf := make([]byte, 10)
	n, err := r.Read(buf)

	fmt.Println(n, err)
}
