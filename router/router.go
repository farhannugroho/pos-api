package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"pos_api/jwt"
	"pos_api/router/endpoint"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	// Templates HTML
	r.LoadHTMLGlob("templates/*")

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Main Router
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	guest := r.Group("/v1")
	{
		// Auth
		guest.POST("/login", endpoint.Login)
		guest.POST("/register", endpoint.Register)

		// Business Type
		businessType := guest.Group("/business_types")
		{
			businessType.GET("", endpoint.GetAllBusinessTypes)
			businessType.GET("/:id", endpoint.GetBusinessTypeById)
			businessType.POST("", endpoint.CreateBusinessType)
			businessType.PUT("", endpoint.UpdateBusinessType)
			businessType.DELETE("/:id", endpoint.DeleteBusinessType)
		}

		// City
		city := guest.Group("/cities")
		{
			city.GET("", endpoint.GetAllCities)
			city.GET("/:id", endpoint.GetCityById)
			city.POST("", endpoint.CreateCity)
			city.PUT("", endpoint.UpdateCity)
			city.DELETE("/:id", endpoint.DeleteCity)
		}

		// Location
		location := guest.Group("/locations")
		{
			location.GET("", endpoint.GetAllLocations)
			location.GET("/:id", endpoint.GetLocationById)
			location.POST("", endpoint.CreateLocation)
			location.PUT("", endpoint.UpdateLocation)
			location.DELETE("/:id", endpoint.DeleteLocation)
		}
	}

	// API Version
	v1 := r.Group("/v1")
	v1.Use(jwt.Middleware)
	{
		// Company
		company := v1.Group("/companies")
		{
			company.GET("", endpoint.GetAllCompanies)
			company.GET("/:id", endpoint.GetCompanyById)
			company.POST("", endpoint.CreateCompany)
			company.PUT("", endpoint.UpdateCompany)
			company.DELETE("/:id", endpoint.DeleteCompany)
		}

		// User
		user := v1.Group("/users")
		{
			user.GET("", endpoint.GetAllUsers)
			user.GET("/:id", endpoint.GetUserById)
			user.POST("", endpoint.CreateUser)
			user.PUT("", endpoint.UpdateUser)
			user.DELETE("/:id", endpoint.DeleteUser)
		}

		// Outlet
		outlet := v1.Group("/outlets")
		{
			outlet.GET("", endpoint.GetAllOutlets)
			outlet.GET("/:id", endpoint.GetOutletById)
			outlet.POST("", endpoint.CreateOutlet)
			outlet.PUT("", endpoint.UpdateOutlet)
			outlet.DELETE("/:id", endpoint.DeleteOutlet)
		}

		// Unit of Measurement
		uom := v1.Group("/uom")
		{
			uom.GET("", endpoint.GetAllUom)
			uom.GET("/:id", endpoint.GetUomById)
			uom.POST("", endpoint.CreateUom)
			uom.PUT("", endpoint.UpdateUom)
			uom.DELETE("/:id", endpoint.DeleteUom)
		}

		// Item Groups
		itemGroup := v1.Group("/item_groups")
		{
			itemGroup.GET("", endpoint.GetAllItemGroups)
			itemGroup.GET("/:id", endpoint.GetItemGroupById)
			itemGroup.POST("", endpoint.CreateItemGroup)
			itemGroup.PUT("", endpoint.UpdateItemGroup)
			itemGroup.DELETE("/:id", endpoint.DeleteItemGroup)
		}

		// Item Categories
		itemCategory := v1.Group("/item_categories")
		{
			itemCategory.GET("", endpoint.GetAllItemCategories)
			itemCategory.GET("/:id", endpoint.GetItemCategoryById)
			itemCategory.POST("", endpoint.CreateItemCategory)
			itemCategory.PUT("", endpoint.UpdateItemCategory)
			itemCategory.DELETE("/:id", endpoint.DeleteItemCategory)
		}

		// Items
		item := v1.Group("/items")
		{
			item.GET("", endpoint.GetAllItems)
			item.GET("/:id", endpoint.GetItemById)
			item.POST("", endpoint.CreateItem)
			item.PUT("", endpoint.UpdateItem)
			item.DELETE("/:id", endpoint.DeleteItem)
		}

		// Item Variants
		itemVariant := v1.Group("/item_variants")
		{
			itemVariant.GET("", endpoint.GetAllItemVariants)
			itemVariant.GET("/:id", endpoint.GetItemVariantById)
			itemVariant.POST("", endpoint.CreateItemVariant)
			itemVariant.PUT("", endpoint.UpdateItemVariant)
			itemVariant.DELETE("/:id", endpoint.DeleteItemVariant)
		}

		// Inventories
		inventory := v1.Group("/inventories")
		{
			inventory.GET("", endpoint.GetAllInventories)
			inventory.GET("/:id", endpoint.GetInventoryById)
			inventory.POST("", endpoint.CreateInventory)
			inventory.PUT("", endpoint.UpdateInventory)
			inventory.DELETE("/:id", endpoint.DeleteInventory)
		}

		// User Role
		userRole := v1.Group("/user_roles")
		{
			userRole.GET("", endpoint.GetAllUserRoles)
			userRole.GET("/:id", endpoint.GetUserRoleById)
			userRole.POST("", endpoint.CreateUserRole)
			userRole.PUT("", endpoint.UpdateUserRole)
			userRole.DELETE("/:id", endpoint.DeleteUserRole)
		}

		// Modules
		module := v1.Group("/modules")
		{
			module.GET("", endpoint.GetAllModules)
			module.GET("/:id", endpoint.GetModuleById)
			module.POST("", endpoint.CreateModule)
			module.PUT("", endpoint.UpdateModule)
			module.DELETE("/:id", endpoint.DeleteModule)
		}

		// Sub Modules
		subModule := v1.Group("/sub_modules")
		{
			subModule.GET("", endpoint.GetAllSubModules)
			subModule.GET("/:id", endpoint.GetSubModuleById)
			subModule.POST("", endpoint.CreateSubModule)
			subModule.PUT("", endpoint.UpdateSubModule)
			subModule.DELETE("/:id", endpoint.DeleteSubModule)
		}
	}

	return r
}
