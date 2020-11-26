package Wrappers

import (
	"context"
	"fmt"
	"proj11-hystrix-wrapper/Services"
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

func defaultProds(rsp interface{}) {
	// 降级方法：尽可能简单，不容易出错，尽量不要有异常
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 4; i++ {
		models = append(models, newProd(200+i, "prodname"+strconv.Itoa(100+int(i))))
	}

	result := rsp.(*Services.ProdListResponse) // 进行断言
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
		Timeout: 5000, // 允许超时5秒
	}
	// 2. 配置command
	hystrix.ConfigureCommand(cmdName, configA)
	return hystrix.Do(cmdName, func() error {
		// 如果错误会进入降级方法
		return this.Client.Call(ctx, req, rsp)
	}, func(e error) error {
		// 降级方法
		fmt.Println("调用降级函数")
		defaultProds(rsp)
		return nil
	})
}

func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{
		c,
	}
}
