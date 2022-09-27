package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	fmt.Println("Signup Service")
}

func Login(context *gin.Context) {
	fmt.Println("Login Service")
}
