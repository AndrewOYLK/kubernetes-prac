package main

import (
	"os"
	"text/template"
)

// Inventory 1. 数据结构
type Inventory struct {
	Material string
	Count    uint
}

func main() {

	// 2. 准备数据
	sweaters := Inventory{
		"wool",
		17,
	}

	// 3. 准备模板
	tmpl := template.New("test")
	tmpl, err := tmpl.Parse("{{.Count}} items are made of {{.Material}}\n")
	if err != nil {
		panic(err.Error())
	}

	// 4. 模板渲染数据
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err.Error())
	}
}
