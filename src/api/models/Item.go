package models

import "time"

// Profile struct
type Profile struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName   string    `gorm:"size:255;not null;" json:"firstName"`
	LastName    string    `gorm:"size:100;not null;" json:"lastName"`
	Description string    `gorm:"size:100;" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// UpdateProfile struct
type UpdateProfile struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Description string `json:"description"`
}