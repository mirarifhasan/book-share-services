/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package models

import (
	"time"
)

type User struct {
	ID             uint
	Name           string
	Phone          string
	Email          string
	Password       string
	Avatar         string
	Rating         uint
	Address        string
	IpAddress      string
	DeviceId       string
	LastActiveTime time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
