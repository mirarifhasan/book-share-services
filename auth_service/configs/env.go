/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvVariables struct {
	Port          string
	DBConstraints string
	AuthServiceBaseUrl string
}

var env *EnvVariables

func initEnvVariable() {
	if env != nil {
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	env = &EnvVariables{
		Port:          os.Getenv("PORT"),
		DBConstraints: os.Getenv("DB_CONSTRAINTS"),
		AuthServiceBaseUrl: os.Getenv("AUTH_SERVICE_BASE_URL"),
	}
}

func GetEnv() *EnvVariables {
	initEnvVariable()
	if env == nil {
		log.Fatal("Error occurred during initEnvVariable method.")
	}
	return env
}
