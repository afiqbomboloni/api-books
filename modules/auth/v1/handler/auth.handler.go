package handler

import (
	"api-books/modules/auth/v1/service"
	"api-books/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *authHandler {
	return &authHandler{authService}
}

func(h *authHandler) Register(ctx *gin.Context) {
	var authRequest request.AuthRequest

	if err := ctx.ShouldBindJSON(&authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}


	_, err := h.authService.SaveUser(authRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return	
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})


}

func(h *authHandler) Login(ctx *gin.Context) {
	var request request.AuthRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}



	user, err := h.authService.AuthValidate(request.Username, request.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "error",
			"errors": err.Error(),
		})
		ctx.Abort()
		return
	}

	token, err := h.authService.GenerateAccessToken(ctx, user)

	ctx.JSON(http.StatusOK, gin.H{
		"data": user.Username,
		"token":token,
	})


}