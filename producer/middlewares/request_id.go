package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Request ID middleware
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate UUID
		id := uuid.New().String()
		// Set context variable
		c.Set("x-request-id", id)
		// Set header
		c.Header("x-request-id", id)
		c.Next()
	}
}
