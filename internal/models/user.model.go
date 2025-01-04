package models

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
