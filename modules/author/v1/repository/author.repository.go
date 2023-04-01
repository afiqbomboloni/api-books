package repository

import (
	"api-books/common/interfaces"
	"api-books/entity"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAll() ([]entity.Author, error)
	FindById(ID int) (*entity.Author, error)
	Create(author entity.Author) (entity.Author, error)
	Update(author entity.Author) (entity.Author, error)
	Delete(author entity.Author) (entity.Author, error)
}

type authorRepository struct {
	db *gorm.DB
	cache interfaces.CacheAble

}

func NewAuthorRepository(db *gorm.DB, cache interfaces.CacheAble) *authorRepository{
	return &authorRepository{
		db, cache}
}

func(r *authorRepository) FindAll() ([]entity.Author, error) {
	var authors []entity.Author

	cacheKey := "authors-all"
	bytes, _ := r.cache.Get(cacheKey)

	if bytes != nil {
		if err := json.Unmarshal(bytes, &authors); err != nil {
			return nil, err
		}

		return authors, nil
	}

	err := r.db.Model(&entity.Author{}).Preload("Books").Find(&authors).Error
	if err != nil {
		return nil, err
	}

	if err := r.cache.Set(cacheKey, authors, 300); err != nil {
		return nil, err
	} 
	
	return authors, err
}

func(r *authorRepository) FindById(ID int) (*entity.Author, error) {
	
	author := &entity.Author{}
	cacheKey := fmt.Sprintf("author:%d", ID)
	bytes,_ := r.cache.Get(cacheKey)

	if bytes != nil {
		if err := json.Unmarshal(bytes, &author); err != nil {
			return nil, err
		}
		return author, nil
	}


	err := r.db.Model(&entity.Author{}).Preload("Books").Find(&author, ID).Error

	if err := r.cache.Set(cacheKey, author, 300); err != nil {
		return nil, err
	}


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