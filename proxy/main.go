package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

const port = ":4000"

func main() {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:3300",
	})
	http.ListenAndServe(port, proxy)
}
