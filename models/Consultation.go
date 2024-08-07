package models

import (
	"gorm.io/gorm"
)

// Consultation representa una consulta del usuario
type Consultation struct {
	gorm.Model
	FirstName   string `gorm:"not null" json:"first_name"`   // Nombre del usuario
	LastName    string `gorm:"not null" json:"last_name"`    // Apellido del usuario
	Email       string `gorm:"uniqueIndex;not null" json:"email"` // Correo electrónico del usuario
	Phone       string `json:"phone"`                       // Teléfono del usuario
	Consultation string `gorm:"type:text;size:3000" json:"consultation"` // Consulta del usuario, texto de hasta 3000 caracteres
	MoreInfo    bool   `json:"more_info"`                   // Información adicional (booleano)
}
