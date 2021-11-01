package helper

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Error interface{} `json:"message"`
}

func Error(c *gin.Context, code int, message interface{}) {
	payload := &ErrorResponse{
		Error: message,
	}
	c.JSON(code, payload)
}
