package routes

import (
	"smartlink/controllers"
	"smartlink/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Routes(router *gin.Engine)
}

type RouterController struct {
	userService    controllers.UserController
	layananService controllers.LayananController
}

func NewUserController(userService controllers.UserController, layanan controllers.LayananController) RouterController {
	return RouterController{userService, layanan}
}
func (c *RouterController) Routes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/login", c.userService.Login)
		v1.POST("/signup", c.userService.SignUp)

		v1.Use(middleware.Authorization())
		v1.GET("/user", c.userService.FindById)
		v1.GET("/users", c.userService.FindAll)
		v1.DELETE("/user/:id", c.userService.Delete)
		v1.PUT("/user/:id", c.userService.Update)
	}
	v2 := router.Group("/v1")
	{
		v3 := v2.Group("/layanan")
		{
			v3.Use(middleware.Authorization())
			v3.POST("/laundry", c.layananService.LayananLaundry)
			v3.GET("/laundry/:id", c.layananService.GetByIdLayanan)
			v3.GET("/laundrys", c.layananService.GetAllLayanan)
		}
	}
}
