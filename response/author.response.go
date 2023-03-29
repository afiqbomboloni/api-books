package response

import "api-books/entity"

type AuthorResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Books []struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
		Theme string `json:"theme"`
	} `json:"books"`
}

func NewAuthorResponse(author entity.Author) AuthorResponse {
	res := AuthorResponse{
		ID : author.ID,
		Name: author.Name,
		Email: author.Email,
		Books : []struct {
			ID    uint   `json:"id"`
			Title string `json:"title"`
			Theme string `json:"theme"`


		}{},
		
	}

	for _, book := range author.Books {
		res.Books = append(res.Books, struct {
			ID    uint   `json:"id"`
			Title string `json:"title"`
			Theme string `json:"theme"`
		} {
			ID : book.ID,
			Title: book.Title,
			Theme: book.Theme,
		})
	}

	return res
}