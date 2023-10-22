package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

func ServerHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		dTime := time.Now().Local().Format(time.RFC1123)
		c.Header("Server", "SD-Made-Server")
		c.Header("Date", dTime)
		c.Next()
	}
}
