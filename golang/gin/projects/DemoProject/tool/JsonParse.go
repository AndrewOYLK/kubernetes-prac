package tool

import (
	"encoding/json"
	"io"
)

type JsonParse struct {
}

func Decode(io io.Reader, v interface{}) error {
//func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}
