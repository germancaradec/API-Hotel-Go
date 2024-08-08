package models

import (
	"gorm.io/gorm"
)

type Consultation struct {
	gorm.Model
	Phone         string `json:"phone"`                       
	Consultation  string `gorm:"type:text;size:3000" json:"consultation"` 
	MoreInfo      bool   `json:"more_info"`                   
	UserID        uint   `json:"user_id"`
}
