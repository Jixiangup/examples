package middleware

import (
	"github.com/gin-gonic/gin"
	error2 "jixiangup.me/examples/gin/internal/error"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) > 0 {
			err := context.Errors.Last().Err
			status := error2.GetStatus(err)
			context.JSON(status, gin.H{"error": err.Error(), "status": status})
		}
	}
}
