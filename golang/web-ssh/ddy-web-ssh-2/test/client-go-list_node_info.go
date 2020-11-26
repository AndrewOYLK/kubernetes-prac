package main

import (
	"context"
	//"fmt"
	"github.com/modood/table"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"time"
)

type NodeInfo struct {
	NodeName      string    `table:"Node节点名称"`
	NodeStatus    string    `table:"Node节点状态"`
	NodeVersion   string    `table:"内核版本"`
	NodeOSVersion string    `table:"Node节点系统版本"`
	K8sVersion    string    `table:"K8s当前版本"`
	OsType        string    `table:"操作系统类型"`
	CreateTime    time.Time `table:"创建时间"`
}

type Namespace struct {
	Name       string    `table:"NameSpace名称"`
	CreateTime time.Time `table:"创建时间"`
	Status     string    `table:"当前运行状态"`
}

type ServiceTable struct {
	Name        string      `table:"Service名称"`
	ServiceType string      `table:"Service类型"`
	ClusterIP   string      `table:"ClusterIP"`
	PortType    string      `table:"端口类型"`
	Port        interface{} `table:"映射出的端口"`
	TargetPort  interface{} `table:"目标端口"`
	NodePort    interface{} `table:"暴露的NodePort"`
}

var loggers *zap.Logger

func init() {
	//初始化日志
	loggers, _ = zap.NewProduction()
}

func main() {
	defer loggers.Sync()
	loggers.Info("The K8s logger started")
	// 初始化config配置信息
	resconf, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		loggers.Panic(err.Error())
		os.Exit(1)
	}

	// 输出读取到的K8s对象信息
	loggers.Info("Kubernetes Target Info",
		zap.String("HOST", resconf.Host),
		zap.String("Username", resconf.Username),
		zap.String("APIPath", resconf.APIPath))

	// 创建需要的clientset对象
	clientset, err := kubernetes.NewForConfig(resconf)
	if err != nil {
		loggers.Panic(err.Error())
		os.Exit(2)
	}

	// 设置Node clientset
	nodelist, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		loggers.Error(err.Error())
	}

	// watch类型的clientset设置
	nodewatch, _ := clientset.CoreV1().Nodes().Watch(context.TODO(), metav1.ListOptions{})
	nodewatch.ResultChan()

	// 构造Node结构体
	nodetable := []NodeInfo{}

	// 打印node信息
	for _, node := range nodelist.Items {
		//fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		//	node.Name,
		//	node.Status.Phase,
		//	node.Status.Addresses,
		//	node.Status.NodeInfo.OSImage,
		//	node.Status.NodeInfo.KubeletVersion,
		//	node.Status.NodeInfo.OperatingSystem,
		//	node.Status.NodeInfo.Architecture,
		//	node.CreationTimestamp,
		//)
		nodetable = append(nodetable, NodeInfo{
			NodeName:      node.Name,
			NodeStatus:    string(node.Status.Conditions[4].Status),
			NodeVersion:   node.Status.NodeInfo.KernelVersion,
			NodeOSVersion: node.Status.NodeInfo.OSImage,
			K8sVersion:    node.Status.NodeInfo.KubeletVersion,
			OsType:        node.Status.NodeInfo.Architecture,
			CreateTime:    node.CreationTimestamp.Time,
		})
	}

	//输出表格
	table.Output(nodetable)
	table.Table(nodetable)

	// 构造ns结构体
	nsstruck := []Namespace{}

	// 设置namespace clientset
	namespaceClient, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		loggers.Error(err.Error())
	}

	// 打印namespace信息
	for _, ns := range namespaceClient.Items {
		//fmt.Printf("NameSpace 名字：%s \t创建时间：%v \t当前状态：%v\n", ns.Name, ns.CreationTimestamp, ns.Status.Phase)
		nsstruck = append(nsstruck, Namespace{
			Name:       ns.Name,
			CreateTime: ns.CreationTimestamp.Time,
			Status:     string(ns.Status.Phase),
		})
	}

	//输出表格
	table.Output(nsstruck)
	table.Table(nsstruck)

	// 设置service client
	serverClient, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		loggers.Error(err.Error())
	}

	// 初始化结构体
	servictab := []ServiceTable{}

	// 打印service信息
	for _, serv := range serverClient.Items {
		//fmt.Printf("Service 名称：%s/%s \t类型：%v \tClusterIP: %v \tPort: %v\n", serv.Namespace, serv.Name, serv.Spec.Type, serv.Spec.ClusterIP, serv.Spec.Ports)
		servictab = append(servictab, ServiceTable{
			Name:        serv.Namespace + "/" + serv.Name,
			ServiceType: string(serv.Spec.Type),
			ClusterIP:   serv.Spec.ClusterIP,
			PortType:    string(serv.Spec.Ports[0].Protocol),
			Port:        serv.Spec.Ports[0].Port,
			TargetPort:  serv.Spec.Ports[0].TargetPort.IntVal,
			NodePort:    serv.Spec.Ports[0].NodePort,
		})
	}

	//输出表格
	table.Output(servictab)
	table.Table(servictab)

	loggers.Info("In The End")
}
