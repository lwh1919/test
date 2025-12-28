package main

import (
	"fmt"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func headers(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		for _, v2 := range v {
			fmt.Fprintln(w, k, v2)
		}
	}
}

func main() {
	http.HandleFunc("/", pong)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":9999", nil)
}
