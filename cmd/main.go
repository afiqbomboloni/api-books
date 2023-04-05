package main

import (
	"api-books/config"
	"api-books/entity"
	authBuilder "api-books/modules/auth/v1/builder"
	authorBuilder "api-books/modules/author/v1/builder"
	bookBuilder "api-books/modules/book/v1/builder"
	publisherBuilder "api-books/modules/publisher/v1/builder"
	"api-books/utils"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func main() {

	config.LoadConfig()
	
	config.ConnectDb().Debug().AutoMigrate(&entity.Author{}, &entity.Book{}, &entity.Publisher{}, &entity.User{})
	fmt.Println("Success")

	r := gin.New()
	redisClient := buildRedisPool()

	BuildHandler(r, redisClient)


	r.Run("localhost:8080")
}

func BuildHandler(router *gin.Engine, client *redis.Pool) {
	authorBuilder.AuthorBuilder(router, client)
	bookBuilder.BookBuilder(router, client)
	publisherBuilder.PublisherBuilder(router, client)
	authBuilder.AuthorBuilder(router)



}


func buildRedisPool() *redis.Pool {
	cfg := config.LoadConfig()
	cachePool := utils.NewPool(cfg.RedisAddr, cfg.DBPass)

	ctx := context.Background()
	_, err := cachePool.GetContext(ctx)

	if err != nil {
		panic(err)
	}

	log.Print("redis successfully connected!")
	return cachePool
}

