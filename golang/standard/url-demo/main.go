package main

import (
	"fmt"
	"net/url"
)

var URL = "https://gitbook.cn/gitchat/column/5dca675eb104917ad887b388/topic/5dca6bdcb104917ad887b3fd"

func main() {
	u, _ := url.Parse(URL)
	fmt.Println(u)
	fmt.Println(u.Port())
	fmt.Println(u.RequestURI())
}
