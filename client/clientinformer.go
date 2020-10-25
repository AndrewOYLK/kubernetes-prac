package client

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
	"time"
)

func TestInformer() {
	var (
		err error
		config *rest.Config
		kubeconfig *string
	)

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

	// =======================================================================
	// 初始化informer工厂
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	// 使用informer工厂拿到我们想要监听的一个资源informer
	deploymentInformer := informerFactory.Apps().V1().Deployments()

	// 这里的Informer和Lister的顺序是否很重要？
	// 创建 Lister
	deployLister := deploymentInformer.Lister()
	// 创建 Informer，将deploymentInformer “注册” 到工厂当中去；返回一个SharedIndexInformer对象
	informer := deploymentInformer.Informer()
	// 注册事件处理程序
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: onAdd, // 业务逻辑，扩展概念：CRD
		UpdateFunc: onUpdate, // 业务逻辑
		DeleteFunc: onDelete, // 业务逻辑
	})
	// =======================================================================
	// 启动informer，将全量的数据缓存下来（List全量 & Watch增量）
	stopper := make(chan struct{})
	defer close(stopper)
	informerFactory.Start(stopper) // 一个协程
	// 等待所有启动的Informer的缓存被同步
	informerFactory.WaitForCacheSync(stopper)
	// =======================================================================
	// 通过Lister获取 "缓存中" 的Deployment的数据
	deployments, err := deployLister.Deployments("default").List(labels.Everything())
	if err != nil {
		panic(err)
	}
	for idx, deploy := range deployments {
		fmt.Printf("%d -> %s\n", idx+1, deploy.Name)
	}
	<-stopper // 挂起
}

func onAdd(obj interface{}) {
	// 断言
	deploy := obj.(*v1.Deployment)
	fmt.Println("add a deployment:", deploy.Name)
}

func onUpdate(old, new interface{}) {
	oldDeploy := old.(*v1.Deployment)
	newDeploy := new.(*v1.Deployment)
	fmt.Println("update deployment", oldDeploy.Name, newDeploy.Name)
}

func onDelete(obj interface{}) {
	deploy := obj.(*v1.Deployment)
	fmt.Println("delete deployment", deploy.Name)
}