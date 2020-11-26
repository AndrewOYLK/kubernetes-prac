熔断器hystrix的使用

- 通用降级方法的编写姿势


```golang
func defaultProds(rsp interface{}) {
	// 降级方法：尽可能简单，不容易出错，尽量不要有异常
	models := make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 4; i++ {
		models = append(models, newProd(200+i, "prodname"+strconv.Itoa(100+int(i))))
	}

    /*
        因为这个方法接收的参数类型是interface{}，它自身没有任何属性和方法的接口实例，
        所以如果想通过它来调用属性和方法，那么就需要进行接口断言
    */
	result := rsp.(*Services.ProdListResponse) 
	result.Data = models
}
```

---

注意：hystrix
	熔断器(进行计算)+降级方法一般都是结合使用，它们都是针对rpc请求访问超时之后要做的一些处理

熔断器的参数设置 - 主要的三个参数

- RequestVolumeThreshold:5	-> 默认20。熔断器请求阈值，意思是有20个请求才进行错误百分比计算
- ErrorPercentThreshold:20 -> 错误百分比。默认50（50%）
- SleepWindow:5000	-> 过多长时间，熔断器再次检测是否开启，单位毫秒（默认5秒）