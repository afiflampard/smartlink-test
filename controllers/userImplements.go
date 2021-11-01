package controllers

import (
	"net/http"
	"smartlink/entities"
	"smartlink/models"
	"smartlink/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct{}

func GetIDUSer(c *gin.Context) string {
	return c.MustGet("user_id").(string)
}

func UserControllers() UserController {
	return User{}
}
func (ctx User) Login(c *gin.Context) {
	userLogin := &entities.Login{}
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Request tidak valid")
		return
	}
	resp, err := services.LoginUser(GetDB(), *userLogin)
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"Message": "Cannot Login, Username or password is wrong",
			"Code":    strconv.Itoa(http.StatusBadGateway),
		})
		return
	}

	c.JSON(http.StatusAccepted, resp)

}
func (ctx User) Delete(c *gin.Context) {
	id := c.Param("id")
	idUser, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Masukkan id")
		return
	}
	resp, err := services.DeleteUser(GetDB(), idUser)
	if err != nil {
		c.JSON(http.StatusBadGateway, resp)
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

func (ctx User) FindAll(c *gin.Context) {
	resp, err := services.GetAllUser(GetDB(), c)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "User Not Found",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
func (ctx User) FindById(c *gin.Context) {
	id := GetIDUSer(c)
	resp, err := services.GetUserByID(GetDB(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "User Not Found",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
func (ctx User) SignUp(c *gin.Context) {
	newUser := models.User{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, "User bad request")
	}
	resp, err := services.CreateUser(GetDB(), &newUser)
	if err != nil {
		c.JSON(http.StatusBadGateway, resp)
		return
	}
	c.JSON(200, entities.SignUpResponse{
		Kode:    http.StatusAccepted,
		Status:  "success",
		Message: "Berhasil terdaftar",
	})
}

func (ctx User) Update(c *gin.Context) {

	userId := GetIDUSer(c)
	updateUser := models.User{}
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, "Request not valid")
		return
	}
	resp, err := services.UpdateUser(GetDB(), userId, &updateUser)
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"message": "user cannot update",
		})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
