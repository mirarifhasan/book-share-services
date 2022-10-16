package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hi")
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
