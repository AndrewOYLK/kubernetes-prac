package rpc

import "fmt"

func main() {
	var a, b int
	a = 1
	b = 2
	c := Add(a, b)
	fmt.Println("计算结果:", c)
}

func Add(a int, b int) int {
	return a + b
}
