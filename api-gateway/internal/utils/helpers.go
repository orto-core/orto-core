package utils

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func CreateReverseProxy(route string) gin.HandlerFunc {
	return func(c *gin.Context) {
		target, err := url.Parse(route)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "failed",
			})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(target)

		proxy.Director = func(r *http.Request) {
			r.Header = c.Request.Header
			r.Host = target.Host
			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			r.URL.Path = c.Param("proxyPath")
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
