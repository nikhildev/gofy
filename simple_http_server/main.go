package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Start the server
	router := http.NewServeMux()

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Health check")
		fmt.Fprintf(w, "OK")
	})

	router.HandleFunc("GET /echo/{message}", func(w http.ResponseWriter, r *http.Request) {
		message := r.PathValue("message")
		fmt.Println("Echo")
		fmt.Fprintf(w, "Echo: %s", message)
	})

	http.ListenAndServe(":8080", router)
}
