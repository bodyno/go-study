package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"nobody/go-study/db"
	"nobody/go-study/models"
)

func main() {
	db.Init()
	productModel := new(models.ProductModel)
	productModel.Init()

	itemModel := new(models.ItemModel)
	itemModel.Init()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		list, count := productModel.Find()
		c.JSON(200, gin.H{
			"list": list,
			"count": count,
		})
	})

	r.GET("/item", func(c *gin.Context) {
		result := itemModel.Find()
		c.JSON(200, gin.H{
			"result": result,
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"Hello": "V1",
			})
		})
	}

	r.Run(":1234")
}