package middleware

import (
	"PasswordProxy/webSys/settings"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func Proxy(pass string, rewrite bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(pass)
		if err != nil {
			c.AbortWithStatus(500)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			if rewrite {
				req.URL.Path = strings.Replace(c.Request.URL.Path, settings.ProxiedPath, "", 1)
			} else {
				req.URL.Path = c.Request.URL.Path
			}
		}

		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}
