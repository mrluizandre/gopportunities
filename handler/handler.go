package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{
		"message": "POST Opening",
	})
}

func ShowOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET Opening",
	})
}

func DeleteOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "DELETE Opening",
	})
}

func UpdateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PUT Opening",
	})
}

func ListOpeningsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "GET Opening",
	})
}
