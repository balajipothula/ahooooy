package model

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Email        string    `gorm:"unique;not null" json:"email"`
    PasswordHash string    `gorm:"not null" json:"-"`
    Provider     string    `gorm:"default:'local'" json:"provider"` // local, google, facebook, twitter
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
