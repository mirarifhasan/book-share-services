package services

import (
	"auth_service/db"
	"auth_service/src/user/dtos"
	"auth_service/src/user/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(c *gin.Context, dto dtos.UserSignupRequest) (res interface{}, err error) {
	var existUsers []models.User
	db.DB.Model(&models.User{}).Where(map[string]interface{}{"phone": dto.Phone}).Find(&existUsers)

	if len(existUsers) != 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "phone already exist"})
		return nil, errors.New("phone already exist")
	}

	var newUserDto *models.User = new(models.User)
	newUserDto.Name = dto.Name
	newUserDto.Phone = dto.Phone
	newUserDto.Password = dto.Password

	newUser := db.DB.Create(newUserDto)

	return newUser, nil
}

func Login(context *gin.Context) {
	fmt.Println("Login Service")
}
