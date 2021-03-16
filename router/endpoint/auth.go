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

type AuthRegister struct {
	User    model.User    `json:"user"`
	Company model.Company `json:"company"`
	Outlet  model.Outlet  `json:"outlet"`
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

func Register(c *gin.Context) {
	var obj AuthRegister
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Email exist or not
	result := config.DB.First(&obj.User, "email = ?", obj.User.Email)
	if result.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exist"})
		return
	}

	// Create company
	if result := config.DB.Create(&obj.Company); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	var companyId = obj.Company.ID
	password, err := util.HashPassword(obj.User.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	obj.User.Password = password
	obj.User.CompanyId = int(companyId)

	// Create User
	obj.User.IsActive = true
	obj.User.IsSuperUser = true
	obj.User.UserRolesId = 12
	if result := config.DB.Create(&obj.User); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	// Create Outlet
	obj.Outlet.Code = "WHS-001"
	obj.Outlet.Name = "Outlet 1"
	obj.Outlet.CompanyId = int(companyId)
	obj.Outlet.IsActive = true
	if result := config.DB.Create(&obj.Outlet); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, obj)
}
