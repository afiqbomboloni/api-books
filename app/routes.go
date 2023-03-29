package app

import (
	"api-books/middleware"
	authHandlerv1 "api-books/modules/auth/v1/handler"
	authServicev1 "api-books/modules/auth/v1/service"
	authorHandlerv1 "api-books/modules/author/v1/handler"
	authorServicev1 "api-books/modules/author/v1/service"
	bookHandlerv1 "api-books/modules/book/v1/handler"
	bookServicev1 "api-books/modules/book/v1/service"
	publisherHandlerv1 "api-books/modules/publisher/v1/handler"
	publisherServicev1 "api-books/modules/publisher/v1/service"

	"github.com/gin-gonic/gin"
)

func AuthorHTTPHandler(router *gin.Engine, as authorServicev1.AuthorService) {
	h := authorHandlerv1.NewAuthorHandler(as)
	v1 := router.Group("v1")

	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("/authors", h.GetAuthors)
		v1.GET("/authors/:id", h.GetAuthor)
		v1.POST("/authors", h.CreateAuthor)
		v1.PUT("/authors/:id", h.UpdateAuthor)
		v1.DELETE("/authors/:id", h.DeleteAuthor)
	}
	
}


func BookHTTPHandler(router *gin.Engine, bs bookServicev1.BookService) {
	h := bookHandlerv1.NewBookHandler(bs)
	v1 := router.Group("v1")

	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("/books", h.GetBooks)
		v1.GET("/books/:id", h.GetBook)
		v1.POST("/books", h.CreateBook)
		v1.PUT("/books/:id", h.UpdateBook)
		v1.DELETE("/books/:id", h.DeleteBook)
	}
	
}

func PublisherHTTPHandler(router *gin.Engine, ps publisherServicev1.PublisherService) {
	h := publisherHandlerv1.NewPublisherHandler(ps)
	v1 := router.Group("v1")

	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("/publishers", h.GetPublishers)
		v1.GET("/publishers/:id", h.GetPublisher)
		v1.POST("/publishers", h.CreatePublisher)
		v1.PUT("/publishers/:id", h.UpdatePublisher)
		v1.DELETE("/publishers/:id", h.DeletePublisher)
	}
	
}

func AuthHTTPHandler(router *gin.Engine, as authServicev1.AuthService) {
	h :=authHandlerv1.NewAuthHandler(as)
	v1 := router.Group("v1")

	
	v1.POST("/register", h.Register)
	v1.POST("/login", h.Login)
}