package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var instance *gorm.DB

func Init() {
	db, err := gorm.Open("mysql", "root:Xuzeyu123@(rm-wz9643u6cranpzqm9o.mysql.rds.aliyuncs.com:3306)/nobody?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		panic("falied to connect database")
	}
	instance = db
}

func GetDB() *gorm.DB {
	return instance
}