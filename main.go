package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"assignment-1/handler"
)

func main() {
	// Render provides PORT via environment variable.
	// Locally we default to 8080 if PORT is not set.
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Set service start time (for uptime calculation)
	handler.SetStartTime(time.Now())

	router := http.NewServeMux()
	
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("CountryInfo API v1 - Use /countryinfo/v1/... endpoints"))
})
	// Status
	router.HandleFunc("/countryinfo/v1/status/", handler.StatusHandler)

	// Info
	router.HandleFunc("/countryinfo/v1/info/", handler.InfoHandler)
	router.HandleFunc("/countryinfo/v1/info/{code}", handler.CountryInfoHandler)

	// Exchange (guide + real endpoint)
	router.HandleFunc("/countryinfo/v1/exchange/", handler.ExchangeHandler)
	router.HandleFunc("/countryinfo/v1/exchange/{code}", handler.ExchangeRatesHandler)

	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
