package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	port = "8080"
)

func main() {
	mux := http.NewServeMux()
	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println("couldn't start server: %w", err)
		os.Exit(-1)
	}
	fmt.Printf("Starting server @ %s", port)
}
