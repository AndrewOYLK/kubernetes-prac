package controller

import (
	"fmt"
	"github.com/AndrewOYLK/k8scode/kube"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"time"
)

// （自定义）Pod 控制器结构体
// 将Reflector、Informer、Indexer、WorkQueue组合起来
type Controller struct {
	indexer cache.Indexer // 缓存
	queue workqueue.RateLimitingInterface // 工作队列（限速队列）
	informer cache.Controller
}

func NewController(queue workqueue.RateLimitingInterface, indexer cache.Indexer, informer cache.Controller) *Controller {
	return &Controller{
		indexer: indexer,
		queue: queue,
		informer: informer,
	}
}

// 控制循环
func (c *Controller) Run(threadiness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()
	// 停止控制器后，需要关掉队列
	defer c.queue.ShutDown()

	// 启动控制器
	klog.Infof("starting pod controller")

	// 启动通用控制器框架
	go c.informer.Run(stopCh)
	// 等待所有相关的缓存同步完成，然后再开始处理队列中的数据
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("time out waiting for cache to sync"))
		return
	}

	// 启动worker处理元素
	for i:=0; i<threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<- stopCh
	klog.Info("stopping pod controller")
}

// 处理元素
func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

func (c *Controller) processNextItem() bool {
	// 从workqueue里面取出一个元素
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	// 告诉队列已经处理了该key
	defer c.queue.Done(key)

	// 根据key去处理业务逻辑
	err := c.syncToStdout(key.(string))
	c.handleErr(err, key)
	return true
}

// 业务逻辑处理
func (c *Controller) syncToStdout(key string) error {
	// 从indexer中获取资源对象
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		klog.Errorf("Fetch object with key %s from indexer failed with %v", key, err)
		return err
	}

	if !exists {
		fmt.Printf("Pod %s does not exist anymore\n", key)
	} else {
		fmt.Printf("Sync/Add/Update for pod %s\n", obj.(*v1.Pod).GetName())
	}
	return nil
}

func (c *Controller) handleErr(err error, key interface{}) {
	if err == nil {
		c.queue.Forget(key)
		return
	}

	// 出现问题，允许当前控制器重试5次
	if c.queue.NumRequeues(key) < 5 {
		// 重新入队列
		c.queue.AddRateLimited(key)
		return
	}
	c.queue.Forget(key)
	runtime.HandleError(err)
	// 不允许继续重试了
}

// =======================================

func initClient() (*kubernetes.Clientset, error) {
	config := kube.GetConfigFromFile()
	// 初始化clientset对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset, nil
}

func Test() {
	clientset, err := initClient()
	if err != nil {
		klog.Fatal(err)
	}

	// 创建pod的 ListWatch
	podListWatch := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(), "pods", v1.NamespaceDefault, fields.Everything())


	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	indexer, informer:= cache.NewIndexerInformer(
		podListWatch,
		&v1.Pod{},
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Println("AddFunc")
				key, err := cache.MetaNamespaceKeyFunc(obj)
				if err == nil {
					queue.Add(key)
				}
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Println("UpdateFunc")
				key, err := cache.MetaNamespaceKeyFunc(newObj)
				if err == nil {
					queue.Add(key)
				}
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Println("DeleteFunc")
				key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
				if err == nil {
					queue.Add(key)
				}
			},
		},
		cache.Indexers{})
	// 实例化pod控制器
	controller := NewController(
		queue,
		indexer,
		informer)

	stopCh := make(chan struct{})
	defer close(stopCh)
	go controller.Run(1, stopCh)

	// 阻塞
	select{
	}
}
