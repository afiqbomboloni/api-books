package repository

import (
	"api-books/entity"


	"gorm.io/gorm"
)


type AuthRepository interface{
	SaveUser(user entity.User) (entity.User, error)
	GetUsername(username string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository{
	return &authRepository{db}
}


func(r *authRepository) SaveUser(user entity.User) (entity.User, error) {
	err := r.db.Debug().Create(&user).Error

	return user, err
}

func(r *authRepository) GetUsername(username string) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("username = ?", username).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil

}




