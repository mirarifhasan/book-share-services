package controllers

import (
	"fmt"
	"net/http"

	sharedDtos "product_service/src/shared/dtos"
	"product_service/src/product/services"
	"product_service/src/product/dtos"
	"github.com/gin-gonic/gin"
)

func CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hi")

		var dto dtos.CrateCategoryRequest

		if err := c.ShouldBind(&dto); err != nil {
			c.JSON(http.StatusBadRequest, sharedDtos.BuildErrorResponse("", err.Error(), nil))
		}

		data, err := services.CreateACategory(c, dto)
		if err != nil {
			c.JSON(http.StatusForbidden, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}
		
		c.JSON(http.StatusCreated, sharedDtos.BuildResponse("SignUp success", data))

	}
}

func GetCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hi")
	}
}

func GetACategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hi")
	}
}
