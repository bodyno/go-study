package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var instance *gorm.DB

func Init() {
	db, err := gorm.Open("mysql", "baobaobooks_dev:Y4GPcRojd8wfZzIqMLD7VUle9OS356@(rm-bp1199ayx32ah1u80o.mysql.rds.aliyuncs.com:3306)/baobaobooks?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
		panic("falied to connect database")
	}
	instance = db
}

func GetDB() *gorm.DB {
	return instance
}