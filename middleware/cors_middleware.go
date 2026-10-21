package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// allowedOrigins lists the browser origins permitted to call the API.
var allowedOrigins = map[string]bool{
	"http://localhost:5173":  true,
	"http://localhost:3000":  true,
	"https://dash-absensi.regaftw.my.id": true,
}

// CORS allows browser frontends on the allowed origins to call the API.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if allowedOrigins[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
