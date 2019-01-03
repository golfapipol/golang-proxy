package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const port = ":4000"

func NewMultipleHostReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		fmt.Println("target", target)
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}

func main() {
	proxy := NewMultipleHostReverseProxy([]*url.URL{
		{
			Scheme: "http",
			Host:   "localhost:3300",
			Path:   "/helloworld",
		},
		{
			Scheme: "http",
			Host:   "localhost:3330",
			Path:   "/helloworld",
		},
	})
	http.ListenAndServe(port, proxy)
}
