package main

import (
	"fmt"
	"k8s.io/client-go/util/workqueue"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"time"
)

var logger *zap.Logger
var queue workqueue.RateLimitingInterface

func main() {
	//初始化日志
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	logger.Info("The K8s logger started")
	//指定k8sconfig文件位置
	rest,err := clientcmd.BuildConfigFromFlags("","/root/.kube/config")
	if err != nil {
		logger.Panic(err.Error())
		os.Exit(1)
	}
	//输出读取到的用户信息
	logger.Info("kubernetes targeted",
		zap.String("host",rest.Host),
		zap.String("username",rest.Username))
	//创建clientset对象
	clientset,err := kubernetes.NewForConfig(rest)
	if err != nil {
		logger.Panic(err.Error())
		os.Exit(1)
	}
	//新建一个informers监视对象（设置再次同步的时间间隔）
	factory := informers.NewSharedInformerFactory(clientset,10*time.Second)
	//获取pod的infomer对象
	informer := factory.Core().V1().Pods().Informer()
	//构造一个阻塞的结构体
	stopper := make(chan struct{})
	//异常是注意关闭通道和informers
	defer close(stopper)
	defer runtime.HandleCrash()
	//创建一个队列
	queue = workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	defer queue.ShutDown()
	//添加监视函数，当触发事件后自动handle
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onADD2,
		UpdateFunc: onUpdate2,
		DeleteFunc: onDelete2,
	})
	logger.Info("Running Informers Success")
	go informer.Run(stopper)
	<- stopper
}

//onADD测试函数2
func onADD2(obj interface{}) {
	var (
		key string
		err error
	)
	if key,err = cache.MetaNamespaceKeyFunc(obj);err !=nil {
		runtime.HandleError(err)
		return
	}
	queue.Add(key)
}

//添加的回调函数
func onADD(obj interface{}) {
	//端砚获取POd资源对象
	pod,ok := obj.(*corev1.Pod)
	if !ok {
		logger.With(zap.String("action","create")).Error("Unserializable pod")
		return
	}
	info := fmt.Sprintln(pod.GetNamespace(),"/",pod.GetName())
	logger.With(zap.String("action","create")).Info(info)


}

//更新回调函数
func onUpdate(oldOBJ,newOBJ interface{}) {
	//获取pod对象
	n := newOBJ.(*corev1.Pod)
	o := oldOBJ.(*corev1.Pod)
	//如果新旧pod版本号一致则跳过
	if n.ResourceVersion == o.ResourceVersion {
		return
	}
	logger.With(zap.String("action","update")).Info(
		fmt.Sprintf("New pod created %s/%s\n",n.GetNamespace(),o.GetName()))
}

//更新回调函数二加入队列但是不做业务操作
func onUpdate2(oldOBJ,newOBJ interface{}) {
	var (
		key string
		err error
	)
	if key,err = cache.MetaNamespaceKeyFunc(newOBJ);err !=nil {
		runtime.HandleError(err)
		return
	}
	queue.Add(key)
}

//删除回调
func onDelete(objob interface{}) {
	pod ,ok := objob.(*corev1.Pod)
	if ! ok {
		logger.With(zap.String("action","delete")).Error("Unserializable pod")
		return
	}

	pod.GetNamespace()
	logger.With(zap.String("action","delete")).Info(fmt.Sprintf("New pod deleted %s/%s",
		pod.GetNamespace(),pod.GetName()))
}

//删除回调2
func onDelete2(objob interface{}) {
	var (
		key string
		err error
	)
	if key,err = cache.MetaNamespaceKeyFunc(objob);err !=nil {
		runtime.HandleError(err)
		return
	}
	queue.Add(key)
}

//队列的处理器
func processQueue(stopper chan struct{}) {
	/*
		在这里使用一个匿名函数来延迟队列
		通过这种方式在函数完成前都会取调用done函数
		否则不会通知队列消息已经被处理
	意义：
		每一个消息都是在一个匿名的函数中执行
		queue.forget(key) 删除一条消息
		queue.addratelimited(key) 失败消息重试
		queue.NumRequeues(key) 判断消息是新还是旧，返回消息处理的次数
	*/
	for {
		func() {
			var (
				key string
				ok bool
			)
			obj,shutdown := queue.Get()
			if shutdown {
				return
			}
			defer queue.Done(obj)
			if key,ok = obj.(string);!ok {
				queue.Forget(obj)
				runtime.HandleError(fmt.Errorf("key is not a string %#v",obj))
				return
			}
			//获取namespace信息
			namespace,name,err := cache.SplitMetaNamespaceKey(key)
			if err != nil {
				queue.Forget(key)
				runtime.HandleError(fmt.Errorf("impossible to split key: %s",key))
				return
			}
			//打印pod对象超时重试次数
			logger.With(zap.Int("queue_len",queue.Len())).
				With(zap.Int("num_retry",queue.NumRequeues(key))).
				Info(fmt.Sprintf("received key %s/%s",namespace,name))
			//实现这个函数
			syncHandler := func() error { return nil}

			if err := syncHandler();err != nil {
				//如果是临时错误可以请求该消息
				queue.AddRateLimited(key)
				return
			}
			//列表中清理
			queue.Forget(key)
		}()

	}

}