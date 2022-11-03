package dtos

import (
	"time"
)

type ProductResponseDto struct {
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
	// Category    Category               `json:"category"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	// DeletedAt   gorm.DeletedAt         `json:"deleted_at"`
}
