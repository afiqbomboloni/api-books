package repository

import (
	"api-books/entity"

	"gorm.io/gorm"
)

type PublisherRepository interface {
	FindAll() ([]entity.Publisher, error)
	FindById(ID int) (entity.Publisher, error)
	Create(publisher entity.Publisher) (entity.Publisher, error)
	Update(publisher entity.Publisher) (entity.Publisher, error)
	Delete(publisher entity.Publisher) (entity.Publisher, error)
}

type publisherRepository struct {
	db *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) *publisherRepository{
	return &publisherRepository{db}
}

func(r *publisherRepository) FindAll() ([]entity.Publisher, error) {
	var publishers []entity.Publisher

	err := r.db.Debug().Find(&publishers).Error

	return publishers, err
}

func(r *publisherRepository) FindById(ID int) (entity.Publisher, error) {
	var publisher entity.Publisher

	err := r.db.Find(&publisher, ID).Error

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