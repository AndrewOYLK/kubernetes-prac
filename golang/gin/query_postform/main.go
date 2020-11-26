package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/post/:k", func(c *gin.Context) {
		// 第一种方式
		para := c.Param("k")
		fmt.Println(para)

		// 第二种方式
		id := c.Query("id")
		page := c.DefaultQuery("page", "10")

		// 第三种方式
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	r.Run(":8080")
}
