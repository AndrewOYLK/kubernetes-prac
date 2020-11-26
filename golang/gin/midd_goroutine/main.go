package main

/*
在中间件或handler中启动新的goroutine时，不能使用原始的上下文，必须使用只读副本
*/

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path " + c.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)

		log.Println("Done! in path " + c.Request.URL.Path)
	})

	r.Run(":8080")
}
