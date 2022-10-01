package services

import (
	"auth_service/db"
	"auth_service/src/user/dtos"
	"auth_service/src/user/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Signup(dto dtos.UserSignupRequest) {

	var user models.User
	result := db.DB.Where(map[string]interface{}{"id": 1}).Find(&user)

	fmt.Sprintln(&result)
	fmt.Println("Signup Service")
}

func Login(context *gin.Context) {
	fmt.Println("Login Service")
}
