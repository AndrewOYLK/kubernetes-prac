package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
}

func main() {
	// 1. map类型的实例
	tmp := map[string]string{
		"Name": "andrew",
	}
	fmt.Println(tmp)

	// 2. 结构体类型实例
	// p := Person{
	// 	Name: "andrew",
	// }

	s := "What is {{.Name}} doing {{-3}} hello world"
	// s := "What is {{- .Name -}} doing {{-3}} hello world"
	tmpl, err := template.New("test").Parse(s)
	if err != nil {
		panic(err.Error())
	}

	err = tmpl.Execute(os.Stdout, tmp)
	if err != nil {
		panic(err.Error())
	}
}
