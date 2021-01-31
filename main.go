package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func main() {
	server := http.Server{
		Addr: "120.0.0.1:8080",
	}
	server.ListenAndServe()
}
