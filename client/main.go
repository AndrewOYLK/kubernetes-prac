package main

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
)

func main() {

}

func TestClientSet() {
	var err error
	var config *rest.Config
	var kubeconfig *string

	if home := homeDir(); home != "" {
		kubeconfig = flag.String(
			"kubeconfig",
			filepath.Join(home, ".kube", "config"),
			"可选",
		)
	} else {
		kubeconfig = flag.String(
			"kubeconfig",
			"",
			"必选",
		)
	}
	flag.Parse() // 解析

	//config = kube.GetConfigFromFile()
	if config, err = rest.InClusterConfig(); err != nil {
		if config, err = clientcmd.BuildConfigFromFlags(
			"",
			*kubeconfig,
		); err != nil {
			panic(err.Error())
		}
	}

	// 初始化clientset对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// metav1.ListOption用于过滤信息
	pods, err := clientset.CoreV1().Pods("kube-system").List(
		metav1.ListOptions{})
	//fmt.Println("====", pods)

	for idx, pod := range pods.Items {
		fmt.Printf("%d -> %s\n", idx+1, pod.Name)
	}
}

func homeDir() string {
	// Linux环境下
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	// Windows环境下
	return os.Getenv("USERPROFILE")
}

func DynamicClient() {
	restconfig, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}
	tektoncli, err := v1beta1.NewForConfig(restconfig)
	if err != nil {
		panic(err)
	}
	tasklist, err := tektoncli.Tasks("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, v := range tasklist.Items {
		fmt.Printf("%v", v)
	}
}
