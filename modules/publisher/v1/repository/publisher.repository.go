package repository

import (
	"api-books/common/interfaces"
	"api-books/entity"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type PublisherRepository interface {
	FindAll() ([]entity.Publisher, error)
	FindById(ID int) (*entity.Publisher, error)
	Create(publisher entity.Publisher) (entity.Publisher, error)
	Update(publisher entity.Publisher) (entity.Publisher, error)
	Delete(publisher entity.Publisher) (entity.Publisher, error)
}

type publisherRepository struct {
	db *gorm.DB
	cache interfaces.CacheAble
}

func NewPublisherRepository(db *gorm.DB, cache interfaces.CacheAble) *publisherRepository{
	return &publisherRepository{db, cache}
}

func(r *publisherRepository) FindAll() ([]entity.Publisher, error) {
	var publishers []entity.Publisher
	cacheKey := "publishers-all"

	bytes,_ := r.cache.Get(cacheKey)
	if bytes != nil {
		if err := json.Unmarshal(bytes, &publishers); err != nil {
			return nil, err
		}

		return publishers, nil
	}

	err := r.db.Debug().Find(&publishers).Error

	if err := r.cache.Set(cacheKey, publishers, 300); err != nil {
		return nil, err
	}

	return publishers, err
}

func(r *publisherRepository) FindById(ID int) (*entity.Publisher, error) {
	publisher := &entity.Publisher{}

	cacheKey := fmt.Sprintf("publisher:%d", ID)

	bytes, _ := r.cache.Get(cacheKey)

	if bytes != nil {
		if err := json.Unmarshal(bytes, &publisher); err != nil {
			return nil, err
		}

		return publisher, nil
	}

	

	err := r.db.Find(&publisher, ID).Error
	if err := r.cache.Set(cacheKey, publisher, 300); err != nil {
		return nil, err
	}

	return publisher, err
}

func(r *publisherRepository) Create(publisher entity.Publisher) (entity.Publisher, error) {
	err := r.db.Debug().Create(&publisher).Error
	return publisher, err
}

func(r *publisherRepository) Update(publisher entity.Publisher) (entity.Publisher, error){
	err := r.db.Debug().Updates(&publisher).Error

	return publisher, err
}

func(r *publisherRepository) Delete(publisher entity.Publisher) (entity.Publisher, error) {
	err := r.db.Debug().Delete(&publisher).Error

	return publisher, err
}