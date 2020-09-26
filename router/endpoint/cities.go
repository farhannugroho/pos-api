package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pos_api/config"
	"pos_api/model"
)

func GetCities(c *gin.Context) {
	var list []model.City
	config.DB.Find(&list)
	c.JSON(http.StatusOK, list)
}
