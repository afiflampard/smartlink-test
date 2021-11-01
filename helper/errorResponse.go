package helper

import "github.com/gin-gonic/gin"

type Response struct {
	Kode    uint
	Message string
	Status  bool
}

func Responses(c *gin.Context, code uint, message string, status bool) {
	payload := &Response{
		Kode:    code,
		Message: message,
		Status:  status,
	}
	c.JSON(int(code), payload)
}
