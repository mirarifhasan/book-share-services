package dtos

// swagger:parameters UserSignUpRequest
type UserSignUpRequest struct {
	// required: true
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
	// required: true
	Phone string `form:"phone" json:"phone" xml:"phone" binding:"required"`
	// required: true
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
	// required: false
	IpAddress string `form:"ip_address" json:"ip_address" xml:"ip_address"`
}

// swagger:parameters UserLoginRequest
type UserLoginRequest struct {
	// required: true
	Phone string `form:"phone" json:"phone" xml:"phone" binding:"required"`
	// required: true
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	// required: false
	IpAddress string `form:"ip_address" json:"ip_address" xml:"ip_address"`
}

type UserInfoResponse struct {
	ID     int    `form:"id" json:"id" xml:"id"`
	Name   string `form:"name" json:"name" xml:"name"`
	Avatar *string `form:"avatar" json:"avatar" xml:"avatar"`
	Rating uint   `form:"rating" json:"rating" xml:"rating"`
}

// swagger:parameters UserSignUpResponse
type UserSignUpResponse struct {
	Token    string           `form:"token" json:"token" xml:"token"`
	UserInfo UserInfoResponse `form:"user_info" json:"user_info" xml:"user_info"`
}
