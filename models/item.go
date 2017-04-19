package models

import (
	"time"
	"nobody/go-study/db"
	"fmt"
)

type Item struct {
	ID uint `gorm:"primary_key" json:"id,omitempty"`
	CodeId uint `json:"code_id,omitempty"`
	Name string `json:"name"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ItemModel struct {}

func (m ItemModel) Init() {
	db.GetDB().AutoMigrate(&Item{})
}

func (m ItemModel) Create() (item Item) {
	item = Item{
		CodeId: 1,
		Name: "Item_One",
	}
	db.GetDB().Create(&item)
	return item
}

func (m ItemModel) Find() (result []struct{
	Item
	Product *Product `json:"product,omitempty"`
}) {

	db.GetDB().Table("items").Select("items.name").Joins("left join products on products.id = items.code_id").Scan(&result)
	fmt.Println(result)
	return result
}