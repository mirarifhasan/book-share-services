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

type Category struct {
	// gorm.Model
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name" gorm:"size:50;uniqueIndex"`
	Thumbnail string         `json:"thumbnail"`
	OrderIdx  int            `json:"order_idx" gorm:"default:0;"`
	IsActive  bool           `json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Products  []Product      `json:"products"`
}

type Product struct {
	// gorm.Model
	ID          uint           `json:"id" gorm:"primarykey"`
	Name        string         `json:"name"`
	Thumbnail   string         `json:"thumbnail"`
	Edition     string         `json:"edition"`
	AuthorName  string         `json:"author_name"`
	Description string         `json:"description"`
	Price       float32        `json:"price"`
	Condition   string         `json:"condition"`
	ApprovedBy  int            `json:"approved_by"`
	SellingBy   int            `json:"selling_by"`
	CategoryID  int            `json:"-"`
	Category    Category       `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type ProductResponseDto struct {
	// gorm.Model
	ID          uint                   `json:"id"`
	Name        string                 `json:"name"`
	Thumbnail   string                 `json:"thumbnail"`
	Edition     string                 `json:"edition"`
	AuthorName  string                 `json:"author_name"`
	Description string                 `json:"description"`
	Price       float32                `json:"price"`
	Condition   string                 `json:"condition"`
	ApprovedBy  int                    `json:"approved_by"`
	SellingBy   int                    `json:"selling_by"`
	Seller      map[string]interface{} `json:"seller"`
	CategoryID  int                    `json:"-"`
	Category    Category               `json:"category"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	DeletedAt   gorm.DeletedAt         `json:"deleted_at"`
}
