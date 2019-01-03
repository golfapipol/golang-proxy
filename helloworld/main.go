package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "helloworld")
	})
	log.Fatal(http.ListenAndServe(":3300", nil))
}
