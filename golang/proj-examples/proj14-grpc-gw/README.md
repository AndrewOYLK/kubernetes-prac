创建grpc网关

## 变成grpc接口
grpc.NewService()

## 创建grpc网关（grpc-gateway）

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

go get -u google.golang.org/grpc

## 添加访问规则