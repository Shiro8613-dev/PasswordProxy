package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		req := c.Request
		dTime := time.Now().Local().Format(time.RFC1123)

		//ip [username] - - [date time] Method URL Proto statusCode
		l := fmt.Sprintf("%s [%s] - - [%s] %s %s %s %d", c.ClientIP(), "name", dTime, req.Method, req.URL.Path, req.Proto, c.Writer.Status())
		fmt.Println(l)
	}
}
