package models

import (
	"testing"
	"github.com/bodyno/go-study/db"
)

func TestInit(t *testing.T) {
	db.Init()
}

func TestItemModel_Create(t *testing.T) {
	item := ItemModel{}.Create()
	if item == nil {
		t.Error("expect item is not nil")
	}
	if item.Name != "Second Item" {
		t.Error("expect item result name equal Second Item")
	}
}

func TestItemModel_Find(t *testing.T) {
	items := ItemModel{}.Find()
	if len(items) == 0 {
		t.Error("expect find results length not equal 0")
	}
}