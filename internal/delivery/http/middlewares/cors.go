package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/radenadri/go-boilerplate/config"
	"github.com/radenadri/go-boilerplate/utils"
)

type CORSConfig struct {
	AllowOrigins []string
	AllowMethods []string
	AllowHeaders []string
}

func DefaultCORSConfig() CORSConfig {
	return CORSConfig{
		AllowOrigins: []string{config.CORSAllowOrigins},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}
}

func CORS(cfg CORSConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set allowed origins
		origin := c.Request.Header.Get("Origin")
		for _, allowedOrigin := range cfg.AllowOrigins {
			if allowedOrigin == "*" || allowedOrigin == origin {
				c.Header("Access-Control-Allow-Origin", origin)
				break
			}
		}

		// Set other CORS headers
		c.Header("Access-Control-Allow-Methods", utils.JoinStrings(cfg.AllowMethods))
		c.Header("Access-Control-Allow-Headers", utils.JoinStrings(cfg.AllowHeaders))
		c.Header("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
