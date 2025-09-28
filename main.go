package main

import (
	"fmt"
	"net/http"
)

func main() {
	_ = fmt.Sprintf          // keep fmt
	_ = http.DefaultServeMux // keep net/http
	// TODO: set up mux, server, ListenAndServe

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	_ = srv.ListenAndServe()
}
