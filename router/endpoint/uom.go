package endpoint

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pos_api/config"
	"pos_api/jwt"
	"pos_api/model"
)

func GetAllUom(c *gin.Context) {
	var list []model.UnitOfMeasurement
	companyId := jwt.GetClaims(c).CompanyId
	config.DB.Where("company_id = ? ", companyId).Find(&list)
	c.JSON(http.StatusOK, list)
}

func GetUomById(c *gin.Context) {
	id := c.Param("id")
	var obj model.UnitOfMeasurement

	// Record Not Found
	companyId := jwt.GetClaims(c).CompanyId
	result := config.DB.Where("company_id = ? ", companyId).First(&obj, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, obj)
}

func CreateUom(c *gin.Context) {
	var obj model.UnitOfMeasurement
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get details from token
	companyId := jwt.GetClaims(c).CompanyId
	obj.CompanyId = companyId

	if result := config.DB.Create(&obj); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, obj)
}

func UpdateUom(c *gin.Context) {
	var body model.UnitOfMeasurement
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get details from token
	companyId := jwt.GetClaims(c).CompanyId
	body.CompanyId = companyId

	id := body.ID
	var obj model.UnitOfMeasurement
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

func DeleteUom(c *gin.Context) {
	id := c.Param("id")
	var obj model.UnitOfMeasurement

	// get details from token
	companyId := jwt.GetClaims(c).CompanyId

	// Record Not Found
	result := config.DB.Where("company_id = ? ", companyId).First(&obj, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	if err := config.DB.Where("id = ? and company_id = ?", id, companyId).Delete(&model.UnitOfMeasurement{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success delete", "data": obj})
	}
}
