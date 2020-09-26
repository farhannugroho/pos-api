package router

import (
	"github.com/gin-gonic/gin"
	"pos_api/router/endpoint"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// api version
	v1 := r.Group("/v1")
	{
		// city
		v1.GET("/cities", endpoint.GetCities)
	}

	return r
}
