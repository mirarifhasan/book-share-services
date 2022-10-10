/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package main

import (
	"auth_service/db"
	"auth_service/routes"
	"auth_service/src/user/models"
)

// @title        Auth Service API Doc
// @version      1.0
// @description  API Documentation
// @BasePath     /api/v1

//@in header
//@name Authorization
//@persistAuthorization true
func main() {
	db.ConnectToDB()
	db.DB.AutoMigrate(models.User{})

	routes.Register()
}
