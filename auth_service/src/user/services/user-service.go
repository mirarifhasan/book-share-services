package services

import (
	"auth_service/db"
	shared "auth_service/src/shared"
	"auth_service/src/user/dtos"
	"auth_service/src/user/models"
	// "encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context, dto dtos.UserSignUpRequest) (res interface{}, err error) {
	var existUsers []models.User
	db.DB.Model(&models.User{}).Where(map[string]interface{}{"phone": dto.Phone}).Find(&existUsers)

	if len(existUsers) != 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "phone already exist"})
		return nil, errors.New("phone already exist")
	}

	var newUser *models.User = new(models.User)
	newUser.Name = dto.Name
	newUser.Phone = dto.Phone
	newUser.Password = dto.Password
	newUser.LastActiveTime = time.Now()

	if err := db.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}

	var responseData dtos.UserSignUpResponse
	responseData.Token = "abc"

	res, _ = shared.StructToMap(responseData)

	return res, nil
}

func Login(context *gin.Context) {
	fmt.Println("Login Service")
}
