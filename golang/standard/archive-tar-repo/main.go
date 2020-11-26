package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("test.txt")

	w := tar.NewWriter(file)

	var tmp = make([]byte, 20)

	n, _ := w.Read(tmp)

	fmt.Println(n)
	fmt.Println(tmp)
}
