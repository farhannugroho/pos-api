package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"pos_api/config"
	"pos_api/model"
	"strings"
	"time"
)

var jwtSecret []byte

func GenerateToken(user model.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	id := uuid.NewV4()

	claims := Claims{
		user.Email,
		int(user.ID),
		user.CompanyId,
		jwt.StandardClaims{Id: id.String(),
			ExpiresAt: expireTime.Unix(),
			Issuer:    "pos_api",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(bearerToken string) (*Claims, error) {
	// remove bearer prefix
	token, _ := stripBearerPrefixFromTokenString(bearerToken)

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

func Setup() {
	jwtSecret = []byte(config.Config.JwtSecret)
}
