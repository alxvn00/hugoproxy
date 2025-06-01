package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(_ http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var target string

		switch {
		case strings.HasPrefix(r.URL.Path, "/api/address"):
			target = "http://geo-service:8081"
		default:
			target = fmt.Sprintf("http://%s:%s", rp.host, rp.port)
		}

		uri, _ := url.Parse(target)

		proxy := httputil.ReverseProxy{Director: func(req *http.Request) {
			req.URL.Scheme = uri.Scheme
			req.URL.Host = uri.Host
			req.Host = uri.Host
		}}

		proxy.ServeHTTP(w, r)
	})
}
