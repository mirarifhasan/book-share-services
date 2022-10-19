package controllers

import (
	"fmt"
	"net/http"
	"product_service/src/product/dtos"
	"product_service/src/product/services"
	sharedDtos "product_service/src/shared/dtos"

	"github.com/gin-gonic/gin"
)

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		var dto dtos.CreateProductRequest

		if err := c.ShouldBind(&dto); err != nil {
			c.JSON(http.StatusBadRequest, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		data, err := services.CreateAProduct(c, dto)
		if err != nil {
			c.JSON(http.StatusForbidden, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		c.JSON(http.StatusCreated, sharedDtos.BuildResponse("", data))

	}
}

func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hi")
	}
}

func GetAProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hi")
	}
}
