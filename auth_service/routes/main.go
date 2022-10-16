/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package routes

import (
	"auth_service/configs"
	"auth_service/src/user/routes"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/ghodss/yaml"
)

func Register() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"app": "Auth Service running",
		})
	})

	swaggerRoute(router)
	routes.UserRouteSetup(router)

	router.Run(":" + configs.GetEnv().Port)
	return router
}

func swaggerRoute(router *gin.Engine) {
	router.GET("swagger/swagger.json", func(c *gin.Context) {
		f, err := os.Open("./docs/swagger.json")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		dec := json.NewDecoder(f)
		data := make(map[string]interface{}, 0)
		if err := dec.Decode(&data); err != nil {
			panic(err)
		}
		appUrl := configs.GetEnv().AuthServiceBaseUrl
		appUrl = strings.Replace(appUrl, "https", "", -1)
		appUrl = strings.Replace(appUrl, "http", "", -1)
		data["host"] = appUrl

		c.JSON(200, data)
		// Res.
		// 	Raw(data).
		// 	Json(c)
	})

	router.GET("swagger/swagger.yaml", func(c *gin.Context) {
		f, err := ioutil.ReadFile("./docs/swagger.yaml")
		if err != nil {
			panic(err)
		}

		data := make(map[string]interface{}, 0)
		err = yaml.Unmarshal([]byte(f), &data)
		if err != nil {
			panic(err)
		}
		appUrl := configs.GetEnv().AuthServiceBaseUrl
		appUrl = strings.Replace(appUrl, "https://", "", -1)
		appUrl = strings.Replace(appUrl, "http://", "", -1)
		data["host"] = appUrl

		c.JSON(200, data)
		// Res.
		// 	Raw(data).
		// 	Yaml(c)
	})

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(configs.GetEnv().AuthServiceBaseUrl+"/swagger/swagger.yaml"), ginSwagger.DefaultModelsExpandDepth(-1)))
}
