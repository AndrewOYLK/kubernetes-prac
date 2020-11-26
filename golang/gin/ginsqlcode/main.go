package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conStr := "root:123qwe@tcp(10.1.0.13:3306)/gsql"
	db, err := sql.Open("mysql", conStr) // 数据库实例
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// 【创建】数据库表
	//_, err = db.Exec("create table person(" +
	//	"id int auto_increment primary key," +
	//	"name varchar(12) not null," +
	//	"age int default 1" +
	//	");")
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//} else {
	//	fmt.Println("数据库表创建成功")
	//}

	// 【插入】数据到数据库表
	//_, err = db.Exec("insert into person(name, age)" +
	//	"values(?, ?);", "Davie", 18)
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//} else {
	//	fmt.Println("数据插入成功")
	//}

	// 【查询】数据表的数据
	rows, err := db.Query("select id, name, age from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	// 方法1:
	//for {
	//	if rows.Next() {
	//		//person := new(Person)
	//		var person Person
	//		if rows.Scan(&person.Id, &person.Name, &person.Age); err != nil {
	//			log.Fatal(err.Error())
	//			return
	//		}
	//		fmt.Println(person.Id, person.Name, person.Age)
	//	} else {
	//		break
	//	}
	//}
	// 方法2: 代码块
scan:
	if rows.Next() {
		//person := new(Person)
		var person Person
		if rows.Scan(&person.Id, &person.Name, &person.Age); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Id, person.Name, person.Age)
		goto scan
	}
}

type Person struct {
	Id   int
	Name string
	Age  int
}
