package routes

import (
	c "auth_service/src/user/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouteSetup(router *gin.Engine) {

	v1 := router.Group("api/v1")
	v1.POST("/signup", c.SignUp())
	v1.POST("/login", c.Login())

}
