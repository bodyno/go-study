package models

import (
	"github.com/bodyno/go-study/db"
	"fmt"
	"time"
	"sync"
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
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		db.GetDB().Limit(1).Select("code, price").Find(&product)
		wg.Done()
	}()
	go func() {
		db.GetDB().Model(&Product{}).Count(&count)
		wg.Done()
	}()
	wg.Wait()
	return product, count
}