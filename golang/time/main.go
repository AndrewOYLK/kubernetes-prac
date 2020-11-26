package main

import (
	"fmt"
	"time"
)

func main() {
	// Time => Unix
	// 准确
	var custom = "2006-01-02T15:04:05Z07:00"

	timeStr := "2020-07-13T00:00:00+08:00"
	t, _ := time.Parse(custom, timeStr)
	fmt.Println("开始:", t.Unix())

	timeStr = "2020-07-13T23:59:59+08:00"
	t, _ = time.Parse(custom, timeStr)
	fmt.Println("结束:", t.Unix())

	/*
		开始: 1594625400
		结束: 1594625402
	*/

	// 	now := time.Now()
	// 	fmt.Println(now.Unix())
	// 	fmt.Println(now.UTC())
	// 	fmt.Println(now.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))

	// Unix => Time
	t = time.Unix(1594483200, 0)
	fmt.Println(t.Format(custom))

	t = time.Unix(1594569540, 0)
	fmt.Println(t.Format(custom))
}
