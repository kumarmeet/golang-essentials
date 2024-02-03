package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimit is a Gin-style middleware function to apply rate limiting
func RateLimit() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(40), 10) // Adjust the rate limit as needed 40 request per 10 seconds

	return func(c *gin.Context) {
		if !limiter.Allow() {
			time.Sleep(1 * time.Second) //slow down
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}
