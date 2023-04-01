package repository

import (
	"api-books/common/interfaces"
	"api-books/entity"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]entity.Book, error)
	FindById(ID int) (*entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
	Delete(book entity.Book) (entity.Book, error)
}

type bookRepository struct {
	db *gorm.DB
	cache interfaces.CacheAble
}

func NewBookRepository(db *gorm.DB, cache interfaces.CacheAble) *bookRepository{
	return &bookRepository{db, cache}
}

func(r *bookRepository) FindAll() ([]entity.Book, error) {
	var books []entity.Book

	cacheKey := "books-all"
	bytes, _ := r.cache.Get(cacheKey)

	if bytes != nil {
		if err := json.Unmarshal(bytes, &books); err != nil {
			return nil, err
		}

		return books, nil
	}

	err := r.db.Preload("Publisher").Find(&books).Error

	if err := r.cache.Set(cacheKey, books, 300); err != nil {
		return nil, err
	}

	return books, err
}

func(r *bookRepository) FindById(ID int) (*entity.Book, error) {
	book := &entity.Book{}

	cacheKey := fmt.Sprintf("book:%d", ID)

	bytes, _ := r.cache.Get(cacheKey)
	
	if bytes != nil {
		if err := json.Unmarshal(bytes, &book); err != nil {
			return nil, err
		}
		return book, nil
	}

	err := r.db.Preload("Publisher").Find(&book, ID).Error

	if err := r.cache.Set(cacheKey, book, 300); err != nil {
		return nil, err
	}

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