package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RateLimiterConfig struct {
	RedisClient *redis.Client
	MaxRequests int
	Window      time.Duration
}

func RateLimiter(cfg RateLimiterConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use IP address as identifier (you can modify this to use user ID or other identifiers)
		identifier := c.ClientIP()
		key := fmt.Sprintf("rate_limit:%s", identifier)

		// Get the current count from Redis
		val, err := cfg.RedisClient.Get(context.Background(), key).Int()
		if err != nil && err != redis.Nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Rate limiter error"})
			c.Abort()
			return
		}

		if val >= cfg.MaxRequests {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		// Increment the counter
		_, err = cfg.RedisClient.Incr(context.Background(), key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Rate limiter error"})
			c.Abort()
			return
		}

		// Set expiry if the key is new
		if val == 0 {
			_, err = cfg.RedisClient.Expire(context.Background(), key, cfg.Window).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Rate limiter error"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
