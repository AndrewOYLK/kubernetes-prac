package common

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	ClientSet *kubernetes.Clientset
	logger    *zap.Logger
)

func NewResAndLog() (*kubernetes.Clientset, *zap.Logger) {

	//指定K8s-config文件
	rest, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}
	ClientSet, err = kubernetes.NewForConfig(rest)
	if err != nil {
		panic(err)
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return ClientSet, logger
}
