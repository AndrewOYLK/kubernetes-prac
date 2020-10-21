package main

import (
	"fmt"
	"github.com/AndrewOYLK/k8scode/kube"
)

func main() {
	conf := kube.GetConfigFromFile()
	fmt.Println(conf.Insecure)
}
