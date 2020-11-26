package main

import (
	"html/template"
	"os"
)

func main() {
	// test1()
	test2()
}

func test1() {
	s := "Hello {{if .flag}}andrew {{else}}david {{end}}\n"

	tmpl := template.New("test")
	tmpl, err := tmpl.Parse(s)
	if err != nil {
		panic(err.Error())
	}

	err = tmpl.Execute(os.Stdout, map[string]bool{
		"flag": false,
	})
	if err != nil {
		panic(err.Error())
	}
}

func test2() {
	s := "Hello {{$name}}"

	name := "andrew"

	tmpl := template.New("test")
	tmpl, err := tmpl.Parse(s)
	if err != nil {
		panic(err.Error())
	}

	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		panic(err.Error())
	}
}
