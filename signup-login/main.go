/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package main

import (
	"help/db"
	"help/models"
	"help/routers"
)

func main() {
	routers.Routers()
	db.ConnectToDB()
	db.DB.AutoMigrate(models.Product{})
}
