package main

import (
	"bufio"
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
	queques "k8s.io/client-go/util/workqueue"
	tools "k8s/auxiliaryway"
	"log"
	"os"
	"time"
)

//clientset只能获取到k8s内置的对象信息无法获取CRD信息
func main() {
	resconfig, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}
	//new一个对象
	clientinfo, err := kubernetes.NewForConfig(resconfig)
	if err != nil {
		panic(err)
	}
	//获取所有ns下的pod信息
	podlist := clientinfo.CoreV1().Pods(apiv1.NamespaceAll)
	//设置获取的对象最大数量
	pods, err := podlist.List(context.Background(), metav1.ListOptions{Limit: 100})
	if err != nil {
		panic(err)
	}
	for _, va := range pods.Items {
		fmt.Printf("NameSpace: %v \t Name: %v\t label: %+v\n", va.Namespace,
			va.Name, va.Labels)
	}

	// 创建Informer的监视对象
	factory := informers.NewSharedInformerFactory(clientinfo, 30*time.Second)
	// 获取事件的informer队列
	eventinfomer := factory.Core().V1().Pods().Informer()
	// 创建一个限速队列
	workquques := queques.NewRateLimitingQueue(queques.DefaultControllerRateLimiter())
	defer workquques.ShutDown()

	//构造监听函数
	eventinfomer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			//获取POd资源对象
			pod, ok := obj.(*corev1.Pod)
			if !ok {
				//fmt.Println("Create Unserializable pod")
				return
			}
			fmt.Printf("[Event] %v/%v is create successfully\n", pod.GetNamespace(), pod.GetName())
		},
		UpdateFunc: nil,
		DeleteFunc: func(obj interface{}) {
			pod, ok := obj.(*corev1.Pod)
			if !ok {
				//fmt.Println("Delete Unserializable pod")
				return
			}
			fmt.Printf("[Delete] %v/%v is Delete successfully\n", pod.GetNamespace(), pod.GetName())
		},
	})
	go eventinfomer.Run(wait.NeverStop)

	//设置namespace名称
	nsname := "test"
	//操作namespaceClient让他拥有对操作的方法
	namespaceClient := clientinfo.CoreV1().Namespaces()
	//构造对象
	namespae := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: nsname},
		Status:     apiv1.NamespaceStatus{Phase: apiv1.NamespaceActive},
	}
	//构造OPt的参数
	opt := metav1.CreateOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
	}
	//创建一个namespace
	fmt.Printf("开始创建NS对象：%s\n", nsname)
	nsobj, err := namespaceClient.Create(context.Background(), namespae, opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %s\tStatus: %s\tUID: %s\n", nsobj.Name, nsobj.Status, nsobj.UID)

	//创建目前ns下的deployment资源对象
	deploymentsClient := clientinfo.AppsV1().Deployments(nsname)

	//写入deployment的参数信息
	deployments := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{},
		//metadata 指定name和ns参数
		ObjectMeta: metav1.ObjectMeta{
			Name:      "demon-deployment",
			Namespace: nsname},
		Spec: appsv1.DeploymentSpec{
			//指定副本数量
			Replicas: tools.Int32Ptr(int32(2)),
			//指定标签
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "test-demo",
				},
			},
			//POD的模板
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					//定义标签选择器
					Labels: map[string]string{
						"app": "test-demo",
					},
					//定义pod-ns
					Namespace: nsname,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:      "nginx",
							Image:     "nginx",
							Resources: apiv1.ResourceRequirements{},
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	//创建deploy对象
	res, err := deploymentsClient.Create(context.Background(), deployments, metav1.CreateOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Success create Deployment", res.GetNamespace(), "/", res.GetObjectMeta().GetName())
	//等待手动触发
	prompt()

	//更新Deployment
	fmt.Println("Updating deployment...")

	//更新Deployment
	if retryErr := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		//获取deployment最新状态然后更新
		result, err := deploymentsClient.Get(context.TODO(), "demon-deployment", metav1.GetOptions{
			TypeMeta: metav1.TypeMeta{
				Kind: "Deployment",
			},
		})
		if err != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment %v", err))
		}
		//更改pod数量
		result.Spec.Replicas = tools.Int32Ptr(1)
		//更改镜像名称
		result.Spec.Template.Spec.Containers[0].Image = "nginx:1.14"
		//更新Deployment对象
		_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	}); retryErr != nil {
		panic(fmt.Errorf("Update Deployment Failed：%v", retryErr))
	}

	prompt()

	time.Sleep(3 * time.Second)
	//执行删除Deployment
	deploymentsClient.Delete(context.Background(), "demon-deployment", metav1.DeleteOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
	})
	//延迟10秒后执行删除
	time.Sleep(10 * time.Second)
	//删除指定的对象
	fmt.Printf("开始删除NS对象：%v\n", nsname)
	deleteObj := metav1.DeleteOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
	}
	err = namespaceClient.Delete(context.Background(), nsname, deleteObj)
	if err != nil {
		log.Fatal(err)
	}
	<-wait.NeverStop
}

func prompt() {
	fmt.Printf("-> Press Return key to Continue.\n")
	//从系统的标准输入接收
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
