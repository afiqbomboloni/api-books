package repository

import (
	"api-books/entity"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAll() ([]entity.Author, error)
	FindById(ID int) (entity.Author, error)
	Create(author entity.Author) (entity.Author, error)
	Update(author entity.Author) (entity.Author, error)
	Delete(author entity.Author) (entity.Author, error)
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *authorRepository{
	return &authorRepository{db}
}

func(r *authorRepository) FindAll() ([]entity.Author, error) {
	var authors []entity.Author

	err := r.db.Model(&entity.Author{}).Preload("Books").Find(&authors).Error

	return authors, err
}

func(r *authorRepository) FindById(ID int) (entity.Author, error) {
	var author entity.Author

	err := r.db.Model(&entity.Author{}).Preload("Books").Find(&author, ID).Error

	return author, err
}

func(r *authorRepository) Create(author entity.Author) (entity.Author, error) {
	err := r.db.Debug().Create(&author).Error
	return author, err
}

func(r *authorRepository) Update(author entity.Author) (entity.Author, error){
	err := r.db.Debug().Updates(&author).Error

	return author, err
}

func(r *authorRepository) Delete(author entity.Author) (entity.Author, error) {
	err := r.db.Debug().Delete(&author).Error

	return author, err
}