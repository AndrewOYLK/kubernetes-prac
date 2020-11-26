package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 200 << 20 // 8MiB

	r.POST("/upload", func(c *gin.Context) {
		// 从MutilForm中获取文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename, " : ", file.Size)

		// 直接读取文件内容
		f, _ := file.Open()
		bys, _ := ioutil.ReadAll(f)
		// fmt.Println("内容:", string(bys))

		// 保存文件
		c.SaveUploadedFile(file, "/opt/test")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", string(bys)),
		})
	})

	r.Run(":8080")
}
