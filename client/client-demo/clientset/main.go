package main

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	/*
		自能访问kubernetes自身内置的资源，不能直接访问CRD自定义资源
		如需clientset访问CRD自定义资源，可以通过client-gen代码生成器重新生成clientSet
		在clientset集合中自动生成与CRD操作相关的接口！

		ClientSet在RESTClient的基础上封装了Resource和Version的管理方法，
		1. 每一个Resource可以理解为一个客户端，而ClientSet就是多个客户端的集合
		2. 每一个资源和版本都是以函数的形式暴露给开发者

		RbacV1
		CoreV1
		NetWorkV1 接口函数
	*/

	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config) // 实例化clientset对象，该对象用于管理所有Resource的客户端
	if err != nil {
		panic(err)
	}

	// core核心组v1版本的pod资源
	PodsResult := getPods(clientset)
	for _, d := range PodsResult.Items {
		fmt.Printf("Namespace: %v \t Name:%v \t Statu:%v\n", d.Namespace, d.Name, d.Status.Phase)
	}

	fmt.Println("===========")

	// apps组v1版本的deployment资源
	DeploymentResult := getDeployments(clientset)
	for _, d := range DeploymentResult.Items {
		fmt.Printf("Namespace: %v \t Name:%v \n", d.Namespace, d.Name)
	}
}

func getPods(clientSet *kubernetes.Clientset) *apiv1.PodList {
	//.组版本.资源对象
	podClient := clientSet.CoreV1().Pods(apiv1.NamespaceDefault)

	list, err := podClient.List(metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}

	return list
}

func getDeployments(clientSet *kubernetes.Clientset) *appsv1.DeploymentList {
	//.组版本.资源对象
	deploymentClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	list, err := deploymentClient.List(metav1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}

	return list
}
