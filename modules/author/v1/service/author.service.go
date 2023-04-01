package service

import (
	"api-books/entity"
	"api-books/modules/author/v1/repository"
	"api-books/request"
)

type AuthorService interface {
	FindAll() ([]entity.Author, error)
	FindById(ID int) (entity.Author, error)
	Create(authorRequest request.AuthorRequest) (entity.Author, error)
	Update(ID int, authorRequest request.AuthorRequest) (entity.Author, error)
	Delete(ID int) (entity.Author, error)
}

type authorService struct {
	authorRepository repository.AuthorRepository
}

func NewAuthorService(authorRepository repository.AuthorRepository) *authorService {
	return &authorService{authorRepository}
}


func(s *authorService) FindAll() ([]entity.Author, error) {
	authors, err := s.authorRepository.FindAll()
	return authors, err

}

func(s *authorService) FindById(ID int) (entity.Author, error) {
	author, err := s.authorRepository.FindById(ID)
	return *author, err
}

func(s *authorService) Create(authorRequest request.AuthorRequest) (entity.Author, error) {
	author := entity.Author {
		Name: authorRequest.Name,
		Email: authorRequest.Email,
	}

	newAuthor, err := s.authorRepository.Create(author)

	return newAuthor, err
}

func(s *authorService) Update(ID int, authorRequest request.AuthorRequest) (entity.Author, error) {
	author, err := s.authorRepository.FindById(ID)

	author.Name = authorRequest.Name
	author.Email = authorRequest.Email

	updatedAuthor, err := s.authorRepository.Update(*author)

	return updatedAuthor, err
}

func(s *authorService) Delete(ID int) (entity.Author, error) {
	author, err := s.authorRepository.FindById(ID)

	deletedAuthor, err := s.authorRepository.Delete(*author)

	return deletedAuthor, err
}