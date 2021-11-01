package migration

import (
	"log"
	"smartlink/models"
	"smartlink/seeder"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	tableExist := (db.Migrator().HasTable(&models.User{}) && db.Migrator().HasTable(&models.Layanan{}))
	if !tableExist {
		dbMigrate := db.Debug().Migrator().DropTable(&models.User{}, &models.Layanan{})
		if dbMigrate != nil {
			log.Fatal("Cannot Drop Table")
		}
		db.AutoMigrate(&models.User{}, &models.Layanan{})
		users := seeder.SeedUser()

		for _, user := range users {
			if err := db.Debug().Create(&user).Error; err != nil {
				log.Fatal("Failed to create User")
			}
		}
	}
}
