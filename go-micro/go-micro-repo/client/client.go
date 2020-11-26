package main

import (
	"context"
	"fmt"
	"gomicro/message"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("student_client"),
	)
	service.Init()

	studentService := message.NewStudentServiceClient("student_service", service.Client())

	res, err := studentService.GetStudent(context.TODO(), &message.StudentRequest{Name: "davie"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	fmt.Println(res.Name)
	fmt.Println(res.Classes)
	fmt.Println(res.Grade)
}
