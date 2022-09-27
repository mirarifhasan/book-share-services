package controllers

import (
	"auth_service/src/user/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Signup Controller")

		services.Signup(c)
		c.JSON(http.StatusCreated, nil)
	}
}
