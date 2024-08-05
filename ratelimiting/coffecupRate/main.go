package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	requests     = make(map[string]int)       // To track requests per IP
	requestTimes = make(map[string]time.Time) // To track the last request time
	mu           sync.Mutex
	limit        = 5              // Max requests per day
	window       = 24 * time.Hour // Time window for rate limiting
)

func rateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		if lastRequestTime, exists := requestTimes[ip]; exists {
			// Check if the time window has expired
			if now.Sub(lastRequestTime) > window {
				// Reset the count and the last request time
				requests[ip] = 0
				requestTimes[ip] = now
			}
		} else {
			// Initialize for new IP
			requestTimes[ip] = now
		}

		if requests[ip] >= limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Rate limit exceeded. You can request up to 5 cups per day.",
			})
			c.Abort()
			return
		}

		// Allow the request and increment the count
		requests[ip]++
		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Group("/cc")
	{
		router.Use(rateLimiter())

		router.GET("/coffee", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Here's your coffee cup!",
			})
		})

		router.NoMethod(func(ctx *gin.Context) {
			ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		})

		router.Run(":8080")
	}
}

// Apply rate limiter middleware
