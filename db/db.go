package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/bodyno/go-study/config"
)

var instance *gorm.DB

func Init() {
	db, err := gorm.Open("mysql", config.Mysql)
	if err != nil {
		panic(err)
		panic("falied to connect database")
	}
	instance = db
}

func GetDB() *gorm.DB {
	return instance
}