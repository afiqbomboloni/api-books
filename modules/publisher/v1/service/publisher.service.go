package service

import (
	"api-books/entity"
	"api-books/modules/publisher/v1/repository"
	"api-books/request"
)

type PublisherService interface {
	FindAll() ([]entity.Publisher, error)
	FindById(ID int) (entity.Publisher, error)
	Create(publisherRequest request.PublisherRequest) (entity.Publisher, error)
	Update(ID int, publisherRequest request.PublisherRequest) (entity.Publisher, error)
	Delete(ID int) (entity.Publisher, error)
}

type publisherService struct {
	publisherRepository repository.PublisherRepository
}

func NewPublisherService(publisherRepository repository.PublisherRepository) *publisherService {
	return &publisherService{publisherRepository}
}


func(s *publisherService) FindAll() ([]entity.Publisher, error) {
	publishers, err := s.publisherRepository.FindAll()
	return publishers, err

}

func(s *publisherService) FindById(ID int) (entity.Publisher, error) {
	publisher, err := s.publisherRepository.FindById(ID)
	return publisher, err
}

func(s *publisherService) Create(publisherRequest request.PublisherRequest) (entity.Publisher, error) {
	publisher := entity.Publisher {
		Name: publisherRequest.Name,
		City: publisherRequest.City,
		BookId: publisherRequest.BookId,
		
	}

	newPublisher, err := s.publisherRepository.Create(publisher)

	return newPublisher, err
}

func(s *publisherService) Update(ID int, publisherRequest request.PublisherRequest) (entity.Publisher, error) {
	publisher, err := s.publisherRepository.FindById(ID)

	publisher.Name = publisherRequest.Name
	publisher.City = publisherRequest.City
	publisher.BookId = publisherRequest.BookId

	updatedPublisher, err := s.publisherRepository.Update(publisher)

	return updatedPublisher, err
}

func(s *publisherService) Delete(ID int) (entity.Publisher, error) {
	publisher, err := s.publisherRepository.FindById(ID)

	deletedPublisher, err := s.publisherRepository.Delete(publisher)

	return deletedPublisher, err
}