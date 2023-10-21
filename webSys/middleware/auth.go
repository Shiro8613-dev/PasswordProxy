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
		auth, err := AuthCheck(admin, c, database)
		if auth {
			c.Next()
		} else {
			if err == "not Login" {
				c.Redirect(http.StatusFound, settings.AuthPath+settings.LoginPagePath)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err})
				c.Abort()
			}
		}
	}
}

func AuthCheck(admin bool, c *gin.Context, database databaseSys.DataBaseStruct) (bool, string) {
	salt, err := database.ReadCrypto()
	if err != nil {
		return false, "server error"
	}

	session := sessions.Default(c)
	token := session.Get("jwt")
	if token == nil {
		return false, "not Login"
	} else {
		token := token.(string)

		username, err := JwtVerify(token, salt.Salt)
		if err != nil {
			return false, "jwt error"
		}

		user, err := database.FindUser(username)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, "user not exists"
		}

		if err != nil {
			return false, "login error"
		}

		if admin {
			if user.Admin {
				return true, ""
			} else {
				return false, "permission denied"
			}
		} else {
			return true, ""
		}
	}
}
