package main

import (
	"fmt"
	"time"
)

func main() {
	// f, err := os.Create("trace.out")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// err = trace.Start(f)
	// if err != nil {
	// 	panic(err)
	// }

	// 正常业务
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello")
	}

	// trace.Stop()
}
