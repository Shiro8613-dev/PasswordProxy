package webSys

import (
	"PasswordProxy/configSys"
	"PasswordProxy/databaseSys"
	"PasswordProxy/webSys/handler"
	"PasswordProxy/webSys/middleware"
	"PasswordProxy/webSys/settings"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebServer struct {
	conf configSys.ListenerConfig
	e    *gin.Engine
}

func init() {
	gin.SetMode(gin.ReleaseMode)

}

// NewWebServer webserver
func NewWebServer(conf configSys.ListenerConfig, proxy configSys.ProxyConfig, database databaseSys.DataBaseStruct, store sessions.Store) WebServer {
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger(), middleware.ServerHeader(), middleware.ErrorHandler())
	r.Use(sessions.Sessions("session", store))

	//after change
	r.Static("/_next", "D:\\DEV\\PasswordProxy\\frontend\\out\\_next")
	r.LoadHTMLGlob("D:\\DEV\\PasswordProxy\\frontend\\out\\*.html")
	//------------

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, settings.ProxiedPath)
	})

	authGroup := r.Group(settings.AuthPath)
	{
		authGroup.GET(settings.LoginPagePath, handler.LoginPage(database))
		authGroup.POST(settings.LoginPagePath, handler.LoginApi(database))
		logoutGroup := authGroup.Group(settings.LogoutPagePath)
		logoutGroup.Use(middleware.Auth(false, database))
		{
			logoutGroup.GET("/", handler.LogoutPage)
			logoutGroup.POST("/", handler.LogoutApi)
		}
	}

	proxyGroup := r.Group(settings.ProxiedPath)
	proxyGroup.Use(middleware.Auth(false, database))
	{
		proxyGroup.Any("/*all", middleware.Proxy(proxy.Pass, proxy.Rewrite))
	}

	return WebServer{
		conf: conf,
		e:    r,
	}
}

func (w WebServer) Start() error {
	fmt.Printf("webserver is started on %s:%d\n", w.conf.Host, w.conf.Port)
	return w.e.Run(fmt.Sprintf("%s:%d", w.conf.Host, w.conf.Port))
}
