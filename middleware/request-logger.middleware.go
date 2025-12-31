package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/malachi190/watchcircle/logger"
)

type RequestLoggerMiddleware struct {
	console *log.Logger
}

// NewRequestLogger returns a Gin handler that prints to CLI only
func NewRequestLogger() gin.HandlerFunc {
	m := &RequestLoggerMiddleware{console: logger.Console}
	return m.Handler()
}

func (m *RequestLoggerMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		ctx.Next()

		latency := time.Since(start)
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		status := ctx.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		m.console.Printf("| status: %3d | latency: %13v | ip: %15s | path: %s %s",
			status,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
