package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	fpRoot = "."
	port   = "8080"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(fpRoot)))

	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Starting server on %s", port)
	log.Fatal(s.ListenAndServe())

}
