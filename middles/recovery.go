package middles

import (
	"gopkg.in/gin-gonic/gin.v1"
	"nobody/go-study/utils"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case utils.ApiError:
					err.(utils.ApiError).JSON(c)
				}
				panic(err)
			}
		}()
		c.Next()
	}
}
