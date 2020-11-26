package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// 迁移schema
	db.AutoMigrate(&Product{})

	// 创建数据
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 查询
	var product Product
	db.First(&product, 1)                   // 查询id为1的product记录
	db.First(&product, "code = ?", "L1212") // 查询code为L1212的product记录

	fmt.Printf("查询：%v \n", product)

	// 更新 - 更新product记录的price为2000
	db.Model(&product).Update("Price", 20000)

	fmt.Printf("更新完毕：%v \n", product)

	// 删除 - 删除product记录
	db.Delete(&product)
}
