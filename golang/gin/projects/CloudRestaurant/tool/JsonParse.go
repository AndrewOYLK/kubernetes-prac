package tool

import (
	"encoding/json"
	"io"
)

/*
	这个文件的作用:
		提供一些方法来完成参数的解析（结构体的解析）
*/

type JsonParse struct {
}

func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
