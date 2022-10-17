package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"product_service/db"
	"product_service/src/product/dtos"
	"product_service/src/product/models"
	shared "product_service/src/shared"
)

func CreateACategory(c *gin.Context, dto dtos.CrateCategoryRequest) (interface{}, error) {

	newCat := models.Category{
		Name:      dto.Name,
		Thumbnail: dto.Thumbnail,
		IsActive:  dto.IsActive,
	}

	if err := db.DB.Create(&newCat).Error; err != nil {
		return nil, err
	}

	res, _ := shared.StructToMap(newCat)
	return res, nil
}

func GetCategories(dto map[string]interface{}) (interface{}, error) {
	fmt.Println("dto", dto)

	return nil, nil
}
