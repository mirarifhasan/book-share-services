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
	Name           string    `json:"name"`
	Phone          string    `json:"phone" gorm:"type:varchar(100);uniqueIndex;not null"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Avatar         string    `json:"avatar"`
	Rating         uint      `json:"rating"`
	Address        string    `json:"address"`
	IpAddress      string    `json:"ip_address"`
	DeviceId       string    `json:"device_id"`
	LastActiveTime time.Time `json:"last_active_time,omitempty"`
}
