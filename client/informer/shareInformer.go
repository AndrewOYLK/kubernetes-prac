package informer

import (
	"fmt"
	"github.com/AndrewOYLK/kubernetes-prac/client/informer/kube"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"time"
)

func ShareInformers() {
	config := kube.GetConfigFromFile()
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 设置每30s重新List一次（本地缓存放到informer里面来）
	informerFactor := informers.NewSharedInformerFactory(clientSet, 30*time.Second)
	deployInformer := informerFactor.Apps().V1().Deployments()

	// 创建informer，相当于注册到工厂中去，这样下面就会去list & watch对应的资源
	informer := deployInformer.Informer() // 这里初始化Informer的时候，他就会带上怎么去list怎么去watch！

	deployLister := deployInformer.Lister()

	// 注册事件处理函数 + 业务逻辑
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: onAdd,
		UpdateFunc: onUpdate,
		DeleteFunc: onDelete,
	})

	stopper := make(chan struct{})
	defer close(stopper)

	// 启动Informer
	informerFactor.Start(stopper)
	// 等待所有启动的Informer的缓存同步
	informerFactor.WaitForCacheSync(stopper)

	// 从本地缓存中获取default中所有的deployment列表
	deployments, err := deployLister.Deployments("default").List(labels.Everything())
	if err != nil {
		panic(err)
	}

	for idx, deploy := range deployments {
		fmt.Printf("%d -> %s\n", idx+1, deploy.Name)
	}
}
