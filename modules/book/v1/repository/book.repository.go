package repository

import (
	"api-books/entity"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]entity.Book, error)
	FindById(ID int) (entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
	Delete(book entity.Book) (entity.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository{
	return &bookRepository{db}
}

func(r *bookRepository) FindAll() ([]entity.Book, error) {
	var books []entity.Book

	err := r.db.Preload("Publisher").Find(&books).Error

	return books, err
}

func(r *bookRepository) FindById(ID int) (entity.Book, error) {
	var book entity.Book

	err := r.db.Preload("Publisher").Find(&book, ID).Error

	return book, err
}

func(r *bookRepository) Create(book entity.Book) (entity.Book, error) {
	err := r.db.Debug().Create(&book).Error
	return book, err
}

func(r *bookRepository) Update(book entity.Book) (entity.Book, error){
	err := r.db.Debug().Updates(&book).Error

	return book, err
}

func(r *bookRepository) Delete(book entity.Book) (entity.Book, error) {
	err := r.db.Debug().Delete(&book).Error

	return book, err
}