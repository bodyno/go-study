package routers

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/bodyno/go-study/controllers"
)

func InitRouters(r *gin.Engine) {

	r.GET("/", controllers.Root)

	r.GET("/main", controllers.Main)

	r.GET("/add", controllers.Add)

	r.GET("/item", controllers.GetItems)

	r.GET("/error", controllers.ShowError)

	r.GET("/error2", controllers.ShowError2)

	r.GET("/qrcode", controllers.QrCode)

	r.GET("/test", controllers.Test)

	v1 := r.Group("/v1")
	{
		v1.GET("/", controllers.Root)
	}
}
