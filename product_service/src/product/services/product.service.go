package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"product_service/db"
	"product_service/src/product/dtos"
	"product_service/src/product/models"
	shared "product_service/src/shared"
	auth_service "product_service/src/shared/http_clients/auth_service/services"

	auth_service_dtos "product_service/src/shared/http_clients/auth_service/dtos"

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
		Edition:     dto.Edition,
		AuthorName:  dto.AuthorName,
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
		queryBuilder.Where(map[string]interface{}{"category_id": query.CategoryID})
	}

	queryBuilder.Preload("Category").Find(&products)

	userIds := []int{}
	for _, v := range products {
		userIds = append(userIds, v.SellingBy)
	}

	var users []auth_service_dtos.User
	if len(userIds) > 0 {
		respByte, _ := auth_service.GetUsersByIds(userIds)
		var respJson map[string]interface{}
		json.Unmarshal(respByte, &respJson)

		usersByte, _ := json.Marshal(respJson["data"])
		json.Unmarshal(usersByte, &users)
	}

	var productRes []models.ProductResponseDto
	for _, p := range products {

		pByte, _ := json.Marshal(p)
		var newP models.ProductResponseDto
		json.Unmarshal(pByte, &newP)

		for _, u := range users {
			if u.ID == uint(p.SellingBy) {
				newP.Seller = shared.StructToMap(u)
			}
		}

		productRes = append(productRes, newP)
	}

	return productRes, nil

}

func GetAProduct(id int) (interface{}, error) {

	var product models.Product
	if err := db.DB.Model(&models.Product{}).Where(map[string]interface{}{"id": id}).Preload("Category").First(&product).Error; err != nil {
		return nil, err
	}

	resp, _ := auth_service.GetUsersByIds([]int{product.SellingBy})
	var result map[string]interface{}
	_ = json.Unmarshal(resp, &result)

	var productResp map[string]interface{}
	data, _ := json.Marshal(product)
	json.Unmarshal(data, &productResp)

	productResp["seller_info"] = result["data"].([]interface{})[0]

	return productResp, nil

}
