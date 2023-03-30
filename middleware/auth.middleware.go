package middleware

import (
	"api-books/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := utils.TokenValid(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"errors":  err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
