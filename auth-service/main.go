/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package main

import (
	"auth_service/db"
	"auth_service/models"
	"auth_service/routers"
)

func main() {
	routers.Routers()
	db.ConnectToDB()
	db.DB.AutoMigrate(models.User{})
}
