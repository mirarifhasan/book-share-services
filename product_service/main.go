package main

import (
	"fmt"
	"product_service/db"
	"product_service/routes"
	"product_service/src/product/models"
)

// @title        Product Service API Doc
// @version      1.0
// @description  API Documentation
// @BasePath     /api/v1

//@in header
//@name Authorization
//@persistAuthorization true
func main() {
	fmt.Println("Product Service starting")

	db.ConnectToDB()
	db.DB.AutoMigrate(models.Product{})

	routes.Register()
}
