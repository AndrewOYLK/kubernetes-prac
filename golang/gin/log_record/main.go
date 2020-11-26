package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	route.Run(":8080")
}
