/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Avatar         string    `json:"avater"`
	Rating         uint      `json:"rating"`
	Address        string    `json:"address"`
	IpAddress      string    `json:"ip_address"`
	DeviceId       string    `json:"device_id"`
	LastActiveTime time.Time `json:"last_active_time"`
	
	// CreatedAt      time.Time `json:"created_at"`
	// UpdatedAt      time.Time `json:"updated_at"`
}
