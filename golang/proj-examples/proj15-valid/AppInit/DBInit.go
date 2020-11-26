/*
	DB初始化
	一般可以光放文档抄
*/
package AppInit

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 数据库驱动引用
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(192.168.189.128:3306)/gmicro?charset=utf8mb4")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("===", db)

	db.DB().SetMaxIdleConns(10) // 最大连接池
	db.DB().SetMaxOpenConns(50) // 最大连接数
}

func GetDB() *gorm.DB {
	return db
}
