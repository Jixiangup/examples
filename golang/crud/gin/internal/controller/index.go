package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(context *gin.Context, exception error) {
	if exception == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "status": 500})
	}
	_ = context.Error(exception)
	context.Abort()
}
