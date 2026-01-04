package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/malachi190/watchcircle/config"
	"github.com/malachi190/watchcircle/handlers"
	"github.com/malachi190/watchcircle/middleware"
)

func Routes(handler *handlers.Handler, cfg *config.Config) func(*gin.Engine) {
	return func(r *gin.Engine) {
		// Add Request Logger
		r.Use(middleware.NewRequestLogger())

		// Health Check
		r.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Welcome to WatchCircle",
			})
		})

		// Versioning
		api := r.Group("/api/v1")

		// Auth
		{
			api.POST("/register", handler.Auth.Register)
			api.POST("/verify-email", handler.Auth.VerifyEmail)
			api.POST("/resend-otp", handler.Auth.ResendOtp)
			api.POST("/login", handler.Auth.Login)
			api.GET("/refetch-token", handler.Auth.FetchRefreshToken)
		}

		// Authenticated Routes
		v1 := api.Use(middleware.NewAuthMiddlware(cfg))
		{
			v1.GET("/test", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"message": "Test completed",
				})
			})
		}
	}
}
