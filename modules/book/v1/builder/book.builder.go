package builder

import (
	"api-books/app"
	"api-books/config"
	"api-books/utils"

	"api-books/modules/book/v1/repository"
	"api-books/modules/book/v1/service"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)
func BookBuilder(router *gin.Engine, client *redis.Pool) {
	
	cache := utils.NewClient(client)
	repo := repository.NewBookRepository(config.ConnectDb(), cache)
    svc := service.NewBookService(repo)
	
	app.BookHTTPHandler(router, svc)
}