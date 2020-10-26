package main

import "github.com/AndrewOYLK/k8scode/controller"

func main() {
	//conf := kube.GetConfigFromFile()
	//fmt.Println("==", conf.Insecure)

	// 20201025
	//client.TestClientSet()
	//client.TestInformer()

	// controller
	controller.Test()
}
