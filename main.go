package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/bodyno/go-study/db"
	"github.com/bodyno/go-study/models"
	"github.com/bodyno/go-study/middles"
	"flag"
	"github.com/bodyno/go-study/config"
	"github.com/bodyno/go-study/routers"
	"github.com/go-bongo/bongo"
	"log"
	"fmt"
)

func main() {

	mongoConfig := &bongo.Config{
		ConnectionString: "mongodb://root:MGhYIjJURXQiS3e@118.178.240.112:3717/admin",
		Database: "mongo-test",
	}
	connection, err := bongo.Connect(mongoConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb连接成功")

	type Person struct {
		bongo.DocumentBase `bson:",inline"`
		FirstName string
		LastName string
		Gender string
	}

	myPerson := &Person{
		FirstName: "Testy",
		LastName: "McGee",
		Gender: "male",
	}
	connection.Collection("people").Save(myPerson)



	flag.IntVar(&config.Age, "age", 2, "age of gopher")
	flag.Parse()



	db.Init()
	models.Init()

	r := gin.Default()
	r.Use(middles.Recovery())

	routers.InitRouters(r)

	r.Run(":1234")
}