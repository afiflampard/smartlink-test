package controllers

import "github.com/gin-gonic/gin"

type LayananController interface {
	LayananLaundry(c *gin.Context)
	GetByIdLayanan(c *gin.Context)
	GetAllLayanan(c *gin.Context)
}
