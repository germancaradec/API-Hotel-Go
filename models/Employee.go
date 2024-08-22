package models

type Employee struct {

	User   *User `gorm:"embedded" json:"user"`

	Position    string  `gorm:"not null" json:"position"`
	Salary      float64 `gorm:"not null" json:"salary"`
	Department  string  `gorm:"not null" json:"department"`
	HireDate    string  `gorm:"not null" json:"hire_date"` 
	PhoneNumber string  `json:"phone_number"`

}