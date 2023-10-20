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
)

type WebServer struct {
	conf configSys.ListenerConfig
	e    *gin.Engine
}

// NewWebServer webserver
func NewWebServer(conf configSys.ListenerConfig, proxy configSys.ProxyConfig, database databaseSys.DataBaseStruct, store sessions.Store) WebServer {
	r := gin.Default()

	r.Use(gin.Recovery(), middleware.Logger())
	r.Use(sessions.Sessions("session", store))

	//after change
	r.Static("/_next", "D:\\DEV\\PasswordProxy\\frontend\\out\\_next")
	r.LoadHTMLGlob("D:\\DEV\\PasswordProxy\\frontend\\out\\*.html")
	//------------
	authGroup := r.Group(settings.AuthPath)
	{
		authGroup.GET(settings.LoginPagePath, handler.LoginPage)
		authGroup.POST(settings.LoginPagePath, handler.LoginApi(database))
		authGroup.GET(settings.LogoutPagePath, handler.LogoutPage)
		authGroup.POST(settings.LogoutPagePath, handler.LogoutApi)
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
	return w.e.Run(fmt.Sprintf("%s:%d", w.conf.Host, w.conf.Port))
}