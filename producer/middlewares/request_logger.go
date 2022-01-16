package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Request logging middleware
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		request_id := c.GetString("x-request-id")
		client_ip := c.ClientIP()
		user_agent := c.Request.UserAgent()
		method := c.Request.Method
		path := c.Request.URL.Path

		t := time.Now()

		c.Next()

		latency := float32(time.Since(t).Seconds())

		status := c.Writer.Status()
		log.Info().Str("request_id", request_id).Str("client_ip", client_ip).
			Str("user_agent", user_agent).Str("method", method).Str("path", path).
			Float32("latency", latency).Int("status", status).Msg("")

	}
}
