package controllers

import (
	"net/http"
	"product_service/src/product/dtos"
	"product_service/src/product/services"
	sharedDtos "product_service/src/shared/dtos"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProduct
// @Summary  create new product
// @Schemes
// @Description  create new product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        request               body     dtos.CreateProductRequest  true  "Create Product Request"
// @Success      201
// @Router       /products [post]
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
		var query dtos.GetProductsQuery

		if err := c.Bind(&query); err != nil {
			c.JSON(http.StatusBadRequest, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		data, err := services.GetProducts(c, query)
		if err != nil {
			c.JSON(http.StatusForbidden, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		c.JSON(http.StatusOK, sharedDtos.BuildResponse("", data))

	}
}

func GetAProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, sharedDtos.BuildErrorResponse("", "id not found or invalid", nil))
			return
		}

		data, err := services.GetAProduct(id)
		if err != nil {
			c.JSON(http.StatusForbidden, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		c.JSON(http.StatusOK, sharedDtos.BuildResponse("", data))
	}
}
