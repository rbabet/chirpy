package main

import (
	"fmt"
	"net/http"
)

func main() {
	_ = fmt.Sprintf          // keep fmt
	_ = http.DefaultServeMux // keep net/http

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))

	_ = srv.ListenAndServe()
}
