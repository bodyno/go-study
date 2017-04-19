package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
	"nobody/go-study/models"
	"nobody/go-study/utils"
	"fmt"
)

func Root(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "world",
	})
}

func Add(c *gin.Context) {
	list, count := new(models.ProductModel).Find()
	c.JSON(200, gin.H{
		"list": list,
		"count": count,
	})
}

func GetItems(c *gin.Context) {
	result := new(models.ItemModel).Find()
	c.JSON(200, gin.H{
		"result": result,
	})
}

func test() {
	panic(utils.Unauthorized)
}

func ShowError(c *gin.Context) {
	test()
	//utils.ApiError{"NotFount", "没有", 404}.JSON(c)
}