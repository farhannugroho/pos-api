package router

import (
	"github.com/gin-gonic/gin"
	"pos_api/router/endpoint"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// home route
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"server": "OK",
		})
	})

	// api version
	v1 := r.Group("/v1")
	{
		// city
		v1.GET("/cities", endpoint.GetCities)
	}

	return r
}
