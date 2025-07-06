package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/tuanngoo192003/gateway-demo-go/gateway/configs"
)

func AuthAPIProxy() http.Handler {
	config := configs.GetConfig()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url, err := url.Parse(config.Endpoints.Auth)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		httpProxy := httputil.NewSingleHostReverseProxy(url)
		httpProxy.ServeHTTP(w, r)
	})
}
