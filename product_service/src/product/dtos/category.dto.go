package dtos

// swagger:parameters CrateCategoryRequest
type CrateCategoryRequest struct {
	// required: true
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
	// required: true
	Thumbnail string `form:"thumbnail" json:"thumbnail" xml:"thumbnail" binding:"required"`
	// required: true
	IsActive bool `form:"is_active" json:"is_active" xml:"is_active"  binding:"required"`
}