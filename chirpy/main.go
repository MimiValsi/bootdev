package main

import (
	"fmt"
	"log"
	"net/http"
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

	fmt.Printf("Starting server on %s", port)
	log.Fatal(s.ListenAndServe())

}
