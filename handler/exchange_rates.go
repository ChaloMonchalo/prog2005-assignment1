package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"assignment-1/models"
)

func ExchangeRatesHandler(w http.ResponseWriter, r *http.Request) {

	// Only GET is allowed
	if r.Method != http.MethodGet {
		writeErrorJSON(w, http.StatusMethodNotAllowed)
		return
	}

	// Extract and normalize country code
	code := strings.ToLower(strings.TrimSpace(r.PathValue("code")))
	if code == "" {
		writeErrorJSON(w, http.StatusBadRequest)
		return
	}

	// Fetch base country data (name, borders, currencies)
	base, status, err := fetchExchangeBaseByCode(code)
	if err != nil {
		writeErrorJSON(w, status)
		return
	}

	// Determine base currency (first key in currency map)
	baseCurrency := ""
	for k := range base.Currencies {
		baseCurrency = k
		break
	}
	if baseCurrency == "" {
		writeErrorJSON(w, http.StatusBadGateway)
		return
	}

	// Fetch all exchange rates for base currency
	allRates, st, err := fetchAllRates(baseCurrency)
	if err != nil {
		writeErrorJSON(w, st)
		return
	}

	exchangeRates := make(map[string]float64)
	seen := map[string]bool{}

	// For each neighbouring country:
	// 1. Get its currency
	// 2. Find corresponding rate
	// 3. Avoid duplicates
	for _, neighbour := range base.Borders {
		cur, _, err := fetchCurrencyByAlpha3(neighbour)
		if err != nil || cur == "" || cur == baseCurrency || seen[cur] {
			continue
		}
		seen[cur] = true

		if rate, ok := allRates[cur]; ok {
			exchangeRates[cur] = rate
		}
	}

	resp := models.ExchangeResponse{
		Country:       base.Name.Common,
		BaseCurrency:  baseCurrency,
		ExchangeRates: exchangeRates,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}