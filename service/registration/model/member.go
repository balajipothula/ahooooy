package model

import "time"

// Member represents a registered user in ahooooy.com
// Member represents a registered user stored in PostgreSQL
type Member struct {
	// 10-digit phone-like identity
	VirtualNumber string `gorm:"column:virtual_number;primaryKey" json:"virtual_number"`
	// primary login field
	Email string `gorm:"column:email;uniqueIndex;not null" json:"email"`
	// true after OTP check
	Verified bool `gorm:"column:verified;default:false" json:"verified"`
	// UTC timestamp
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	// UTC timestamp
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
