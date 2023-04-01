package builder

import (
	"api-books/app"
	"api-books/config"
	"api-books/utils"

	// "api-books/utils"

	"api-books/modules/author/v1/repository"
	"api-books/modules/author/v1/service"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	// "github.com/gomodule/redigo/redis"
)
func AuthorBuilder(router *gin.Engine, client *redis.Pool) {
	cache := utils.NewClient(client)

	repo := repository.NewAuthorRepository(config.ConnectDb(), cache)
    svc := service.NewAuthorService(repo)
	
	app.AuthorHTTPHandler(router, svc)
}