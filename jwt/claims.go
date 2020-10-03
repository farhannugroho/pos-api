package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Email     string `json:"email"`
	UserId    int    `json:"user_id"`
	CompanyId int    `json:"company_id"`
	jwt.StandardClaims
}

func GetClaims(c *gin.Context) *Claims {
	token := c.Request.Header.Get("Authorization")
	claims, _ := ParseToken(token)
	return claims
}
