package kube

import (
	"fmt"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

// 有时候go开发，你很难知道第三方库的某个函数是返回结构体对象指针类型呢，还是值类型，必须回到
// 源码区查看下，然后你才能在自己的函数返回类型上加上具体的类型
func GetConfigFromFile() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", fmt.Sprintf("%s\\%s", GetCurrentPath(), "kube\\config"))
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func HomeDir() string {
	// Linux环境下
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	// Windows环境下
	return os.Getenv("USERPROFILE")
}