package entity

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	ID    uint `gorm:"primary_key"`
	Name  string
	Email string
	Books []Book	`json:"books"`
}
