package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config := getK8sConfig()

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}

	_, APIResourceList, err := discoveryClient.ServerGroupsAndResources() // 组和资源
	if err != nil {
		panic(err)
	}

	for _, list := range APIResourceList {
		gv, err := schema.ParseGroupVersion(list.GroupVersion) // 组 + 版本
		if err != nil {
			panic(err)
		}

		for _, resource := range list.APIResources { // 资源对象
			fmt.Printf("name: %v, group:%v, version:%v, kind:%v \n", resource.Name, gv.Group, gv.Version, resource.Kind)
		}
	}
}

func getK8sConfig() *rest.Config {
	// 有时候go开发，你很难知道第三方库的某个函数是返回结构体对象指针类型呢，还是值类型，必须回到
	// 源码区查看下，然后你才能在自己的函数返回类型上加上具体的类型
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}

	return config
}
