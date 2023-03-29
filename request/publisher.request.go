package request

type PublisherRequest struct {
	Name      string	`json:"name"`
	City     	string	`json:"city"`
	BookId    uint		`json:"book_id"`
}