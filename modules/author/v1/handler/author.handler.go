package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"api-books/modules/author/v1/service"
	"api-books/request"
	"api-books/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type authorHandler struct {
	authorService service.AuthorService
}

func NewAuthorHandler(authorService service.AuthorService) *authorHandler {
	return &authorHandler{authorService}
}

func(h *authorHandler) GetAuthors(ctx *gin.Context) {

	author, err := h.authorService.FindAll()

	if err != nil {
        ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }

	var authorsResponse []response.AuthorResponse

	for _, a := range author {
		authorsResponse = append(authorsResponse, response.NewAuthorResponse(a))
	}
	

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": authorsResponse,
	})

	
}

func(h *authorHandler) GetAuthor(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.authorService.FindById(id)

	if b.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Content Empty",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	authorResponse := response.NewAuthorResponse(b)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": authorResponse,
	})
}

func(h *authorHandler) CreateAuthor(ctx *gin.Context) {
	var authorRequest request.AuthorRequest

	err := ctx.ShouldBindJSON(&authorRequest)

	if err != nil {

		errorMessages := []string{}
		for _,e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, is %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error": errorMessages,
		})
		return
	}

	newAuthor, err := h.authorService.Create(authorRequest)

	authorResponse := response.NewAuthorResponse(newAuthor)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": authorResponse,
	})
}


func(h *authorHandler) UpdateAuthor(ctx *gin.Context) {

	var authorRequest request.AuthorRequest

	err := ctx.ShouldBindJSON(&authorRequest)

	if err != nil {

		errorMessages := []string{}
		for _,e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, is %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error": errorMessages,
		})
		return
	}


	

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	updatedAuthor, err := h.authorService.Update(id, authorRequest)

	authorResponse := response.NewAuthorResponse(updatedAuthor)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": authorResponse,
	})
}

func(h *authorHandler) DeleteAuthor(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.authorService.Delete(id)
	if b.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":"Id Not found",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}


	authorResponse := response.NewAuthorResponse(b)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": authorResponse,
	})
}


// func AuthorResponse(author entity.Author) (response.Author) {
// 	return response.Author{
// 		ID: author.ID,
// 		Name: author.Name,
// 		Email: author.Email,
// 		Books: []response.Book{
			
// 		},
		
// 	}
// }