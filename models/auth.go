package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID string `json:"user_id"`
	jwt.StandardClaims
}
