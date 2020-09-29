package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pos_api/router/endpoint"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Templates HTML
	r.LoadHTMLGlob("templates/*")

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Main Router
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// API Version
	v1 := r.Group("/v1")
	{
		// City
		city := v1.Group("/cities")
		{
			city.GET("", endpoint.GetAllCities)
			city.GET("/:id", endpoint.GetCityById)
			city.POST("", endpoint.CreateCity)
			city.PUT("", endpoint.UpdateCity)
			city.DELETE("/:id", endpoint.DeleteCity)
		}

		// Business Type
		businessType := v1.Group("/business_types")
		{
			businessType.GET("", endpoint.GetAllBusinessTypes)
			businessType.GET("/:id", endpoint.GetBusinessTypeById)
			businessType.POST("", endpoint.CreateBusinessType)
			businessType.PUT("", endpoint.UpdateBusinessType)
			businessType.DELETE("/:id", endpoint.DeleteBusinessType)
		}
	}

	return r
}
