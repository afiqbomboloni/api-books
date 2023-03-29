package builder

import (
	"api-books/app"
	"api-books/config"

	"api-books/modules/author/v1/repository"
	"api-books/modules/author/v1/service"

	"github.com/gin-gonic/gin"
)
func AuthorBuilder(router *gin.Engine) {
	repo := repository.NewAuthorRepository(config.ConnectDb())
    svc := service.NewAuthorService(repo)
	
	app.AuthorHTTPHandler(router, svc)
}