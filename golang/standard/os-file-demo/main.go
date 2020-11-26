package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 1. 测试
	// fmt.Println(os.Getpagesize())

	// 2. 测试
	// f, err := os.Open("message")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(f.)

	// 3. 测试线程打开
	// p, err := os.FindProcess(107915)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(p)
	// p.Kill()
	// fmt.Println(p)

	// 4. 测试
	// f_info, err := os.Stat("message")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Printf("%#v\n", f_info)
	// fmt.Println(f_info)
	// fmt.Println(os.Hostname())
	// fmt.Println(os.Getwd())

	// 5. 测试
	ex := exec.Command("touch", "/opt/abcde.txt")
	fmt.Println(ex.Start())
}
