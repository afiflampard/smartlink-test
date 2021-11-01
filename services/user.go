package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"smartlink/entities"
	"smartlink/helper"
	"smartlink/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB, user entities.Login) (*entities.LoginResponse, error) {
	var userTemp models.User
	if err := db.Where("username = ?", user.Username).First(&userTemp).Error; err != nil {
		return nil, err
	}
	result := bcrypt.CompareHashAndPassword([]byte(userTemp.Password), []byte(user.Password))
	if result != nil {
		return nil, result
	}
	token, err := GenerateToken(&userTemp)
	if err != nil {
		return nil, err
	}
	return &entities.LoginResponse{
		Code:   http.StatusAccepted,
		Status: "success",
		Data: entities.Data{
			ID:       userTemp.UserId,
			Nama:     userTemp.Nama,
			Username: userTemp.Username,
			Token:    token,
		},
	}, nil
}

var regex, _ = regexp.Compile(`([A-Za-z]+[\d@]+[\w@]*|[\d@]+[A-Za-z]+[\w@]*)`)
var regexBilangan, _ = regexp.Compile(`^[0-9]*$`)

func Validation(user *models.User) (map[string]string, error) {
	fmt.Println(regexBilangan.MatchString(user.Telepon))
	if len(user.Nama) > 50 {
		return map[string]string{
			"Message": "Nama tidak boleh lebih dari 50 karakter",
		}, errors.New("Error")
	}

	if !regex.MatchString(user.Username) {
		return map[string]string{
			"Message": "Username harus terdiri dari huruf dan angka",
		}, errors.New("Error")
	}

	if len(user.Telepon) > 15 || !regexBilangan.MatchString(user.Telepon) {
		return map[string]string{
			"Message": "Nomer telepon tidak boleh diatas 15 karakter dan harus angka",
		}, errors.New("Error")
	}
	return nil, nil
}

func CreateUser(db *gorm.DB, user *models.User) (*helper.Response, error) {
	resp, err := Validation(user)
	if err != nil {
		return &helper.Response{
			Kode:    http.StatusBadRequest,
			Message: resp["Message"],
		}, err
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	addPerson := models.User{
		Username: user.Username,
		Password: string(pass),
		Nama:     user.Nama,
		Telepon:  user.Telepon,
	}
	if err := db.Create(&addPerson).Error; err != nil {
		return &helper.Response{
			Kode:    http.StatusBadRequest,
			Message: "Cannot add User",
			Status:  false,
		}, err
	}
	addPerson.UserId = "USR" + fmt.Sprintf("%03d", addPerson.ID)
	if err := db.Save(&addPerson).Error; err != nil {
		return &helper.Response{
			Kode:    http.StatusBadRequest,
			Message: "Cannot add ID",
			Status:  false,
		}, err
	}

	return &helper.Response{
		Kode:    http.StatusAccepted,
		Message: "User Success created",
		Status:  true,
	}, nil
}

func GetUserByID(db *gorm.DB, id string) (*models.User, error) {
	var user *models.User
	if err := db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUser(db *gorm.DB, c *gin.Context) (user []models.User, err error) {
	if err := db.Find(&user).Error; err != nil {
		helper.Responses(c, http.StatusBadRequest, "User Not Found", false)
	}
	return user, nil
}

func UpdateUser(db *gorm.DB, id string, user *models.User) (*models.User, error) {
	_, err := Validation(user)
	if err != nil {
		return nil, err
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	var updateUser models.User
	if err := db.Where("user_id = ?", id).First(&updateUser).Error; err != nil {
		return nil, err
	}
	if user.Username != "" {
		updateUser.Username = user.Username
		updateUser.Password = string(pass)
		// updateUser.RoleID = user.RoleID
		db.Save(&updateUser)
	}
	return &updateUser, nil
}

func DeleteUser(db *gorm.DB, id int) (map[string]string, error) {
	var user models.User
	if err := db.Delete(&user, id).Error; err != nil {
		return map[string]string{
			"message": "User tidak ada",
		}, err
	}
	return map[string]string{
		"message": "User telah terhapus",
	}, nil
}
