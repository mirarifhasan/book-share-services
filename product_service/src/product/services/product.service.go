package services

import (
	"errors"
	"fmt"
	"product_service/db"
	"product_service/src/product/dtos"
	"product_service/src/product/models"
	shared "product_service/src/shared"

	"github.com/gin-gonic/gin"
)

func CreateAProduct(c *gin.Context, dto dtos.CreateProductRequest) (interface{}, error) {

	var category models.Category

	err := db.DB.Model(&models.Category{}).Where(map[string]interface{}{"id": dto.CategoryID}).First(&category).Error
	if err != nil {
		return nil, errors.New("category not exist")
	}

	newProd := models.Product{
		Name:        dto.Name,
		Thumbnail:   dto.Thumbnail,
		Description: dto.Description,
		Price:       dto.Price,
		Condition:   dto.Condition,
		SellingBy:   0,
		ApprovedBy:  0,
		CategoryID:  dto.CategoryID,
	}

	if err := db.DB.Create(&newProd).Error; err != nil {
		return nil, err
	}

	return shared.StructToMap(newProd), nil
}

func GetProducts(c *gin.Context, query dtos.GetProductsQuery) (interface{}, error) {

	fmt.Println(query)

	var products []models.Product
	queryBuilder := db.DB.Model(&models.Product{})

	if len(query.Name) > 0 {
		queryBuilder = queryBuilder.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if len(query.CategoryID) > 0 {
		queryBuilder.Where(map[string]interface{}{"category_id":query.CategoryID})
	}

	queryBuilder.Find(&products)

	return products, nil
}
