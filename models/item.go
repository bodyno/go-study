package models

import (
	"time"
	"github.com/bodyno/go-study/db"
)

type Item struct {
	ID uint `gorm:"primary_key" json:"id,omitempty"`
	CodeId uint `json:"code_id,omitempty" sql:"not null"`
	Name string `json:"name" sql:"not null"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type ItemModel struct {}

func (m ItemModel) Init() {
	db.GetDB().AutoMigrate(&Item{})
}

func (m ItemModel) Create() (item *Item) {
	item = &Item{
		CodeId: 2,
		Name: "Second Item",
	}
	db.GetDB().Create(&item)
	return item
}

func (m ItemModel) Find() (items []*Item) {
	db.GetDB().Table("items").Select("code_id, name").Scan(&items)
	return
}