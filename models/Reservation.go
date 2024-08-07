package models

import (
	"time"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	Adults         int       `json:"adults"`             
	Checkin        time.Time `gorm:"not null" json:"check_in"`
	Checkout       time.Time `gorm:"not null" json:"check_out"`
	Children       int       `json:"children"`           
	Email          string    `gorm:"uniqueIndex;not null" json:"email"`
	NumberOfRooms  int       `json:"number_of_rooms"`    
	RoomType       string    `json:"room_type"`          
	UserID         uint      `json:"user_id"`
}
