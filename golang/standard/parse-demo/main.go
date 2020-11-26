package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/json_post", func(c *gin.Context) {

		// buf := bytes.Buffer{}
		bs, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println(string(bs))

		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	router.Run(":8080")
}
