package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"assignment-1/models"
)

// Base URL for self-hosted REST Countries service
const restCountriesBase = "http://129.241.150.113:8080/v3.1"

// fetchCountryByCode retrieves full country information
// Used by /info/{code}
func fetchCountryByCode(code string) (models.RestCountriesCountry, int, error) {

	// HTTP client with timeout to avoid hanging requests
	client := http.Client{Timeout: 4 * time.Second}

	url := restCountriesBase + "/alpha/" + code
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.RestCountriesCountry{}, http.StatusInternalServerError, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return models.RestCountriesCountry{}, http.StatusServiceUnavailable, err
	}
	defer resp.Body.Close()

	// Propagate upstream status if not OK
	if resp.StatusCode != http.StatusOK {
		return models.RestCountriesCountry{}, resp.StatusCode, errors.New("upstream returned non-200")
	}

	// REST Countries returns an array → take first element
	var arr []models.RestCountriesCountry
	if err := json.NewDecoder(resp.Body).Decode(&arr); err != nil {
		return models.RestCountriesCountry{}, http.StatusBadGateway, err
	}
	if len(arr) == 0 {
		return models.RestCountriesCountry{}, http.StatusNotFound, errors.New("no country in upstream response")
	}

	return arr[0], http.StatusOK, nil
}

// fetchExchangeBaseByCode retrieves minimal country data
// Used by /exchange/{code}
func fetchExchangeBaseByCode(code string) (models.RestCountriesExchangeBase, int, error) {

	client := http.Client{Timeout: 4 * time.Second}

	url := restCountriesBase + "/alpha/" + code
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.RestCountriesExchangeBase{}, http.StatusInternalServerError, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return models.RestCountriesExchangeBase{}, http.StatusServiceUnavailable, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.RestCountriesExchangeBase{}, resp.StatusCode, errors.New("upstream returned non-200")
	}

	// Only required subset of fields is decoded
	var arr []models.RestCountriesExchangeBase
	if err := json.NewDecoder(resp.Body).Decode(&arr); err != nil {
		return models.RestCountriesExchangeBase{}, http.StatusBadGateway, err
	}
	if len(arr) == 0 {
		return models.RestCountriesExchangeBase{}, http.StatusNotFound, errors.New("no country in upstream response")
	}

	return arr[0], http.StatusOK, nil
}

// fetchCurrencyByAlpha3 retrieves currency code for a neighbouring country
// Alpha-3 codes (e.g., SWE, FIN) are accepted by REST Countries
func fetchCurrencyByAlpha3(alpha3 string) (string, int, error) {

	client := http.Client{Timeout: 4 * time.Second}

	url := restCountriesBase + "/alpha/" + alpha3
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", http.StatusServiceUnavailable, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", resp.StatusCode, errors.New("upstream returned non-200")
	}

	var arr []models.RestCountriesCurrencyOnly
	if err := json.NewDecoder(resp.Body).Decode(&arr); err != nil {
		return "", http.StatusBadGateway, err
	}
	if len(arr) == 0 {
		return "", http.StatusNotFound, errors.New("no country in upstream response")
	}

	// Extract first currency key from map
	for k := range arr[0].Currencies {
		return k, http.StatusOK, nil
	}

	return "", http.StatusBadGateway, errors.New("no currencies in upstream response")
}