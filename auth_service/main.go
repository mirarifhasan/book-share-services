/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package main

import (
	"auth_service/db"
	"auth_service/src/user/models"
	"auth_service/routes"
)

func main() {
	routes.Register()
	db.ConnectToDB()
	db.DB.AutoMigrate(models.User{})
}