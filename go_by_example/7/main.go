package main

import (
	"fmt"
	"net/http"
	"time"
)

func pong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("pong-st")
	defer fmt.Println("pong-ed")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "pong")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(err)
		InternalError := http.StatusInternalServerError
		http.Error(w, http.StatusText(InternalError), InternalError)
	}

}

func main() {
	http.HandleFunc("/", pong)
	http.ListenAndServe(":9999", nil)
}
