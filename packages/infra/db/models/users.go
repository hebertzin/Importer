package models

import (
	"time"
)

type Users struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
