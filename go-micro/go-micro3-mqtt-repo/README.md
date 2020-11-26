```shell
# 终端执行命令

# 1. 在10.1.0.13机器上以开发模式运行consul，用于测试
consul agent -dev -client=0.0.0.0

# 2. 运行服务端进行服务注册
go run main.go --registry=consul

# 输出：注意每次启动服务，端口都不一样
# 2020-04-02 14:17:01.539301 I | Transport [http] Listening on [::]:35210
# 2020-04-02 14:17:01.539522 I | Broker [http] Connected to [::]:18210
# 2020-04-02 14:17:01.539908 I | Registry [consul] Registering node: student_service-4052cde6-7b4c-4813-8ae8-73603ef2acb0
```

```shell
# 终端执行命令

# 该命令表示我们每间隔5秒钟向服务注册组件注册一次，每次有效期限是10秒
go run main.go --register_ttl=10 --register_interval=5
```