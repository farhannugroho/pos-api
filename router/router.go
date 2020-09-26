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
		v1.GET("/cities", endpoint.GetCities)
	}

	return r
}
