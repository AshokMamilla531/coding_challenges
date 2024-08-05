package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	requests = make(map[string]time.Time)
	mu       sync.Mutex
	limit    = 5               // Max requests
	window   = 1 * time.Minute // Time window
)

func rateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		lastRequest, exists := requests[ip]
		if exists && time.Since(lastRequest) < window {
			if len(requests) >= limit {
				c.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
				mu.Unlock()
				c.Abort()
				return
			}
		}
		requests[ip] = time.Now()
		mu.Unlock()

		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(rateLimiter())

	router.GET("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Request successful",
		})
	})

	router.Run(":8080")
}
