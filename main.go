package main

import (
	"log"
	"net/http"
)

func main() {

	const port = "8080"
	const filePathRoot = "."

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(filePathRoot)))

	log.Printf("Serving files from %s on port: %s\n", filePathRoot, port)

	log.Fatal(srv.ListenAndServe())
}
