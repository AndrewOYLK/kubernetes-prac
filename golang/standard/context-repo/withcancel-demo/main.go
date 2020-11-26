package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	cancel()
	for {
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		fmt.Println("start")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("到这里啦")
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
