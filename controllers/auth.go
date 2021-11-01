package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTServices interface {
	GenerateToken(c *gin.Context, ID string, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
