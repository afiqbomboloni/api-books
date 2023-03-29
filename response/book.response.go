package response

import "api-books/entity"

type BookResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Theme     string `json:"theme"`
	AuthorId  uint   `json:"author_id"`
	Publisher struct {
		Name string `json:"name"`
		City string `json:"city"`
	} `json:"publisher"`
}

func NewBookResponse(book entity.Book) BookResponse {
	res := BookResponse {
		ID: book.ID,
		Title: book.Title,
		Theme: book.Theme,
		AuthorId: book.AuthorId,
		Publisher: struct {
			Name string `json:"name"`
			City string `json:"city"`
		} {
			Name: book.Publisher.Name,
			City: book.Publisher.City,
		},
	}

	return res
}
