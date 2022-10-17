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
			return
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
		var query dtos.GetCategoriesQuery

		if err := c.Bind(&query); err != nil {
			c.JSON(http.StatusBadRequest, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		fmt.Println("query", query)

		data, err := services.GetCategories(query)
		if err != nil {
			c.JSON(http.StatusForbidden, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		fmt.Println(data)

		c.JSON(http.StatusOK, sharedDtos.BuildResponse("", data))

	}
}
