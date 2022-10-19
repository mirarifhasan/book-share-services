package dtos

// swagger:parameters CrateProductRequest
type CreateProductRequest struct {
	// required: true
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
	// required: false
	Thumbnail string `form:"thumbnail" json:"thumbnail" xml:"thumbnail"`
	// required: true
	Description string `form:"description" json:"description" xml:"description"  binding:"required"`
	// required: true
	Price float32 `form:"price" json:"price" xml:"price"  binding:"required"`
	// required: true
	Condition string `form:"condition" json:"condition" xml:"condition"  binding:"required"`
	// required: true
	SellingBy int `form:"selling_by" json:"selling_by" xml:"selling_by"  binding:"required"`
	// required: true
	CategoryID int `form:"category_id" json:"category_id" xml:"category_id"  binding:"required"`
}