package builder

import (
	"api-books/app"
	"api-books/config"
	"api-books/utils"

	"api-books/modules/publisher/v1/repository"
	"api-books/modules/publisher/v1/service"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)
func PublisherBuilder(router *gin.Engine, client *redis.Pool) {
	cache := utils.NewClient(client)
	repo := repository.NewPublisherRepository(config.ConnectDb(), cache)
    svc := service.NewPublisherService(repo)
	
	app.PublisherHTTPHandler(router, svc)
}