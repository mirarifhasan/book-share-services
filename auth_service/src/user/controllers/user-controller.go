package controllers

import (
	"auth_service/src/user/dtos"
	"auth_service/src/user/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserSignUp
// @Summary  SignUp a user
// @Schemes
// @Description  SignUp a user for the first time
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request               body     dtos.UserSignUpRequest  true  "User SignUp Request"
// @Success      201 				   {object} dtos.UserSignUpResponse
// @Router       /signup [post]
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("SignUp Controller")

		var dto dtos.UserSignUpRequest

		if err := c.ShouldBind(&dto); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		res, err := services.SignUp(c, dto)

		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}

		c.JSON(http.StatusCreated, res)
	}
}
