package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/malachi190/watchcircle/config"
	"github.com/malachi190/watchcircle/helper"
	"github.com/malachi190/watchcircle/logger"
)

type AuthMiddleware struct {
	Config *config.Config
}

func NewAuthMiddlware(cfg *config.Config) gin.HandlerFunc {
	a := &AuthMiddleware{Config: cfg}
	return a.Handler()
}

func (a *AuthMiddleware) Handler() gin.HandlerFunc {
	secret := []byte(a.Config.Jwt.AccessSecret)
	return func(ctx *gin.Context) {
		auth := strings.TrimSpace(ctx.GetHeader("Authorization"))

		const prefix = "Bearer "

		if !strings.HasPrefix(auth, prefix) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "missing or invalid bearer token",
			})
			return
		}

		tokenStr := auth[len(prefix):]

		claims, err := helper.DecodeJwt(tokenStr, secret)

		if err != nil {
			logger.Error.Printf("error while decoding token: %v", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "invalid or expired token",
			})
			return
		}

		userId, err := claims.GetSubject()

		if err != nil || userId == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "missing subject",
			})
			return
		}

		ctx.Set("user_id", userId)

		ctx.Next()
	}

}
