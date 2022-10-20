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
		IpAddress:      dto.IpAddress,
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

func Login(context *gin.Context, dto dtos.UserLoginRequest) (interface{}, error) {
	var user models.User
	if err := db.DB.Model(&models.User{}).Where(map[string]interface{}{"phone": dto.Phone, "password": dto.Password}).First(&user).Error; err != nil {
		return nil, err
	}

	if len(dto.IpAddress) > 0 {
		db.DB.Model(&models.User{}).Where(map[string]interface{}{"id": user.ID}).Update("IpAddress", dto.IpAddress)
	}

	responseObj := dtos.UserSignUpResponse{
		Token: GenerateToken(int(user.ID)),
		UserInfo: dtos.UserInfoResponse{
			ID:     int(user.ID),
			Name:   user.Name,
			Avatar: nil,
			Rating: user.Rating,
		},
	}

	res, _ := shared.StructToMap(responseObj)
	return res, nil
}

func GetUsersByIds(ids []int) (interface{}, error){
	fmt.Println("Hi")
	return nil, nil
}
