package services

import (
	"errors"
	"fmt"
	"net/http"
	"smartlink/entities"
	"smartlink/models"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func ValidationUnit(unit string) error {
	unit = strings.ToLower(unit)
	if unit == "kg" || unit == "pcs" || unit == "cm" || unit == "m2" {
		return nil
	}
	return errors.New("Error")
}

func PostLayanan(db *gorm.DB, userId string, userLayanan entities.Layanan) (*entities.LayananResponse, error) {
	err := ValidationUnit(userLayanan.Unit)
	if err != nil {
		return nil, err
	}
	var user models.User
	if err := db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	userLayanan.Harga = strings.Replace(userLayanan.Harga, ".", "", -1)
	tempHarga := strings.Replace(userLayanan.Harga, ",", ".", -1)
	harga, err := strconv.ParseFloat(tempHarga, 64)
	if err != nil {
		return nil, err
	}
	addLayanan := models.Layanan{
		Unit:   userLayanan.Unit,
		Nama:   userLayanan.Nama,
		Harga:  harga,
		UserId: user.ID,
	}
	if err := db.Create(&addLayanan).Error; err != nil {
		return nil, err
	}
	addLayanan.IdLayanan = "LYN" + fmt.Sprintf("%03d", addLayanan.ID)
	if err := db.Save(&addLayanan).Error; err != nil {
		return nil, err
	}
	return &entities.LayananResponse{
		Code:   http.StatusAccepted,
		Status: "success",
		Data: entities.DataResponse{
			ID:     addLayanan.IdLayanan,
			Nama:   addLayanan.Nama,
			Unit:   addLayanan.Unit,
			Harga:  fmt.Sprintf("%.2f", addLayanan.Harga),
			UserID: user.UserId,
		},
	}, nil
}

func GetLayananById(db *gorm.DB, id uint) (*models.Layanan, error) {
	var layanan models.Layanan
	if err := db.Preload("User").First(&layanan, id).Error; err != nil {
		return nil, err
	}
	return &layanan, nil
}

func GetAllLayanan(db *gorm.DB) (*[]models.Layanan, error) {
	var layanan []models.Layanan
	if err := db.Preload("User").Find(&layanan).Error; err != nil {
		return nil, err
	}
	return &layanan, nil
}
