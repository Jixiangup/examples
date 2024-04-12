package controller

import "github.com/gin-gonic/gin"

func Error(context *gin.Context, exception error) {
	if exception == nil {
		context.JSON(500, gin.H{"message": "Internal Server Error", "status": 500})
	}
	_ = context.Error(exception)
	context.Abort()
}
