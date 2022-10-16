package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hi")
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
