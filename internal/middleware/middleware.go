package middleware

import (
	"errors"
	"github.com/alinadsm04/backend-for-highload/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var SecretKey = []byte("MY_VERY_VERY_VERY_SECRET_KEY_100%_secure_anyone_can_not_find_out_:)")

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, gin.Error{
					Err:  errors.New("invalid signing method"),
					Type: gin.ErrorTypePrivate,
				}
			}
			return SecretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

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
