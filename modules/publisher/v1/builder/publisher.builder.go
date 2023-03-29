package builder

import (
	"api-books/app"
	"api-books/config"

	"api-books/modules/publisher/v1/repository"
	"api-books/modules/publisher/v1/service"

	"github.com/gin-gonic/gin"
)
func PublisherBuilder(router *gin.Engine) {
	repo := repository.NewPublisherRepository(config.ConnectDb())
    svc := service.NewPublisherService(repo)
	
	app.PublisherHTTPHandler(router, svc)
}