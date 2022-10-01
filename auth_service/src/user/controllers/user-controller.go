package controllers

import (
	"auth_service/src/user/dtos"
	"auth_service/src/user/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserSigngup
// @Summary  Signup a user
// @Schemes
// @Description  Signup a user for the first time
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request               body    dtos.UserSignupRequest  true  "User Signup Request"
// @Success      201
// @Router       /signup [post]
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Signup Controller")

		var dto dtos.UserSignupRequest

		if err := c.ShouldBind(&dto); err != nil {
			fmt.Println("Error khaise")
		}

		services.Signup(dto)
		c.JSON(http.StatusCreated, nil)
	}
}
