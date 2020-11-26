package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"strings"
	"time"
)

func main() {
	//设置client-go的客户端
	rest,err := clientcmd.BuildConfigFromFlags("","/root/.kube/config")
	if err != nil {
		panic(err)
	}
	//new一个新的client对象
	clientsets,err := kubernetes.NewForConfig(rest)
	if err !=nil {
		panic(err)
	}
	//设置时间间隔
	resyncPeriod := 30 *time.Minute
	//新建事件资源的监听器（传递一个informedd对象并且设置时间周期）AddEventHandler当事件被触发时候会调用函数处理
	si := informers.NewSharedInformerFactory(clientsets,resyncPeriod)
	si.Core().V1().Events().Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: newEvent(clientsets),
		})
	si.Start(wait.NeverStop)
}

//新建一个事件对象
func newEvent(kbc *kubernetes.Clientset) func(obj interface{}) {
	return func(obj interface{}) {
		//获取事件对象的值
		event := obj.(*corev1.Event)
		//对各种情况下的事件进行判断和处理
		switch obj {
		case event.InvolvedObject.Kind != "Pod":
			return
		case event.Reason != "FailedMount":
			return
		case strings.Contains(event.Message,"MountVolume.WaitForAttach failed for volume")&&strings.Contains(event.Message,"devicePath is empty"):
			err := kbc.CoreV1().Pods(event.InvolvedObject.Namespace).
				Delete(context.Background(),event.InvolvedObject.Name,metav1.DeleteOptions{})
			if err != nil {
				log.Panic(err.Error())
			}
			log.Print(fmt.Sprintf("Pod %s/%s Deleted because effected by the devicePath issue", event.InvolvedObject.Namespace,event.InvolvedObject.Name))

		}


	}
}
