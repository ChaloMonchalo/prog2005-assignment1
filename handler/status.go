package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type statusResponse struct {
	RestCountriesAPI int    `json:"restcountriesapi"`
	CurrenciesAPI    int    `json:"currenciesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}

// Self-hosted upstream services (must use these)
const restCountriesURL = "http://129.241.150.113:8080/v3.1/alpha/no"
const currenciesURL = "http://129.241.150.113:9090/currency/NOK"

// helper function to check upstream status
func probe(url string) int {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return http.StatusServiceUnavailable
	}

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusServiceUnavailable
	}
	defer resp.Body.Close()

	return resp.StatusCode
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	restStatus := probe(restCountriesURL)
	currStatus := probe(currenciesURL)

	response := statusResponse{
		RestCountriesAPI: restStatus,
		CurrenciesAPI:    currStatus,
		Version:          "v1",
		Uptime:           UptimeSeconds(),
	}

	// 1) header first
	w.Header().Set("Content-Type", "application/json")

	// 2) status code once
	if restStatus != http.StatusOK || currStatus != http.StatusOK {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	// 3) body
	_ = json.NewEncoder(w).Encode(response)
}
