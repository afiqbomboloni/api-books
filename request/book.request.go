package request

type BookRequest struct {
	Title     string	`json:"title" binding:"required"`
	Theme     string	`json:"theme" binding:"required"`
	AuthorId	uint	`json:"author_id" binding:"required"`
}