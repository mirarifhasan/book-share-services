/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	Port string `default:"3000"`

	DbUsername string `default:"root"`
	DbPassword string `default:"password"`
	DbHostPort string `default:"localhost:3306"`
	DbName     string `default:"auth_db_name"`

	AuthServiceBaseUrl    string `default:"http://localhost:3000"`
	ProductServiceBaseUrl string `default:"http://localhost:3000"`
}

var env *EnvVariables

func initEnvVariable() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	env = &EnvVariables{
		Port: os.Getenv("PORT"),

		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHostPort: os.Getenv("DB_HOST_PORT"),
		DbName:     os.Getenv("DB_NAME"),

		AuthServiceBaseUrl:    os.Getenv("AUTH_SERVICE_BASE_URL"),
		ProductServiceBaseUrl: os.Getenv("PRODUCT_SERVICE_BASE_URL"),
	}
}

func GetEnv() *EnvVariables {
	initEnvVariable()
	if env == nil {
		log.Fatal("Error occurred during initEnvVariable method.")
	}
	return env
}
