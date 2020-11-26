整合hystrix到框架中，使用到wrapper


熔断机制目前感觉是配置在请求端的。

                                          ------ grpc服务后端1
浏览器 ------ restful api（熔断配置在这里）  ------ grpc服务后端2
                                          ------ grpc服务后端3