package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"nobody/go-study/routers"
	"nobody/go-study/db"
	"nobody/go-study/models"
	"nobody/go-study/middles"
)

func main() {

	db.Init()
	models.Init()

	r := gin.Default()
	r.Use(middles.Recovery())

	routers.InitRouters(r)

	r.Run(":1234")
}