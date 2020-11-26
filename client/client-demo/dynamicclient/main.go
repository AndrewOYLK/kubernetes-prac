package main

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	/*
		dynamicclient内置实现了unstructed，用于处理非结构化数据结构（即无法提前预知的数据结构），这是dynamicclient能访问crd自定义资源的关键
	*/

	config := getK8sConfig()

	// 实例化dynamicClient对象
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// pod对象 - 资源组+资源版本+资源名称配置对象
	gvr := schema.GroupVersionResource{
		Version:  "v1",
		Resource: "pods",
	}

	unstructObj, err := dynamicClient.Resource(gvr).Namespace(apiv1.NamespaceDefault).
		List(metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}

	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), podList)
	if err != nil {
		panic(err)
	}

	for _, d := range podList.Items {
		fmt.Printf("Namespace: %v \t Name:%v \t Statu:%v\n", d.Namespace, d.Name, d.Status.Phase)
	}

	// deployment对象
	gvrD := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}

	// default命名空间
	unstructObj, err = dynamicClient.Resource(gvrD).Namespace(apiv1.NamespaceDefault).
		List(metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}

	deploymentList := &appsv1.DeploymentList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), deploymentList)
	if err != nil {
		panic(err)
	}

	for _, d := range deploymentList.Items {
		fmt.Printf("Namespace: %v \t Name:%v \t \n", d.Namespace, d.Name)
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
