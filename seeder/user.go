package seeder

import (
	"log"
	"smartlink/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedUser() []models.User {
	pass, err := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	var users = []models.User{
		{
			UserId:   "USR001",
			Username: "fif",
			Password: string(pass),
			Nama:     "Afif Musyayyidin",
			Telepon:  "081615962254",
		},
		{
			UserId:   "USR002",
			Username: "fifa",
			Password: string(pass),
			Nama:     "Afif MM",
			Telepon:  "081615962254",
		},
	}
	return users

}
