```golang
ginRouter := gin.Default()

v1 := ginRouter.Group("/v1")
{
    v1.Handle("POST", "/test", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "data": "test",
        })
    })
}

server := &http.Server {
    Addr: ":8088",
    Handler: ginRouter,
}

handler := make(chan error)

go(func() {
    handler<-server.ListenAndServe()
})()

go(func() {
    // 信号监听
    notify := make(chan os.Signal)
    signal.Notify(notify, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
    // <-notify
    handler <- fmt.Errorf("%s", <-notify)
})()

go(func() {
    // 注册服务
})()

getHandler := <-handler
fmt.Println(getHanlder.Error())
// 反注册服务
err := server.Shutdown(ctx.Background())
if err != nil {
    log.Fatal(err)
}
```