package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId   string `gorm:"column:user_id" json:"user_id"`
	Username string `gorm:"column:username; type:varchar(255); not null; unique" json:"username"`
	Password string `gorm:"column:password; type:varchar(255); not null" json:"password"`
	Nama     string `gorm:"column:nama" json:"nama"`
	Telepon  string `gorm:"column:telepon" json:"telepon"`
}
