package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// Base URL for self-hosted Currency API
const currencyBaseURL = "http://129.241.150.113:9090/currency"

// fetchAllRates retrieves all exchange rates for a given base currency.
// Used by /exchange/{code}
func fetchAllRates(base string) (map[string]float64, int, error) {

	// HTTP client with timeout to prevent long blocking calls
	client := http.Client{Timeout: 4 * time.Second}

	url := currencyBaseURL + "/" + base
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, http.StatusServiceUnavailable, err
	}
	defer resp.Body.Close()

	// Propagate upstream error status if not OK
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, errors.New("currency upstream returned non-200")
	}

	// Read full response body for flexible JSON parsing
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusBadGateway, err
	}

	// First attempt: API returns direct map format
	var direct map[string]float64
	if err := json.Unmarshal(body, &direct); err == nil && len(direct) > 0 {
		return direct, http.StatusOK, nil
	}

	// Second attempt: API returns wrapped format { "rates": {...} }
	var wrapper struct {
		Rates map[string]float64 `json:"rates"`
	}
	if err := json.Unmarshal(body, &wrapper); err == nil && len(wrapper.Rates) > 0 {
		return wrapper.Rates, http.StatusOK, nil
	}

	// If neither format matches, treat as upstream error
	return nil, http.StatusBadGateway, errors.New("unknown currency format")
}