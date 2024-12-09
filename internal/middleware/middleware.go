package middleware

import (
	"github.com/alinadsm04/backend-for-highload/internal/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Calculate request duration
		duration := time.Since(start).Seconds()

		// Get status code
		status := c.Writer.Status()

		// Update metrics
		utils.RequestCount.WithLabelValues(c.Request.Method, c.FullPath(), string(status)).Inc()
		utils.RequestLatency.WithLabelValues(c.Request.Method, c.FullPath(), string(status)).Observe(duration)
	}
}
