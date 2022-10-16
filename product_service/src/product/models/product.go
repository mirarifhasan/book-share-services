/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */

package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	condition   string `json:"condition"`
	ApprovedBy  int    `json:"approved_by"`
	SellingBy   int    `json:"selling_by"`

	//CategoryID int      `json:"category_id"`
	Category Category //` gorm:"foreignKey:ID"`
}
