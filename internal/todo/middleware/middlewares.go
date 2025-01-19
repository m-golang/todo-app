package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SecureHeaders is a middleware that adds secure and common headers to the HTTP response.
// It sets the 'Connection', 'Content-Type', and 'Date' headers for each request.
func SecureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Connection", "keep-alive")
		c.Header("Content-Type", "application/json")
		c.Header("Date", time.Now().UTC().Format(time.RFC1123))

		c.Next()
	}
}

// RecoverPanic is a middleware that handles panics in the Gin application.
// If a panic occurs during request processing, it will recover and return a 500 Internal Server Error response.
func RecoverPanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Defer function to recover from panic if any occurs during the request lifecycle
		defer func() {
			if err := recover(); err != nil {
				c.Header("Connection", "close")
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": http.StatusText(http.StatusInternalServerError),
				})
			}
		}()
		c.Next()
	}
}
