package kube

import (
	"fmt"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

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