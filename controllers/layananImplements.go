package controllers

import (
	"net/http"
	"smartlink/entities"
	"smartlink/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Layanan struct{}

func LayananControllers() LayananController {
	return Layanan{}
}

func (ctx Layanan) LayananLaundry(c *gin.Context) {
	idUser := GetIDUSer(c)
	layanan := entities.Layanan{}
	if err := c.ShouldBindJSON(&layanan); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Request tidak valid")
		return
	}
	resp, err := services.PostLayanan(GetDB(), idUser, layanan)
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"Message": "Cannot Login, Username or password is wrong",
			"Code":    strconv.Itoa(http.StatusBadGateway),
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
func (ctx Layanan) GetByIdLayanan(c *gin.Context) {
	id := c.Param("id")
	idLayanan, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Masukkan id")
		return
	}
	resp, err := services.GetLayananById(GetDB(), uint(idLayanan))
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Layanan Not Found",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
func (ctx Layanan) GetAllLayanan(c *gin.Context) {
	resp, err := services.GetAllLayanan(GetDB())
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Layanan Not Found",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
