package model

import "time"

// Member represents a registered user stored in PostgreSQL
type Member struct {
	// 10-digit phone-like identity
	VirtualNumber string `gorm:"column:virtual_number;primaryKey" json:"virtual_number"`
	// primary login field
	Email string `gorm:"column:email;uniqueIndex;not null" json:"email"`
	// true after OTP check
	Verified bool `gorm:"column:verified;default:false" json:"verified"`

	// Profile fields
	FirstName  string    `gorm:"column:first_name" json:"first_name"`
	FamilyName string    `gorm:"column:family_name" json:"family_name"`
	Dob        time.Time `gorm:"column:dob" json:"dob"`
	Gender     string    `gorm:"column:gender" json:"gender"`

	// UTC timestamps – set dynamically in Go
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}
