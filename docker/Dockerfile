FROM golang:1.15 as builder
WORKDIR /root
COPY . ./
RUN export GO111MODULE=on GOPROXY="https://goproxy.cn,https://goproxy.io,direct" \
    && CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /root /root
WORKDIR /root
ENTRYPOINT ["./main"]