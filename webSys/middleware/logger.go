package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger(salt string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		req := c.Request
		dTime := time.Now().Local().Format(time.RFC1123)
		name := "NotLogin"
		session := sessions.Default(c)
		key := session.Get("jwt")
		if key != nil {
			name, _ = JwtVerify(key.(string), salt)
		}

		//ip [username] - - [date time] Method URL Proto statusCode
		l := fmt.Sprintf("%s [%s] - - [%s] %s %s %s %d", c.ClientIP(), name, dTime, req.Method, req.URL.Path, req.Proto, c.Writer.Status())
		fmt.Println(l)
	}
}
