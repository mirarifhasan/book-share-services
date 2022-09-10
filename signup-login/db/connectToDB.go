/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"help/configs"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := configs.GetEnv().DBConstraints
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}
	DB = db
}
