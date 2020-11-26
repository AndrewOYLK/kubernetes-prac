package main

import "fmt"

func main() {
	tmp := "abcd"
	for i, v := range tmp {
		fmt.Println(i, v)
	}
}
