package middleware

import (
	"fmt"
	"net/http"
	"smartlink/controllers"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Split(authHeader, " ")
		token, err := controllers.JWTAuthService().ValidateToken(tokenString[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			fmt.Println("Masuk sini")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user_id", claims["user_id"])
	}
}
