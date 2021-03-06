package endpoint

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pos_api/config"
	"pos_api/model"
)

func GetAllItems(c *gin.Context) {
	var list []model.Item
	config.DB.Preload("ItemVariant").Find(&list)
	c.JSON(http.StatusOK, list)
}

func GetItemById(c *gin.Context) {
	id := c.Param("id")
	var obj model.Item

	// Record Not Found
	result := config.DB.Preload("ItemVariant").First(&obj, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, obj)
}

func CreateItem(c *gin.Context) {
	obj := &model.Item{}
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//
	//obj.ItemId = int(obj.ID)
	//if result := config.DB.Create(&obj); result.Error != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	//	return
	//}
	c.JSON(http.StatusCreated, obj)
}

func UpdateItem(c *gin.Context) {
	var body model.Item
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := body.ID
	var obj model.Item

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

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var obj model.Item

	// Record Not Found
	result := config.DB.First(&obj, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	if err := config.DB.Where("id = ?", id).Delete(&model.Item{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success delete", "data": obj})
	}
}
