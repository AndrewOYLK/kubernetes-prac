package main

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1" // Deployment对象的结构体类型
	corev1 "k8s.io/api/core/v1" // Pod对象的结构体类型

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" // 限制显示相关

	"k8s.io/client-go/kubernetes/scheme" // 序列化与反序列化相关
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd" // 配置文件对象相关
)

func main() {
	resultPods := getPodList()
	for _, d := range resultPods.Items {
		fmt.Printf("Namespace: %v \t Name:%v \t Statu:%v\n", d.Namespace, d.Name, d.Status.Phase)
	}

	resultDeployments := getDeploymentList()
	for _, d := range resultDeployments.Items {
		fmt.Printf("Namespace: %v \t Name:%v \t \n", d.Namespace, d.Name)
	}
}

func getPodList() *corev1.PodList {
	// kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}

	var result = &corev1.PodList{}

	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion // 请求的资源组 + 资源版本 （ schema.GroupVersion{Group: "", Version: "v1"} ）
	config.NegotiatedSerializer = scheme.Codecs      // 数据的编解码器

	restClient, err := rest.RESTClientFor(config) // 实例化
	if err != nil {
		panic(err)
	}

	// 以下进行对restClient对象构建HTTP请求参数
	h := restClient.Get()
	fmt.Println(h.URL()) // https://10.200.10.77:6443/api/v1

	err = h.Namespace("default").Resource("pods").VersionedParams( // 查询选项（如limit、timeoutseconds等添加到请求参数中）
		&metav1.ListOptions{Limit: 2}, // 限制显示pod的个数
		scheme.ParameterCodec).Do().Into(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(h.URL()) // https://10.200.10.77:6443/api/v1/namespaces/default/pods?limit=2

	return result
}

func getDeploymentList() *appsv1.DeploymentList {
	// kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}

	var result = &appsv1.DeploymentList{}

	config.APIPath = "apis" // 设置请求的HTTP路径
	config.GroupVersion = &appsv1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs // 数据的编解码器

	restClient, err := rest.RESTClientFor(config) // 实例化
	if err != nil {
		panic(err)
	}

	// 以下进行对restClient对象构建HTTP请求参数
	h := restClient.Get()
	fmt.Println(h.URL()) // https://10.200.10.77:6443/api/v1
	err = h.Namespace("default").
		Resource("deployments").
		VersionedParams( // 查询选项（如limit、timeoutseconds等添加到请求参数中）
			&metav1.ListOptions{Limit: 2}, // 限制显示pod的个数
			scheme.ParameterCodec).Do().Into(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(h.URL()) // https://10.200.10.77:6443/api/v1/namespaces/default/pods?limit=2

	return result
}
