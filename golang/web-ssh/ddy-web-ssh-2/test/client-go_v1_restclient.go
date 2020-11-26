package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/scheme"
	"time"
)

func main(){
	cmdconf,err :=clientcmd.BuildConfigFromFlags("","/root/.kube/config")
	if err !=nil {
		fmt.Println("clientcmd.BuildConfigFromFlags:")
		panic(err)
	}
	//设置请求超时时间
	cmdconf.Timeout = 1 *time.Minute
	//设置api
	cmdconf.APIPath = "api"
	//设置注册对象的版本
	cmdconf.GroupVersion = &corev1.SchemeGroupVersion
	//设置编解码的
	cmdconf.NegotiatedSerializer = scheme.Codecs
	//设置访问的客户端,实例化对象
	restclient,err := rest.RESTClientFor(cmdconf)
	if err != nil {
		panic(err)
	}
	//构建一个对象的列表
	podlistRES := &corev1.PodList{}
	//构造请求参数从APIserver获取对象传入解码器
	if err := restclient.Get().
		Namespace("kube-system").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 500},scheme.ParameterCodec).
		Do(context.Background()).Into(podlistRES);err !=nil {
			panic(err)
	}
	for _,va := range podlistRES.Items{
		fmt.Printf("NameSpace: %v \t Name: %v\t label: %+v\n",va.Namespace,
			va.Name,va.Labels)
	}

}
