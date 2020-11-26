package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			fmt.Println(err)
			c.Status(http.StatusServiceUnavailable)
			return
		}

		// ===========================================
		// 保存文件
		reqURL := response.Request.URL.RequestURI()
		_, filename := path.Split(reqURL)
		newFile, _ := os.Create("/opt/" + filename)
		go saveFile(newFile, response.Body)
		fmt.Println("后台保存中")
		// ===========================================

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	router.Run(":8080")
}

func saveFile(dst *os.File, src io.ReadCloser) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println(err)
	}
}
