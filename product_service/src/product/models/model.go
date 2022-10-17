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
	Name      string `json:"name" gorm:"size:50;uniqueIndex"`
	Thumbnail string `json:"thumbnail"`
	OrderIdx  int    `json:"order_idx" gorm:"default:0;"`
	IsActive  bool   `json:"is_active"`
	// Product   []Product `gorm:"foreignKey:ID"`
}

type Product struct {
	gorm.Model
	Name        string   `json:"name"`
	Thumbnail   string   `json:"thumbnail"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Condition   string   `json:"condition"`
	ApprovedBy  int      `json:"approved_by"`
	SellingBy   int      `json:"selling_by"`
	CategoryID  int      `json:"-"`
	Category    Category `gorm:"foreignKey:CategoryID;references:ID"`
}
