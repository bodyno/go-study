package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/bodyno/go-study/db"
	"github.com/bodyno/go-study/models"
	"github.com/bodyno/go-study/middles"
	"github.com/bodyno/go-study/routers"
)

func main() {

	db.Init()
	models.Init()

	r := gin.Default()
	r.Use(middles.Recovery())

	routers.InitRouters(r)

	r.Run(":1234")
}