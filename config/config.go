package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	viperUser := os.Getenv("DB_NAME")
	viperPassword := os.Getenv("DB_PASSWORD")
	viperDb := os.Getenv("DB_DATABASE")
	viperProtocol := os.Getenv("DB_PROTOCOL")
	viperHost := os.Getenv("DB_HOST")
	viperPort := os.Getenv("DB_PORT")

	prosgretConname := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viperUser,
		viperPassword,
		viperProtocol,
		viperHost,
		viperPort,
		viperDb)
	fmt.Println("conname is\t\t", prosgretConname)
	db, err := gorm.Open(mysql.Open(prosgretConname))
	if err != nil {

		panic(err)
	}
	return db
}
