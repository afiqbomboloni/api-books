package builder

import (
	"api-books/app"
	"api-books/config"

	"api-books/modules/auth/v1/repository"
	"api-books/modules/auth/v1/service"

	"github.com/gin-gonic/gin"
)
func AuthorBuilder(router *gin.Engine) {
	repo := repository.NewAuthRepository(config.ConnectDb())
    svc := service.NewAuthRepository(repo)
	
	app.AuthHTTPHandler(router, svc)
}