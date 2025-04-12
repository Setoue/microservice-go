package controllers

import (
	"net/http"

	"github.com/Setoue/microservice-go-category/internal/services"
	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(context *gin.Context) {

	var body createCategoryInput

	if err := context.ShouldBindJSON(&body); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request",
			"message": err.Error(),
		})
		return
	}

	service := services.NewCreateCategoryServices()

	err := service.Execute(body.Name)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request",
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
	})

}
