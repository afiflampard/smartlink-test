package models

import "gorm.io/gorm"

type Layanan struct {
	gorm.Model
	IdLayanan string  `gorm:"column:layanan_id" json:"layanan_id"`
	Nama      string  `gorm:"column:nama" json:"nama"`
	Unit      string  `gorm:"column:unit" json:"unit"`
	Harga     float64 `gorm:"column:harga" json:"harga"`
	UserId    uint    `gorm:"column:user_id" json:"user_id"`
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
