package request


type AuthorRequest struct {
	Name       string      `json:"name" binding:"required"`
	Email       string   `json:"email" binding:"required,email"`

}