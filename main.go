package main

import (
	"log"
	"smartlink/config"
	"smartlink/controllers"
	"smartlink/migration"
	"smartlink/routes"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db := config.Connect()
	migration.Migrate(db)
	controllers.InitiateDB(db)
	r := gin.Default()
	routes := routes.NewUserController(controllers.UserControllers(), controllers.LayananControllers())
	r.Use(cors.Default())
	routes.Routes(r)
	log.Fatal(r.Run(":8000"))
}
