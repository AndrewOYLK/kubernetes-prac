package client

import (
	"context"
	"fmt"

	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

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
