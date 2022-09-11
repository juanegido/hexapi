package main

import (
	"fmt"
	"log"
	"net/http"
)

const httpAddr = ":8080"

func main() {
	fmt.Println("Starting server on", httpAddr)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(httpAddr, mux))
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
}
