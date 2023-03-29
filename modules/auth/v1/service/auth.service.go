package service

import (
	"api-books/entity"
	"api-books/modules/auth/v1/repository"
	"api-books/request"
	"context"
	"errors"

	"api-books/utils"

	"golang.org/x/crypto/bcrypt"
)


type AuthService interface {
	SaveUser(authRequest request.AuthRequest) (entity.User, error)
	GenerateAccessToken(ctx context.Context, user *entity.User) (string, error)
	AuthValidate(username, password string) (*entity.User, error)
}


type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthRepository(authRepository repository.AuthRepository) *authService {
	return &authService{authRepository}
}

func(s *authService) SaveUser(authRequest request.AuthRequest) (entity.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(authRequest.Password), bcrypt.DefaultCost)

	auth := entity.User {
		Username: authRequest.Username,
		Password:string(hash),
	}
	newAuth, err := s.authRepository.SaveUser(auth)

	return newAuth, err
}

func(s *authService) GenerateAccessToken(ctx context.Context, user *entity.User) (string, error) {

	token, err := utils.GenerateToken(user.ID)


	return token, err
}

func(s *authService) AuthValidate(username, password string) (*entity.User, error) {
	user, err := s.authRepository.GetUsername(username)

	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.New("Invalid Credentials")
	}

	return user, err
}