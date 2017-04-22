package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
	"nobody/go-study/models"
	"nobody/go-study/utils"
	"github.com/skip2/go-qrcode"
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

func QrCode(c *gin.Context) {
	png, _ := qrcode.Encode("https://bodyno.com", qrcode.High, 256)
	//c.Set("Content-Type", "image/png")
	c.Data(200, "image/png", png)

}

func test() {
	panic(utils.Unauthorized)
}

func ShowError(c *gin.Context) {
	test()
	//utils.ApiError{"NotFount", "没有", 404}.JSON(c)
}