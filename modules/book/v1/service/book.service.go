package service

import (
	"api-books/entity"
	"api-books/modules/book/v1/repository"
	"api-books/request"
)

type BookService interface {
	FindAll() ([]entity.Book, error)
	FindById(ID int) (entity.Book, error)
	Create(bookRequest request.BookRequest) (entity.Book, error)
	Update(ID int, bookRequest request.BookRequest) (entity.Book, error)
	Delete(ID int) (entity.Book, error)
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) *bookService {
	return &bookService{bookRepository}
}


func(s *bookService) FindAll() ([]entity.Book, error) {
	books, err := s.bookRepository.FindAll()
	return books, err

}

func(s *bookService) FindById(ID int) (entity.Book, error) {
	book, err := s.bookRepository.FindById(ID)
	return *book, err
}

func(s *bookService) Create(bookRequest request.BookRequest) (entity.Book, error) {
	book := entity.Book {
		Title: bookRequest.Title,
		Theme: bookRequest.Theme,
		AuthorId: bookRequest.AuthorId,
		
		
	}

	newBook, err := s.bookRepository.Create(book)

	return newBook, err
}

func(s *bookService) Update(ID int, bookRequest request.BookRequest) (entity.Book, error) {
	book, err := s.bookRepository.FindById(ID)

	book.Title = bookRequest.Title
	book.Theme = bookRequest.Theme
	book.AuthorId = bookRequest.AuthorId

	updatedBook, err := s.bookRepository.Update(*book)

	return updatedBook, err
}

func(s *bookService) Delete(ID int) (entity.Book, error) {
	book, err := s.bookRepository.FindById(ID)

	deletedBook, err := s.bookRepository.Delete(*book)

	return deletedBook, err
}