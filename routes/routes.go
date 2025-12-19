package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/malachi190/watchcircle/handlers"
)

func Routes(handler *handlers.Handler) func(*gin.Engine) {
	return func(r *gin.Engine) {
		// Add Request Logger and CORS middlwares

		// Health Check
		r.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Welcome to WatchCircle",
			})
		})

		// Versioning
		api := r.Group("/api/v1")

		{
			api.GET("/test")
		}

	}
}
