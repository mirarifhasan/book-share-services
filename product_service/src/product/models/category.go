/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name      string    `json:"name"`
	Thumbnail string    `json:"thumbnail"`
	IsActive  bool      `json:"is_active"`
	ParentId  *Category `json:"parent_id" gorm:"foreignKey:ID"`

	//Product []Product `gorm:"foreignKey:ID"`
}
