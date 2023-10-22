package middleware

import (
	"PasswordProxy/webSys/settings"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func Proxy(pass string, rewrite bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(pass)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
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

		proxy.ModifyResponse = func(res *http.Response) error {
			res.Header.Del("Content-Length")
			res.Header.Del("Server")
			res.Header.Del("Date")

			if slices.Contains(GetErrorCodeList(), res.StatusCode) {
				res.Body = &http.NoBody
			}

			return nil
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
