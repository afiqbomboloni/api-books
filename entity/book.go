package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID        uint		`gorm:"primary_key"`
	Title     string
	Theme     string
	AuthorId	uint	`json:"-"`
	Publisher	Publisher 
}