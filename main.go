package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/bodyno/go-study/db"
	"github.com/bodyno/go-study/models"
	"github.com/bodyno/go-study/middles"
	"flag"
	"github.com/bodyno/go-study/config"
	"github.com/bodyno/go-study/routers"
)

func main() {

	flag.IntVar(&config.Age, "age", 2, "age of gopher")
	flag.Parse()


	db.Init()
	models.Init()

	r := gin.Default()
	r.Use(middles.Recovery())

	routers.InitRouters(r)

	r.Run(":1234")
}