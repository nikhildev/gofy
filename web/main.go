package main

import (
	"log"
	"net/http"

	"github.com/nikhildev/gofy/web/handlers"
)

func main() {
	log.Println("Starting Weather Application on port 8080")
	log.Println("Access at: http://localhost:8080")

	http.HandleFunc("/", handlers.WeatherHandler)
	http.HandleFunc("/api/search-location", handlers.SearchLocationHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
