package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()

	// h.Write([]byte("i am andrew"))
	io.WriteString(h, "i am andrew")

	fmt.Printf("%x\n", h.Sum(nil))
}
