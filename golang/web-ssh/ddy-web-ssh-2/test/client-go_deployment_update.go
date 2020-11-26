package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/apimachinery/pkg/watch"
	common "k8s/common"
	"strconv"
	"strings"
	"time"
)

var (
	Path string
)

func init() {
	flag.StringVar(&Path, "path", "", "<request>必须项目")
}

func main() {
	var (
		deploymentYaml, depliymentJson []byte
		err                            error
		deploymentstruct               = &appsv1.Deployment{}
		k8sDeployment                  *appsv1.Deployment
	)
	flag.Parse()
	// 初始化日志和客户端
	clientset, logger := common.NewResAndLog()

	// 注册一个新的Deployment的监听器
	Deploymentwatch, err := clientset.AppsV1().Deployments("").Watch(context.TODO(), metaV1.ListOptions{})
	clientset.AppsV1().RESTClient()
	// 打印事件日志
	go func(d watch.Interface) {
		defer d.Stop()
		for {
			select {
			case event := <-d.ResultChan():
				ob := event.Object.(*appsv1.Deployment)
				if ob.Status.AvailableReplicas != *(ob.Spec.Replicas) && strings.Contains(ob.Name, "nginx") {
					logger.Info("Deployment Event sync", zap.String("Event", fmt.Sprintf("%v/%v runningpod: %v/%v", ob.Namespace, ob.Name, *(ob.Spec.Replicas), ob.Status.AvailableReplicas)))
				}
			default:
				time.Sleep(time.Microsecond * 1000)
			}
		}
	}(Deploymentwatch)

	// 读取yaml配置文件
	if deploymentYaml, err = ioutil.ReadFile(Path); err != nil {
		panic(err)
	}

	// yaml转换为Json
	if depliymentJson, err = yaml.ToJSON(deploymentYaml); err != nil {
		panic(err)
	}

	// Json转换结构体
	if err := json.Unmarshal(depliymentJson, deploymentstruct); err != nil {
		logger.Error("Josn Unmarshal is Failed", zap.String("err", err.Error()))
	}

	// 给pod新增label
	deploymentstruct.Spec.Template.Labels["deply_time"] = strconv.Itoa(int(time.Now().Unix()))

	fmt.Println(deploymentstruct.Spec.Template.Labels)

	// 更新Deployment
	if _, err := clientset.AppsV1().Deployments("default").Update(context.TODO(), deploymentstruct, metaV1.UpdateOptions{}); err != nil {
		logger.Warn("Update Deployment is Failed", zap.String("ERROR", err.Error()))
	} else {
		fmt.Println("更新完成！！")
	}

	// 等待更新完成
	for {
		// 获取K8s中Deployment对象的当前状态
		if k8sDeployment, err = clientset.AppsV1().Deployments("default").Get(
			context.TODO(),
			deploymentstruct.Name,
			metaV1.GetOptions{}); err != nil {
			goto RETRY
		}

		/*			Deployment状态的判定
					1.当前的运行中的deployment的replic更新数量是否等于当前的定义的replica的数量
					2.当前的运行中的Depolyment副本数量是否等于资源列表清单定义的数量
					3.当前的运行中的Depolyment可用副本数量是否等于定义的数量
					4.replica controller的生成器是否等于已经设置的值
		*/
		if k8sDeployment.Status.UpdatedReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.Replicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.AvailableReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.ObservedGeneration == k8sDeployment.Generation {
			// 升级完成则跳出循环
			break
			fmt.Println("部署成功！")
		}

		// 打印运行中的Pod比例
		fmt.Printf("正在更新中的对象：(%d/%d)\n", k8sDeployment.Status.AvailableReplicas, *(k8sDeployment.Spec.Replicas))

	RETRY:
		time.Sleep(1 * time.Second)
	}

	//　打印每一个POD的当前状态（最终展示出pod列表）
	if podList, err := clientset.CoreV1().Pods("default").List(context.TODO(), metaV1.ListOptions{
		LabelSelector: "app=nginx",
	}); err != nil {
		for _, pod := range podList.Items {
			// 获取POD名字和状态
			podName := pod.Name
			podStatus := string(pod.Status.Phase)

			// running 表示至少
			if podStatus == string(corev1.PodRunning) {
				// 汇总错误的原因
				if pod.Status.Reason != "" {
					podStatus = pod.Status.Reason
					goto KO
				}

				// condition有错误信息
				for _, cond := range pod.Status.Conditions {
					if cond.Type == corev1.PodReady {
						if cond.Status == corev1.ConditionTrue {
							podStatus = cond.Reason
						}
						goto KO
					}
				}

				// 如果没有read condition则设置状态为UNKNOWN
				podStatus = "Unknown"
			}
		KO:
			fmt.Printf("[name=%s status=%s]\n", podName, podStatus)
		}
	}
	time.Sleep(100 * time.Second)
}
