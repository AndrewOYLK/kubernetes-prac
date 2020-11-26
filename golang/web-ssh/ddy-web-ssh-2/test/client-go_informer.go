package main

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"time"
)


func main() {
	restconfig,err := clientcmd.BuildConfigFromFlags("","/root/.kube/config")
	if err != nil{
		panic(err)
	}
	clientset,err := kubernetes.NewForConfig(restconfig)
	if err != nil{
		panic(err)
	}
	log.Println("开始监听Informers事件流")
	//创建一个关闭通道,因为infomer是一个持久化的goroutine
	stopch := make(chan struct{})
	defer close(stopch)
	//实例化informers并且设置1分钟rsync一次
	sharedinfomers := informers.NewSharedInformerFactory(clientset,time.Minute)
	//获取具体的资源对象的informers
	cacheinfomer := sharedinfomers.Core().V1().Pods().Informer()
	//当创建资源的时候触发事件的回调方法
	cacheinfomer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mobj := obj.(v1.Object)
			log.Printf("new pod added to store: %s\t",mobj.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			old := oldObj.(v1.Object)
			new := newObj.(v1.Object)
			log.Printf("%s pod has been Update %s",old.GetName(),new.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			mobj := obj.(v1.Object)
			log.Printf("%s Pod Delete",mobj.GetName())
		},
	})
	cacheinfomer.Run(stopch)
}

