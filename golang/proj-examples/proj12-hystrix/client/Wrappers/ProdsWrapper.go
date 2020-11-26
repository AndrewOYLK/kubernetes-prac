package Wrappers

import (
	"context"
	"fmt"
	"proj12-hystrix/client/Services"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)

func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

// 通用降级方法
func defaultData(rsp interface{}) {
	// 类型判断
	switch t := rsp.(type) {
	case *Services.ProdListResponse:
		defaultProds(rsp)
	case *Services.ProdDetailResponse:
		t.Data = newProd(10, "降级商品")
	}
}

// 商品列表降级方法
func defaultProds(rsp interface{}) {
	// 降级方法：尽可能简单，不容易出错，尽量不要有异常
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 4; i++ {
		models = append(models, newProd(200+i, "prodname"+strconv.Itoa(100+int(i))))
	}

	result := rsp.(*Services.ProdListResponse) // 进行断言类型
	result.Data = models
}

type ProdsWrapper struct {
	client.Client
}

func (this *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	// 同步方法
	// 服务名 + "." + 方法名
	cmdName := req.Service() + "." + req.Endpoint()

	// 这里为了体现降级
	// 1. 配置config
	configA := hystrix.CommandConfig{
		Timeout:                2000, // 允许超时2秒，用户每次访问依然需要等待2秒！前端响应比较慢（重视）
		RequestVolumeThreshold: 2,    // 2个请求里面只要有1个出错（50%）就熔断器就需要打开
		ErrorPercentThreshold:  50,
		SleepWindow:            5000, // 5秒过后再次去探测方法是否ok，如果还是不行，熔断器继续打开
	}
	// 2. 配置command
	hystrix.ConfigureCommand(cmdName, configA)
	return hystrix.Do(cmdName, func() error {
		// 如果错误会进入降级方法
		return this.Client.Call(ctx, req, rsp)
	}, func(e error) error {
		// 降级方法
		fmt.Println("调用降级函数")
		// defaultProds(rsp)
		defaultData(rsp)
		return nil
	})
}

func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{
		c,
	}
}
