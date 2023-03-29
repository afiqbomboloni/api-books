package response

import "api-books/entity"

type PublisherResponse struct {
	Name string `json:"name"`
	City string `json:"city"`
	BookId   uint   `json:"book_id"`
}

func NewPublisherResponse(publisher entity.Publisher) PublisherResponse {
	res := PublisherResponse {
		Name: publisher.Name,
		City: publisher.City,
		BookId: publisher.BookId,
	}

	return res
}