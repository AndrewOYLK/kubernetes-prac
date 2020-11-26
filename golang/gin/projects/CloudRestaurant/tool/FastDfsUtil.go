package tool

import (
	"bufio"
	"fmt"
	"github.com/tedcy/fdfs_client"
	"os"
	"strings"
)

/*
 上传文件到fastDFS系统
*/
func UploadFile(filename string) string {
	client, err := fdfs_client.NewClientWithConfig("./config/fastdfs.conf")
	defer client.Destory()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	fileId, err := client.UploadByFilename(filename)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return fileId
}

/**
从配置文件中读取文件服务器的IP和端口相关配置
*/
func FileServerAddr() string {
	file, err := os.Open("./config/fastdfs.conf")
	if err != nil {
		return ""
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		str := strings.SplitN(line, "=", 2)
		//fmt.Println(str[0], str[1])
		switch str[0] {
		case "http_server_port":
			return str[1]
		}
		if err != nil {
			return ""
		}
	}
}
