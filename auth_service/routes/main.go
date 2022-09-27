/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package routes

import (
	"auth_service/configs"
	"auth_service/src/user/routes"
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"app": "Auth Service running",
		})
	})

	routes.UserRouteSetup(router)

	router.Run(configs.GetEnv().Port)
	return router
}
