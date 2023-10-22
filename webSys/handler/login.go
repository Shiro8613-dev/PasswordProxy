package handler

import (
	"PasswordProxy/databaseSys"
	"PasswordProxy/utils/cryptoSys"
	"PasswordProxy/utils/jwtSys"
	"PasswordProxy/webSys/middleware"
	"PasswordProxy/webSys/settings"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type loginJsonStruct struct {
	Username  string `json:"username"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
	Password3 string `json:"password3"`
	PinCode   int    `json:"pinCode"`
}

func LoginPage(database databaseSys.DataBaseStruct) gin.HandlerFunc {
	return func(c *gin.Context) {
		b, err := middleware.AuthCheck(false, c, database)

		if b {
			c.Redirect(http.StatusFound, settings.ProxiedPath)
		} else {
			if err != "not Login" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
				c.Abort()
			} else {
				c.HTML(http.StatusOK, "login.html", nil)
			}
		}
	}
}

func LoginApi(database databaseSys.DataBaseStruct) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginJson loginJsonStruct
		if err := c.ShouldBindJSON(&loginJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this json struct is bad"})
		}

		user, err := database.FindUser(loginJson.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not Exist"})
		}

		if cryptoSys.VerifyPassword(loginJson.Password1, user.Password1) &&
			cryptoSys.VerifyPassword(loginJson.Password2, user.Password2) &&
			cryptoSys.VerifyPassword(loginJson.Password3, user.Password3) &&
			cryptoSys.VerifyPassword(strconv.Itoa(loginJson.PinCode), user.PinCode) {

			//loginOk
			session := sessions.Default(c)

			token, err := jwtSys.JwtGenerate(loginJson.Username, database)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "jwt error"})
			}

			session.Set("jwt", token)
			session.Options(sessions.Options{
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
				MaxAge:   int(time.Hour.Seconds()) * 24,
			})
			err = session.Save()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "session error"})
			}

			c.JSON(http.StatusOK, gin.H{"message": "login"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		}
	}
}
