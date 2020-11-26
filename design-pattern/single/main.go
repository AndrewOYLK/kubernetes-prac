package main

import (
	"fmt"
	"sync"
)

type test struct {
	Name string
}

var t *test
var once sync.Once

func getInterface() *test {
	// once.Do(func() { t = &test{"Andrew"} })
	t = &test{"Andrew"}
	return t
}

func main() {
	i1 := getInterface()
	i2 := getInterface()

	if i1 == i2 {
		fmt.Println("相同")
	} else {
		fmt.Println("不相同")
	}
}
