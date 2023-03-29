package entity

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name      string
	City     	string
	BookId   uint
	// BookId		uint	`json:"book_id"`
	// Book	Book
}