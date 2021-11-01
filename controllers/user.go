package controllers

import "github.com/gin-gonic/gin"

type UserController interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
