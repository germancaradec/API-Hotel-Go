package models

import (
	"time"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	Checkin     time.Time `gorm:"not null" json:"check_in"`
	Checkout    time.Time `gorm:"not null" json:"check_out"`
	Email       string    `gorm:"uniqueIndex;not null" json:"email"` 
	UserID      uint      `json:"user_id"`
}
