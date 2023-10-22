package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
)

func GetErrorCodeList() []int {
	return []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusInternalServerError,
		http.StatusBadGateway,
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if slices.Contains(GetErrorCodeList(), c.Writer.Status()) {
			if c.Request.Method == "GET" {
				status := c.Writer.Status()
				filename := fmt.Sprintf("HTTP%d.html", status)
				c.HTML(status, filename, nil)
				c.Abort()
			}
		}
	}
}
