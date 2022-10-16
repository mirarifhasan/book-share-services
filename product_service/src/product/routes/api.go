package routes

import (
	c "product_service/src/product/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRouteSetup(router *gin.Engine) {

	v1 := router.Group("api/v1")

	v1.POST("/categories", c.CreateCategory())
	v1.GET("/categories", c.GetCategories())
	v1.GET("/categories/:id", c.GetACategory())

	v1.POST("/products", c.CreateProduct())
	v1.GET("/products", c.GetProducts())
	v1.GET("/products/:id", c.GetAProduct())

}
