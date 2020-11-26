package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func main() {
	conf,err := clientcmd.BuildConfigFromFlags("","/root/.kube/config")
	if err != nil {
		panic(err)
	}
	//新建discovery对象
	disClient,err := discovery.NewDiscoveryClientForConfig(conf)
	if err != nil {
		panic(err)
	}
	//获取资源列表
	_,apisource,err := disClient.ServerGroupsAndResources()
	if err != nil {
		panic(err)
	}
	for _,list := range  apisource {
		groupv,_ := schema.ParseGroupVersion(list.GroupVersion)
		for _,resource := range list.APIResources {
			fmt.Printf("name=%v\tgroup=%v\tVersion=%v\n",resource.Name,groupv.Group,groupv.Version)
		}
	}
}





type RateLimiter interface {
	// 获取指定的元素应该等待的时间
	When(item interface{}) time.Duration
	/*
	释放指定的元素，情况该元素的排队数
	注意：这里有一个非常重要的概念——限速周期
	一个限速周期是指从执行AddRateLimited方法到执行完Forget方法之间的时间
	如果该元素被Forget方法处理完，则清空排队数
	*/
	Forget(item interface{})
	//获取指定的元素排队位置
	NumRequeues(item interface{}) int
}