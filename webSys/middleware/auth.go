package middleware

import (
	"PasswordProxy/databaseSys"
	"PasswordProxy/webSys/settings"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Auth(admin bool, database databaseSys.DataBaseStruct) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("jwt")
		if token == nil {
			c.Redirect(http.StatusFound, settings.AuthPath+settings.LoginPagePath)
		} else {
			token := token.(string)

			username, err := JwtVerify(token, database)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "jwt error"})
			}

			user, err := database.FindUser(username)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.Redirect(http.StatusFound, settings.AuthPath+settings.LoginPagePath)
			}

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user error"})
			}

			if admin {
				if user.Admin {
					c.Next()
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "permission denied"})
				}
			} else {
				c.Next()
			}
		}
	}
}
