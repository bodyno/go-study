package models

import (
	"nobody/go-study/db"
	"fmt"
	"time"
)

type Product struct {
	ID uint `gorm:"primary_key" json:"id,omitempty"`
	Code string `gorm:"index:idx_code_price" json:"code,omitempty"`
	Price uint `gorm:"index:idx_code_price" json:"price,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ProductModel struct {}

func (m ProductModel) Init() {
	db.GetDB().AutoMigrate(&Product{})
}

func getCount(ch chan int) {
	var count int

	fmt.Println(count)
	ch <- count
}

func (m ProductModel) Find() (product []Product, count int) {
	db.GetDB().Create(&Product{Code: "1111", Price: 100})
	ch := make(chan bool, 10)
	go func() {
		db.GetDB().Limit(1).Select("code, price").Find(&product)
		fmt.Println(product)
		ch <- true
	}()
	go func() {
		db.GetDB().Model(&Product{}).Count(&count)
		ch <- true
	}()
	<-ch
	<-ch
	return product, count
}