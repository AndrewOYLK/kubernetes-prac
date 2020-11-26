package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	client := http.DefaultClient

	rsp, _ := client.Get("http://www.baidu.com")

	/*
		解释：这个resp.Body实例是io.ReadCloser类型的
	*/
	lr := io.LimitReader(rsp.Body, 10)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
