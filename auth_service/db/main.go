/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package db

import (
	"auth_service/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := configs.GetEnv().DbUsername + ":" + configs.GetEnv().DbPassword + "@tcp(" + configs.GetEnv().DbHost + ":" + configs.GetEnv().DbPort + ")/"
	fmt.Println("dsn: ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in DB server connection")
	}

	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + configs.GetEnv().DbName + ";")

	dsn = dsn + configs.GetEnv().DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in DB connection:", err)
	}

	DB = db
	fmt.Println("DB Connected")
}
