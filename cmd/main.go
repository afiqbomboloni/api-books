package main

import (
	"api-books/config"
	"api-books/entity"
	authorBuilder "api-books/modules/author/v1/builder"
	bookBuilder "api-books/modules/book/v1/builder"
	publisherBuilder "api-books/modules/publisher/v1/builder"
	authBuilder "api-books/modules/auth/v1/builder"
	"fmt"


	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()
	
	config.ConnectDb().Debug().AutoMigrate(entity.Author{}, entity.Book{}, entity.Publisher{}, entity.User{})
	fmt.Println("Success")

	r := gin.New()

	BuildHandler(r)


	r.Run("localhost:8080")
}

func BuildHandler(router *gin.Engine) {
	authorBuilder.AuthorBuilder(router)
	bookBuilder.BookBuilder(router)
	publisherBuilder.PublisherBuilder(router)
	authBuilder.AuthorBuilder(router)


}

