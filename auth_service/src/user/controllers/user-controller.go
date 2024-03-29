package controllers

import (
	sharedDtos "auth_service/src/shared/dtos"
	"auth_service/src/user/dtos"
	"auth_service/src/user/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// SignUp
// @Summary  SignUp a user
// @Schemes
// @Description  SignUp a user for the first time
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request               body     dtos.UserSignUpRequest  true  "User SignUp Request"
// @Success      201 				   {object} sharedDtos.Response{data=dtos.UserSignUpResponse}
// @Router       /signup [post]
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("SignUp Controller")

		var dto dtos.UserSignUpRequest

		if err := c.ShouldBind(&dto); err != nil {
			c.JSON(http.StatusBadRequest, sharedDtos.BuildErrorResponse("", err.Error(), nil))
		}

		data, err := services.SignUp(c, dto)

		if err != nil {
			c.JSON(http.StatusForbidden, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		c.JSON(http.StatusCreated, sharedDtos.BuildResponse("SignUp success", data))
	}
}

// Login
// @Summary  Login a user
// @Schemes
// @Description  Login a user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request               body     dtos.UserLoginRequest  true  "User Login Request"
// @Success      201 				   {object} sharedDtos.Response{data=dtos.UserSignUpResponse}
// @Router       /login [post]
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto dtos.UserLoginRequest

		if err := c.ShouldBind(&dto); err != nil {
			c.JSON(http.StatusBadRequest, sharedDtos.BuildErrorResponse("", err.Error(), nil))
		}

		data, err := services.Login(c, dto)

		if err != nil {
			c.JSON(http.StatusNotFound, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		c.JSON(http.StatusCreated, sharedDtos.BuildResponse("Login success", data))
	}
}

func GetUsersByIds() gin.HandlerFunc {
	return func(c *gin.Context) {

		var dto dtos.GetUsersByIdsRequest
		if err := c.ShouldBind(&dto); err != nil {
			c.JSON(http.StatusBadRequest, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		idsInStr := strings.Split(dto.Ids, ",")

		var ids []int
		for i := 0; i < len(idsInStr); i++ {
			val := strings.Trim(idsInStr[i], " ")
			id, err := strconv.Atoi(val)
			if err != nil {
				c.JSON(http.StatusBadRequest, sharedDtos.BuildErrorResponse("", err.Error(), nil))
				return
			}
			ids = append(ids, id)
		}
		fmt.Println("ids", ids)

		data, err := services.GetUsersByIds(ids)
		if err != nil {
			c.JSON(http.StatusNotFound, sharedDtos.BuildErrorResponse("", err.Error(), nil))
			return
		}

		c.JSON(http.StatusCreated, sharedDtos.BuildResponse("", data))
	}
}
