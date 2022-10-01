package dtos

// swagger:parameters UserSignupRequest
type UserSignupRequest struct {
	// required: true
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
	// required: true
	Phone string `form:"phone" json:"phone" xml:"phone" binding:"required"`
	// required: true
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}
