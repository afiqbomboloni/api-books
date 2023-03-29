package handler

import (
	"api-books/modules/book/v1/service"
	"api-books/request"
	"api-books/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *bookHandler {
	return &bookHandler{bookService}
}

func(h *bookHandler) GetBooks(ctx *gin.Context) {

	book, err := h.bookService.FindAll()

	if err != nil {
        ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
        return
    }

	var booksResponse []response.BookResponse

	for _, b := range book {
		booksResponse = append(booksResponse, response.NewBookResponse(b))
	}
	

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": booksResponse,
	})

	
}

func(h *bookHandler) GetBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindById(id)

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

	bookResponse := response.NewBookResponse(b)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": bookResponse,
	})
}

func(h *bookHandler) CreateBook(ctx *gin.Context) {
	var bookRequest request.BookRequest

	err := ctx.ShouldBindJSON(&bookRequest)

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

	newBook, err := h.bookService.Create(bookRequest)

	bookResponse := response.NewBookResponse(newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": bookResponse,
	})
}

func(h *bookHandler) UpdateBook(ctx *gin.Context) {

	var bookRequest request.BookRequest

	err := ctx.ShouldBindJSON(&bookRequest)

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

	updatedBook, err := h.bookService.Update(id, bookRequest)

	bookResponse := response.NewBookResponse(updatedBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": bookResponse,
	})
}

func(h *bookHandler) DeleteBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.Delete(id)
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


	bookResponse := response.NewBookResponse(b)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": bookResponse,
	})
}
