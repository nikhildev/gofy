// main initializes and starts a simple HTTP server on port 8080.
// The server has two endpoints:
// 1. GET /health - returns "OK" to indicate the server is running.
// 2. GET /echo/{message} - echoes back the provided message in the response.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	log.Printf("Starting server on port %d", port)
	router := http.NewServeMux()

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	router.HandleFunc("GET /echo/{message}", func(w http.ResponseWriter, r *http.Request) {
		message := r.PathValue("message")
		w.Write([]byte("Echo: " + message))
	})

	err := http.ListenAndServe(fmt.Sprintf(`:%d`, port), router)

	if err != nil {
		log.Fatalln("Failed to start server", err)
	}

}
