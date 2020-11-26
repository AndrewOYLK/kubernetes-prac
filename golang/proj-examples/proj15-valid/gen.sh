#!/bin/bash

# wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip
# go get -u github.com/golang/protobuf/protoc-gen-go --> 插件
# go get -u github.com/favadi/protoc-go-inject-tag

# 使用micro api网关功能，编译proto文件，需要生成micro文件。
# 编译生成该文件需要使用到一个新的protoc-gen-micro库
# go get -u github.com/micro/protoc-gen-micro

cd Services/protos
protoc --go_out=../ Models.proto
protoc --micro_out=../ --go_out=../ UserService.proto
protoc-go-inject-tag -input=../Models.pb.go
protoc-go-inject-tag -input=../UserService.pb.go
cd ..
cd .. 