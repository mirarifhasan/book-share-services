package services

import (
	"auth_service/db"
	shared "auth_service/src/shared"
	"auth_service/src/user/dtos"
	"auth_service/src/user/models"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context, dto dtos.UserSignUpRequest) (res interface{}, err error) {
	var existUsers []models.User
	db.DB.Model(&models.User{}).Where(map[string]interface{}{"phone": dto.Phone}).Find(&existUsers)

	if len(existUsers) != 0 {
		return nil, errors.New("phone already exist")
	}

	newUser := models.User{
		Name:           dto.Name,
		Phone:          dto.Phone,
		Password:       dto.Password,
		LastActiveTime: time.Now(),
	}

	if err := db.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}

	responseObj := dtos.UserSignUpResponse{
		Token: GenerateToken(int(newUser.ID)),
		UserInfo: dtos.UserInfoResponse{
			ID:     int(newUser.ID),
			Name:   newUser.Name,
			Avatar: nil,
			Rating: newUser.Rating,
		},
	}

	res, _ = shared.StructToMap(responseObj)

	return res, nil
}

func Login(context *gin.Context) {
	fmt.Println("Login Service")
}
