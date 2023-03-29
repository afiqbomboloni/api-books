package handler

import (
	"api-books/modules/publisher/v1/service"
	"api-books/request"
	"api-books/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type publisherHandler struct {
	publisherService service.PublisherService
}

func NewPublisherHandler(publisherService service.PublisherService) *publisherHandler {
	return &publisherHandler{publisherService}
}


func(h *publisherHandler) GetPublishers(ctx *gin.Context) {

	publisher, err := h.publisherService.FindAll()

	if err != nil {
        ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
        return
    }

	var publishersResponse []response.PublisherResponse

	for _, p := range publisher {
		publisherResponse := response.NewPublisherResponse(p)

		publishersResponse = append(publishersResponse, publisherResponse)
	}
	

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": publisher,
	})

	
}

func(h *publisherHandler) GetPublisher(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	p, err := h.publisherService.FindById(id)

	if p.ID == 0 {
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

	publisherResponse := response.NewPublisherResponse(p)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": publisherResponse,
	})
}

func(h *publisherHandler) CreatePublisher(ctx *gin.Context) {
	var publisherRequest request.PublisherRequest

	err := ctx.ShouldBindJSON(&publisherRequest)

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

	newPublisher, err := h.publisherService.Create(publisherRequest)


	publisherResponse := response.NewPublisherResponse(newPublisher)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": publisherResponse,
	})
}

func(h *publisherHandler) UpdatePublisher(ctx *gin.Context) {

	var publisherRequest request.PublisherRequest

	err := ctx.ShouldBindJSON(&publisherRequest)

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

	updatedPublisher, err := h.publisherService.Update(id, publisherRequest)

	publisherResponse := response.NewPublisherResponse(updatedPublisher)
	

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": publisherResponse,
	})
}

func(h *publisherHandler) DeletePublisher(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	p, err := h.publisherService.Delete(id)
	if p.ID == 0 {
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


	publisherResponse := response.NewPublisherResponse(p)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": publisherResponse,
	})
}

