package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/bodyno/go-study/models"
	"github.com/bodyno/go-study/utils"
	"github.com/skip2/go-qrcode"
	"os"
	"time"
)

func Root(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "world",
	})
}

func Add(c *gin.Context) {
	item := models.ItemModel{}.Create()
	c.JSON(200, item)
}

func GetItems(c *gin.Context) {
	result := models.ItemModel{}.Find()
	c.JSON(200, result)
}

func QrCode(c *gin.Context) {
	png, _ := qrcode.Encode("https://bodyno.com", qrcode.High, 256)
	c.Data(200, "image/png", png)

}

func test() {
	panic(utils.Unauthorized)
}

func ShowError(c *gin.Context) {
	test()
	//utils.ApiError{"NotFount", "没有", 404}.JSON(c)
}

func ShowError2(c *gin.Context) {
	os.Exit(1)
}

func Test(c *gin.Context) {
	time.Sleep(3 * time.Second)
	c.JSON(200, gin.H{
		"hello": "world",
	})
}