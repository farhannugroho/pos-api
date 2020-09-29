package endpoint

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pos_api/config"
	"pos_api/model"
)

func GetAllLocations(c *gin.Context) {
	var list []model.Location
	config.DB.Find(&list)
	c.JSON(http.StatusOK, list)
}

func GetLocationById(c *gin.Context) {
	id := c.Param("id")
	var obj model.Location

	// Record Not Found
	result := config.DB.First(&obj, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, obj)
}

func CreateLocation(c *gin.Context) {
	var obj model.Location
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := config.DB.Create(&obj); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusCreated, obj)
}

func UpdateLocation(c *gin.Context) {
	var body model.Location
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := body.ID
	var obj model.City

	// Record Not Found
	result := config.DB.First(&obj, id)
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

func DeleteLocation(c *gin.Context) {
	id := c.Param("id")
	var obj model.Location

	// Record Not Found
	result := config.DB.First(&obj, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	if err := config.DB.Where("id = ?", id).Delete(&model.Location{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success delete", "data": obj})
	}
}
