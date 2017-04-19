package utils

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type ApiError struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Status int `json:"-"`
}

var (
	Unauthorized = ApiError{
		"Unauthorized",
		"权限不足",
		401,
	}
)

func (apiError ApiError) JSON (c *gin.Context) {
	c.JSON(apiError.Status, apiError)
}