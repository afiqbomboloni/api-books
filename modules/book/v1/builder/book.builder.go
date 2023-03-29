package builder

import (
	"api-books/app"
	"api-books/config"

	"api-books/modules/book/v1/repository"
	"api-books/modules/book/v1/service"

	"github.com/gin-gonic/gin"
)
func BookBuilder(router *gin.Engine) {
	repo := repository.NewBookRepository(config.ConnectDb())
    svc := service.NewBookService(repo)
	
	app.BookHTTPHandler(router, svc)
}