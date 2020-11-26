micro --registry consul --registry_address 192.168
.189.128:8500 list services


## 调用服务 方式1

micro --registry consul --registry_address 192.168
.189.128:8500 get service test.jtthink.com

micro --registry consul --registry_address 192.168
.189.128:8500 call test.jtthink.com TestService.Call "{\"id\": 3}"


## 调用服务 方式2

使用POSTMAN GET请求：http://localhost:8001

修改：
Headers
    Content-Type: application/json

提交：
Body
    使用json rpc的方式进行请求
    {
        "jsonrpc": "2.0",
        "method": "TestService.Call",
        "params": [{
            "id": 3
        }],
        "id": 1
    }

结果
    {
        "id": "1",
        "result": {
            "data": "test3"
        },
        error: null
    }

## 调用服务 方式3 - api gateway

# 备注因为新版的micro的
# 到我们实际环境的时候，就可以把网关用容器进行部署，部署成一个专门的网关

./apigw.sh
    export MICRO_REGISTRY=consul
    export MICRO_REGISTRY_ADDRESS=10.1.0.13:8500
    export MICRO_API_NAMESPACE=api.jtthink.com
    export MICRO_API_HANDLER=rpc # 可以是http
    micro api # 默认端口是8080

http://10.1.0.13:8080/test/TestService/call

POST：
{
	"id": 456
}

OUTPUT:
{
  "data": "test456"
}


---

## 借助gin框架的验证做法

https://github.com/go-playground/validator

go get gopkg.io/go-playground/validator.v9


-----------------------------------------------------
Models.proto
**这里存储与服务相关的model，可以和数据库的字段不一致**


------------------------------------------------------

重新安装项目涉及的库：
go get -u github.com/micro/go-micro
go get -u github.com/micro/go-plugins@master // 需要做replace操作

etcdctl get / --prefix --keys-only=true  -> 只看key

micro工具可以帮助我们方便查看etcd：`micro web` 启动一个内置的web服务


------------------------------------


jsonrpc协议