package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(context *gin.Context, err error) {
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"error":   err.Error(),
		"message": err.Error(),
	})
}

func Ok(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
	})
}