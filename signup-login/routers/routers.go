/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"help/configs"
)

func Routers() *gin.Engine {
	router := gin.Default()

	router.Run(configs.GetEnv().Port)
	return router
}
