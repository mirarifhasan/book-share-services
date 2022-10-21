package dtos

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint           `json:"id" `
	Name           string         `json:"name"`
	Phone          string         `json:"phone" `
	Email          string         `json:"email"`
	Password       string         `json:"password"`
	Avatar         string         `json:"avatar"`
	Rating         uint           `json:"rating"`
	Address        string         `json:"address"`
	IpAddress      string         `json:"ip_address"`
	DeviceId       string         `json:"device_id"`
	LastActiveTime time.Time      `json:"last_active_time,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}
