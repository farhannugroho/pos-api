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

func GetAllUsers(c *gin.Context) {
	var list []model.User
	companyId := jwt.GetClaims(c).CompanyId
	config.DB.Where("company_id = ?", companyId).Preload("UserRole").Find(&list)
	c.JSON(http.StatusOK, list)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var obj model.User

	// Record Not Found
	companyId := jwt.GetClaims(c).CompanyId
	result := config.DB.Where("company_id = ? ", companyId).Preload("UserRole").First(&obj, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, obj)
}

func CreateUser(c *gin.Context) {
	var obj model.User
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := util.HashPassword(obj.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	obj.Password = password

	// get details from token
	companyId := jwt.GetClaims(c).CompanyId
	obj.CompanyId = companyId

	if result := config.DB.Create(&obj); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, obj)
}

func UpdateUser(c *gin.Context) {
	var body model.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get details from token
	companyId := jwt.GetClaims(c).CompanyId
	body.CompanyId = companyId

	id := body.ID
	var obj model.User
	// Record Not Found
	result := config.DB.Where("company_id = ? ", companyId).First(&obj, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	if result := config.DB.Save(&body); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success Update", "data": body})
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var obj model.User

	// get details from token
	companyId := jwt.GetClaims(c).CompanyId

	// Record Not Found
	result := config.DB.Where("company_id = ? ", companyId).First(&obj, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	if err := config.DB.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success delete", "data": obj})
	}
}
