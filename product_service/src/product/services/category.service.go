package services

import (
	"github.com/gin-gonic/gin"
	"product_service/db"
	"product_service/src/product/dtos"
	"product_service/src/product/models"
	shared "product_service/src/shared"

	"gorm.io/gorm"
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

	return shared.StructToMap(newCat), nil
}

func GetCategories(query dtos.GetCategoriesQuery) (interface{}, error) {

	var filterQuery = map[string]interface{}{}
	if query.IsActive == "true" {
		filterQuery["is_active"] = true
	} else if query.IsActive == "false" {
		filterQuery["is_active"] = false
	}

	var categories []models.Category
	var dbRes *gorm.DB

	if len(filterQuery) > 0 {
		dbRes = db.DB.Model(&models.Category{}).Where(filterQuery).Find(&categories)
	} else {
		dbRes = db.DB.Model(&models.Category{}).Find(&categories)
	}

	if err := dbRes.Error; err != nil {
		return nil, err
	}

	return categories, nil

}
