package services

import (
	"smartlink/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user *models.User) (string, error) {
	expTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		ID: user.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("I love Malang"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
