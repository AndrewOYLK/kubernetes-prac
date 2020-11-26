package main

import (
	"context"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	"k8s/common"
)

func main() {
	var (
	//tailfLines int64
	)

	// 获取Container数据
	rest, _ := common.NewResAndLog()

	// 生成获取POD日志的请求
	req := rest.CoreV1().Pods("kube-system").GetLogs("kube-proxy-h4rmx", &coreV1.PodLogOptions{})

	//
	fmt.Println("开始获取log")

	// 发送请求
	res := req.Do(context.TODO())
	if res.Error() != nil {
		err := res.Error()
		fmt.Println(err)
		return
	}

	// 获取输出的结果
	logs, err := res.Raw()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 输出结果
	fmt.Println("Container Output：", string(logs), res.Error())

}
