package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

		c.Next()

		log.Printf("Response Status: %d", c.Writer.Status())
	}
}
