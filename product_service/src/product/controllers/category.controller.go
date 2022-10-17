package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"product_service/src/product/dtos"
	"product_service/src/product/services"
	sharedDtos "product_service/src/shared/dtos"
)

func CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {

		var dto dtos.CrateCategoryRequest

		if err := c.ShouldBind(&dto); err != nil {
			c.JSON(http.StatusBadRequest, sharedDtos.BuildErrorResponse("", err.Error(), nil))
		}

		data, err := services.CreateACategory(c, dto)
		if err != nil {
			c.JSON(http.StatusForbidden, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		c.JSON(http.StatusCreated, sharedDtos.BuildResponse("", data))

	}
}

func GetCategories() gin.HandlerFunc {
	return func(c *gin.Context) {

		dto := map[string]interface{}{
			"is_active": c.DefaultQuery("is_active", ""),
		}

		data, err := services.GetCategories(dto)
		if err != nil {
			c.JSON(http.StatusForbidden, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		c.JSON(http.StatusCreated, sharedDtos.BuildResponse("", data))

	}
}

func GetACategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hi")
	}
}
