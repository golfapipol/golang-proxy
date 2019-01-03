package main

import (
	"net/http"
)

const port = ":4000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		println("--->", port, req.URL.String())
	})
	http.ListenAndServe(port, nil)
}
