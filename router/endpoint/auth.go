package endpoint

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pos_api/config"
	"pos_api/jwt"
	"pos_api/model"
	"pos_api/util"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var obj Auth
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Record Not Found
	var user model.User
	result := config.DB.Where("email = ?", obj.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "Email Not Found"})
		return
	}

	isMatch := util.CheckPasswordHash(obj.Password, user.Password)
	if isMatch != true {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong Password"})
		return
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
