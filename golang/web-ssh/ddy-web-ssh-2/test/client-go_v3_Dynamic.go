package main

import (
	"context"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/clock"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"sync"
	"time"
)

func main() {
	conf,err :=clientcmd.BuildConfigFromFlags("","/root/.kube/config")
	if err !=nil {
		panic(err)
	}
	//新建Dynamic的对象
	dynamicClinet,err := dynamic.NewForConfig(conf)
	if err != nil {
		panic(err)
	}
	//构建资源对象
	obdj := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}
	//构造对象信息获取所有ns下的pod信息【返回结构体对象和错误】
	resobj,err := dynamicClinet.Resource(obdj).Namespace(apiv1.NamespaceAll).List(context.Background(),metav1.ListOptions{Limit: 100})
	if err !=nil{
		panic(err)
	}
	//声明类型
	podList := &corev1.PodList{}
	//将原本的指针类型转换成podlist类型
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(resobj.UnstructuredContent(),podList)
	if err !=nil {
		panic(err)
	}
	//循环遍历所有的资源对象
	for _,va := range podList.Items {
		fmt.Printf("NameSpace: %v \t Name: %v\t label: %+v\n",va.Namespace,
			va.Name,va.Labels)
	}

}




