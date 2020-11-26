package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// fmt.Println(time.Now().Add(2 * time.Minute))

	d := time.Now().Add(60 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	fmt.Println(time.Now())
	select {
	case <-time.After(10 * time.Second): // 这里time.After返回一个channel类型的值（<-chan Time）
		fmt.Println(time.Now())
		return
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
