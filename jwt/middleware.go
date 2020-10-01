package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Header struct {
	Authorization string `json:"authorization"`
}

func Middleware(c *gin.Context) {
	var msg = "OK"
	header := Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		msg = err.Error()
	}

	var bearerToken = header.Authorization
	token, err := stripBearerPrefixFromTokenString(bearerToken)
	if err != nil {
		msg = err.Error()
	}

	if token == "" {
		msg = "Token Invalid"
	} else {
		_, err := ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				msg = "Token Expired"
			default:
				msg = "Token Error"
			}
		}
	}

	if msg != "OK" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": msg})
		c.Abort()
		return
	}

	c.Next()
}
